package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/config"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
	"github.com/spf13/viper"
)

func main() {

	config.Load()
	repo.ConnectDB()

	router := http.NewServeMux()

	httpPort := viper.GetString("HTTP_PORT")
	server := &http.Server{
		Addr:    httpPort,
		Handler: router,
	}

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server working"))
	})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		log.Printf("Server running on %s", httpPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()
	<-done
	fmt.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown failed")
	}

}
