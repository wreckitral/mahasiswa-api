package main

import "net/http"


type APIServer struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
    }
}

