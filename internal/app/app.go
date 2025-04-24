package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/khayss/seek/internal/handlers"
	"github.com/khayss/seek/internal/routes"
)

type App struct {
	DB *sql.DB
}

func NewApp(dbUrl string) (*App, error) {
	var app App

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return &app, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return &app, err
	}

	app.DB = db

	return &app, nil
}

func (a *App) GetDB() *sql.DB {
	return a.DB
}

func (a *App) RegisterRoutes(router *gin.Engine) {
	router.GET("/healthz", handlers.HealthCheck)

	v1 := router.Group("/v1")
	routes.AddTransactionRoutes(v1, a)

}
