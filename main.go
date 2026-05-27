package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "8080"

// main is the entry point of the application.
// it sets up the HTTP server, starts it in a separate goroutine,
// and listens for shutdown signals to gracefully stop the server.
func main() {
	port := getPort()
	server := newServer(port, routes())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on port %s", port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server stopped.")
}

// getPort returns the HTTP server port.
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}

// routes configures and returns the HTTP request multiplexer.
func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	return mux
}

// newServer creates and configures a new HTTP server instance.
func newServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
}

// ----- helper functions -----
// helloHandler handles requests to the root endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}

// healthHandler handles health check requests.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
