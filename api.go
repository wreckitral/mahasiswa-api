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
    router.HandleFunc("/max-ipk", makeHttpHandlerFunc(s.getMahasiswaMaxIpk))
    router.HandleFunc("/min-ipk", makeHttpHandlerFunc(s.getMahasiswaMinIpk))
    router.HandleFunc("/max-suliet", makeHttpHandlerFunc(s.getMahasiswaMaxSuliet))
    router.HandleFunc("/min-suliet", makeHttpHandlerFunc(s.getMahasiswaMinSuliet))
    router.HandleFunc("/agregat", makeHttpHandlerFunc(s.getAgregat))

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

    return writeJSON(res, http.StatusOK, map[string]float64{"rata-rata-ipk":avgIpk})
}

func(s *APIServer) getRataRataSuliet(res http.ResponseWriter, req *http.Request) error {
    avgSuliet, err := s.Store.GetRataRataSuliet()
    if err != nil {
        return err
    }
    
    return writeJSON(res, http.StatusOK, map[string]float64{"rata-rata-suliet": avgSuliet})
}

func(s *APIServer) getMahasiswaMaxIpk(res http.ResponseWriter, req *http.Request) error {
    mahasiswa, err := s.Store.GetMahasiswaMaxIpk()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, mahasiswa)
}

func(s *APIServer) getMahasiswaMinIpk(res http.ResponseWriter, req *http.Request) error {
    mahasiswa, err := s.Store.GetMahasiswaMinIpk()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, mahasiswa)
}

func(s *APIServer) getMahasiswaMaxSuliet(res http.ResponseWriter, req *http.Request) error {
    mahasiswa, err := s.Store.GetMahasiswaMaxSuliet()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, mahasiswa)
}

func(s *APIServer) getMahasiswaMinSuliet(res http.ResponseWriter, req *http.Request) error {
    mahasiswa, err := s.Store.GetMahasiswaMinSuliet()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, mahasiswa)
}

func(s *APIServer) getAgregat(res http.ResponseWriter, req *http.Request) error {
    pujian, sangatMemuaskan, memuaskan, err := s.Store.GetPujianSangatMemuaskanMemuaskan()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, map[string]int{
        "dengan-pujian": pujian,
        "sangat-memuaskan": sangatMemuaskan,
        "memuaskan": memuaskan,
    })
}
