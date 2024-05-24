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

    router.HandleFunc("/rata-rata", makeHttpHandlerFunc(s.getRataRata))

    server := http.Server{
        Addr: s.listenAddr,
    }

    log.Println("mahasiswa-api Server is running on port:", s.listenAddr)

    return server.ListenAndServe()
}

func(s *APIServer) getRataRata(res http.ResponseWriter, req *http.Request) error {
    return writeJSON(res, http.StatusOK, map[string]string{"message":"rata-rata is called"})
}
