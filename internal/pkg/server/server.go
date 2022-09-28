package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"littlevincee.com/pokemon-api-gateway/internal/pkg/logger"
)

type Server interface {
	Start(port string, handler http.Handler)
}

type server struct {
	Server
}

func New() *server {
	return &server{}
}
func (s server) Start(port string, handler http.Handler) {
	log := logger.New()

	apiServer := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	log.Infof("Pokemon API Gateway is listening at %s", port)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatalf("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		log.Infof("graceful shutdown in process")

		err := apiServer.Shutdown(shutdownCtx)

		if err != nil {
			log.Fatalf(err.Error())
		}

		serverStopCtx()
	}()

	// Run the server
	err := apiServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
