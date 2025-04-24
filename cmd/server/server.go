package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khayss/seek/internal/app"
	"github.com/khayss/seek/internal/config"
)

func main() {
	// Load configuration
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("[SEEK] failed to load config: %v", err)
	}

	app, err := app.NewApp(config.GetDbUrl())
	if err != nil {
		log.Fatalf("[SEEK] failed to create app: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	// Register routes
	app.RegisterRoutes(router)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[SEEK] Listen:	%s\n", err)
		}
	}()

	<-ctx.Done()
	stop()

	log.Println("[SEEK] Shutting down server gracefully...")
	log.Println("[SEEK] Press Ctrl+C again to force shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("[SEEK] Server forced to shutdown: %v", err)
	}

	log.Println("[SEEK] Server exiting")
}
