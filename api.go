package main

import (
	"log"
	"net/http"
)

type APIServer struct {
    listenAddr string
    Store   Storage
}

func NewAPIServer(listenAddr string, s Storage) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
        Store: s,
    }
}

func (s *APIServer) Run() error {
    router := http.NewServeMux()

    router.HandleFunc("/rata-rata", makeHttpHandlerFunc(s.getRataRata))
    router.HandleFunc("/rata-rata-suliet", makeHttpHandlerFunc(s.getRataRataSuliet))

    server := http.Server{
        Addr: s.listenAddr,
        Handler: router,
    }

    log.Println("mahasiswa-api Server is running on port:", s.listenAddr)

    return server.ListenAndServe()
}

func(s *APIServer) getRataRata(res http.ResponseWriter, req *http.Request) error {
    avgIpk, err := s.Store.GetRataRata()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, map[string]float64{"rata-rata IPK":avgIpk})
}

func(s *APIServer) getRataRataSuliet(res http.ResponseWriter, req *http.Request) error {
    avgSuliet, err := s.Store.getRataRataSuliet()
    if err != nil {
        return err
    }
    
    return writeJSON(res, http.StatusOK, map[string]float64{"rata-rata SULIET": avgSuliet})
}
