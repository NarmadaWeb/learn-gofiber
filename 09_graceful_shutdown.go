package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("Initializing Fiber application...")
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		reqID := time.Now().UnixNano()
		log.Printf("[Req %d] Received request to /", reqID)
		select {
		case <-time.After(4 * time.Second):
			log.Printf("[Req %d] Finished processing request to /", reqID)
			return c.SendString("Request processed successfully!")
		case <-c.Context().Done():
			log.Printf("[Req %d] Client disconnected before processing finished for /", reqID)
			return c.Context().Err()
		}
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		log.Println("Received request to /ping")
		return c.SendString("pong")
	})

	serverErrors := make(chan error, 1)

	go func() {
		listenAddr := ":3000"
		log.Printf("Server starting to listen on %s...", listenAddr)
		serverErrors <- app.Listen(listenAddr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case receivedSignal := <-quit:
		log.Printf("Received shutdown signal: %s. Initiating graceful shutdown...", receivedSignal)

	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("FATAL: Server failed to start: %v", err)
		}
		log.Println("Server error channel received:", err)
	}


	log.Println("Starting server shutdown process...")

	if err := app.Shutdown(); err != nil {
		log.Printf("ERROR during server shutdown: %v", err)
		os.Exit(1)
	}

	log.Println("Server shutdown completed gracefully.")

	log.Println("Application exited.")
	os.Exit(0)
}
