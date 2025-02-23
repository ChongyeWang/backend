# How to Run a Gin Project

This guide will walk you through the steps to set up and run this project.

## Prerequisites

1. **Go installed**: Ensure you have Go installed on your system.
2. **GOPATH setup**: Make sure your `GOPATH` and `GOBIN` environment variables are properly configured.
3. **Git installed**: You'll need Git to clone the project repository.

## Steps to Run a Gin Project

Clone the Project Repository

go build

go run main.go

# API Documentation

## **GET /orders

Retrieves a list of all orders.

## **POST /orders

Add an order to database.

# CICD

A Git CI/CD pipeline will automatically trigger deployments to an EC2 instance running Docker. 

The branch used here default is master.

