package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/app"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/config"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
	"github.com/spf13/viper"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("error loading config file %s", err.Error())
	}

	db, err := repo.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
	}

	services := app.InitServices(db)

	router := app.InitRouter(services)

	httpPort := viper.GetString("HTTP_PORT")
	server := &http.Server{
		Addr:    httpPort,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		log.Printf("Server running on %s", httpPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()
	<-done
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed")
	}
}
