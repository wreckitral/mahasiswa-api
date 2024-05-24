package main

import (
	"log"
	"net/http"
)

type APIServer struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
    }
}

func (s *APIServer) Run() error {
    router := http.NewServeMux()

    server := http.Server{
        Addr: s.listenAddr,
    }

    log.Println("mahasiswa-api Server is running on port:", s.listenAddr)

    return server.ListenAndServe()
}
