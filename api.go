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
    router.HandleFunc("/agregat-predikat", makeHttpHandlerFunc(s.getAgregatPredikat))
    router.HandleFunc("/agregat-masa-studi", makeHttpHandlerFunc(s.getAgregatMasaStudi))
    router.HandleFunc("/data-umur", makeHttpHandlerFunc(s.getDataUmur))

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

func(s *APIServer) getAgregatPredikat(res http.ResponseWriter, req *http.Request) error {
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

func(s *APIServer) getAgregatMasaStudi(res http.ResponseWriter, req *http.Request) error {
    agregat, err := s.Store.GetAgregatMasaStudi()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, agregat)
}

func(s *APIServer) getDataUmur(res http.ResponseWriter, req *http.Request) error {
    dataUmur, err := s.Store.GetDataUmurMahasiswa()
    if err != nil {
        return err
    }

    return writeJSON(res, http.StatusOK, dataUmur)
}
