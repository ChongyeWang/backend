package websocket

import (
	database "backend/db"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

var statusMap = map[string]string{
	"1": "pending",
	"2": "processing",
	"3": "completed",
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		statusCode := string(msg)
		log.Printf("Received status code: %s\n", statusCode)

		status, exists := statusMap[statusCode]
		if !exists {
			log.Println("Invalid status code received:", statusCode)
			continue
		}

		orderID := 1

		err = database.UpdateOrderStatus(orderID, status)
		if err != nil {
			log.Println("Error updating order status:", err)
			continue
		}

		BroadcastMessage("Order " + strconv.Itoa(orderID) + " status updated to " + status)
	}
}

func BroadcastMessage(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("WebSocket broadcast error:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
