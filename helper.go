package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(res http.ResponseWriter, status int, body any) error {
    res.Header().Add("Content-Type", "application/json")
    res.WriteHeader(status)

    return json.NewEncoder(res).Encode(body)
}
