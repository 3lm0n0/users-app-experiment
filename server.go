package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	defaultPort 		= "8008"
	idleTimeout       	= 30 * time.Second
	writeTimeout      	= 180 * time.Second
	readHeaderTimeout 	= 10 * time.Second
	readTimeout       	= 10 * time.Second
    shutdownTimeout 	= 10 * time.Second
)

type Server struct {
    httpServer *http.Server
    shutdownCh chan struct{}
    wg         sync.WaitGroup
    mu         sync.Mutex
    logger     *log.Logger
}

func NewServer(listenAddress string, logger *log.Logger) *Server {
	if listenAddress == "" {
		listenAddress = defaultPort
	}
    return &Server{
        httpServer: &http.Server{
            Addr:               listenAddress,
            IdleTimeout:        idleTimeout,
            WriteTimeout:       writeTimeout,
            ReadTimeout:        readTimeout,
            ReadHeaderTimeout:  readHeaderTimeout,
            ErrorLog:           log.New(log.Writer(), "ERROR: ", log.LstdFlags),
        },
        shutdownCh: make(chan struct{}),
        logger:     logger,
    }
}

func (s *Server) Start() error {
	s.logger.Printf("Server is starting on port: %s", s.httpServer.Addr)
	s.httpServer.ListenAndServe()
    /*err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatal(err)
		return err
	}*/
	return nil
}

func (s *Server) Shutdown() error {
	s.logger.Printf("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	close(s.shutdownCh)
	s.wg.Wait()

	s.logger.Printf("Server gracefully stopped")

	return nil
}

func (s *Server) ShutdownChannel() <-chan struct{} {
	return s.shutdownCh
}