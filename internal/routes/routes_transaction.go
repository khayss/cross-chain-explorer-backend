package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khayss/seek/internal/handlers"
	"github.com/khayss/seek/pkg/interfaces"
)

func AddTransactionRoutes(group *gin.RouterGroup, app interfaces.AppInterface) {

	tx := group.Group("/tx")
	tx.POST("/", handlers.CreateTransactionHandler(app))
	tx.GET("/", handlers.GetTransactionsHandler)
	tx.GET("/:id", handlers.GetTransactionByIDHandler)
	tx.PUT("/:id", handlers.UpdateTransactionHandler)
	tx.DELETE("/:id", handlers.DeleteTransactionHandler)
}
