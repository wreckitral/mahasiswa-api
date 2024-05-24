package main

import (
	"encoding/json"
	"net/http"
)

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
