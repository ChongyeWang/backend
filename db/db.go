package database

import (
	"backend/models"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./orders.db")
	if err != nil {
		return err
	}

	query := `
    CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		symbol TEXT NOT NULL,
		price REAL NOT NULL,
		quantity INTEGER NOT NULL,
		type TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending'
	);`
	_, err = DB.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Database initialized successfully")
	return nil
}

func CreateOrder(order models.Order) error {
	query := `INSERT INTO orders (symbol, price, quantity, type) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, order.Symbol, order.Price, order.Quantity, order.Type)
	return err
}

func GetOrders() ([]models.Order, error) {
	rows, err := DB.Query("SELECT id, symbol, price, quantity, type FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Symbol, &order.Price, &order.Quantity, &order.Type)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrderStatus(orderID int, status string) error {
	query := `UPDATE orders SET status = ? WHERE id = ?`
	_, err := DB.Exec(query, status, orderID)
	return err
}
