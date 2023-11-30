package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const port = ":8080"

func runWithGracefulShutdown(router any) {
	srv := &http.Server{
		Addr:    port,
		Handler: router.(http.Handler),
	}

	go func() {
		log.Printf("Listening on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)
	<-q

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Could not shutdown: %s\n", err)
	}

}
