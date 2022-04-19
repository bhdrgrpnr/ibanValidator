package main

import (
	api "IbanValidator/internal/http"
	"IbanValidator/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	ctx := context.Background()

	handler := api.NewHandler()
	service.InitService()

	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("starting parcel API on port %d...", ":8080")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()



	server.Shutdown(ctx)
}
