package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khayss/seek/pkg/interfaces"
)

type TransactionForCreate struct {
	SourceTxID       string `json:"source_tx_id" binding:"required"`
	FromChainID      int    `json:"from_chain_id" binding:"required"`
	ToChainID        int    `json:"to_chain_id" binding:"required"`
	FromAddress      string `json:"from_address" binding:"required"`
	ToAddress        string `json:"to_address" binding:"required"`
	Amount           string `json:"amount" binding:"required"`
	FromTokenAddress string `json:"from_token_address" binding:"required"`
	ToTokenAddress   string `json:"to_token_address" binding:"required"`
}

func CreateTransactionHandler(app interfaces.AppInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// db := app.GetDB()
		// Implement the logic to create a transaction
		var tx TransactionForCreate

		if err := c.ShouldBindJSON(&tx); err != nil {
			c.JSON(400, gin.H{
				"error":   "Invalid request payload",
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Transaction created",
		})
	}
}

func GetTransactionsHandler(c *gin.Context) {
	// Implement the logic to get transactions
	c.JSON(200, gin.H{
		"message": "Get transactions",
	})
}

func GetTransactionByIDHandler(c *gin.Context) {
	// Implement the logic to get a transaction by ID
	c.JSON(200, gin.H{
		"message": "Get transaction by ID",
	})
}

func UpdateTransactionHandler(c *gin.Context) {
	// Implement the logic to update a transaction
	c.JSON(200, gin.H{
		"message": "Transaction updated",
	})
}

func DeleteTransactionHandler(c *gin.Context) {
	// Implement the logic to delete a transaction
	c.JSON(200, gin.H{
		"message": "Transaction deleted",
	})
}
