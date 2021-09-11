package main

import (
	"context"
	"go-tutorial/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodBye(l)

	// Create a ServeMux
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	// Spins up a http server, first arg is the port, second one is the http handler (default handler, see DefaultServerMux)
	// ServeMux is like a "controller" that maps a pattern to a handler
	s := &http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal()
		}
	}()

	// Signal channel
	// Kill the server when the channel receives appropriate signal e.g. Ctrl C in cmd
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// Graceful shutdown. Useful for cases like server upgrades.
	// Waits for existing requests to finish before shutdown.
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
