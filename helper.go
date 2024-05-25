package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

/*
HELPER FUNCTION API
*/
type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
    Error string `json:"error"`
}

func writeJSON(res http.ResponseWriter, status int, body any) error {
    res.Header().Add("Content-Type", "application/json")
    res.WriteHeader(status)

    return json.NewEncoder(res).Encode(body)
}

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
    return func(res http.ResponseWriter, req *http.Request) {
        if err := f(res, req); err != nil {
            writeJSON(res, http.StatusInternalServerError, apiError{err.Error()})
        }

        if req.Method != "GET" {
            writeJSON(res, http.StatusInternalServerError, apiError{"Wrong Method"})
        }
    } 
}

/*
HELPER FUNCTION FOR STORAGE
*/
func scanIntoMahasiswa(rows *sql.Rows) (*Mahasiswa, error) {
    mahasiswa := Mahasiswa{}

    err := rows.Scan(&mahasiswa.Nama, &mahasiswa.NIM, &mahasiswa.TmptLahir, 
        &mahasiswa.TglLahir, &mahasiswa.IPK, &mahasiswa.Predikat, &mahasiswa.Prodi, &mahasiswa.MasaStudi, &mahasiswa.Suliet)

    return &mahasiswa, err

}

/*
HELPER FUNCTION FOR AGE FETCHING
*/
func scanIntoDataUmur(rows *sql.Rows) (*DataUmur, error) {
    dataUmur := DataUmur{}

    err := rows.Scan(&dataUmur.UmurMahasiswa)

    return &dataUmur, err
}
