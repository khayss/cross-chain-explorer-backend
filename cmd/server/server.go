package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khayss/cross-chain-explorer-backend/internal/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/healthz", handlers.HealthCheck)
	router.POST("/transaction", handlers.CreateTransactionHandler)
	router.GET("/transactions", handlers.GetTransactionsHandler)
	router.GET("/transaction/:id", handlers.GetTransactionByIDHandler)
	router.PUT("/transaction/:id", handlers.UpdateTransactionHandler)
	router.DELETE("/transaction/:id", handlers.DeleteTransactionHandler)

	router.Run(":8080")
}
