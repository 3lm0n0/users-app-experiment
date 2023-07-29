package main

import (
	"net/http"
	"time"
)

const (
	defaultPort = "8008"
	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

type Server interface {
	Start() error
}

type ServerParams struct {
	listenAddress string
}

func NewServer(listenAddress string) Server {
	return &ServerParams{
		listenAddress: listenAddress,
	}
}

func(sp *ServerParams) Start() error {
	server := &http.Server{
		Addr:    "0.0.0.0:" + sp.listenAddress,

		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}
	return server.ListenAndServe()
}

