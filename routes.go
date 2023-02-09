package main

import (
	"encoding/json"
	"finder/internal/api"
	"finder/internal/search"
	"net/http"
)

type Storage interface {
	Exists(int, int) bool
	Category(int) []int
}

func writeResult(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}

func createHTTPHandler(config Config, storage Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/v1/search":
			var payload api.Payload
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				writeResult(w, r, http.StatusBadRequest, api.Result{Error: err.Error()})
				return
			}

			result := search.Search(payload)
			writeResult(w, r, http.StatusOK, result)
		default:
			writeResult(w, r, http.StatusBadRequest, api.Result{Error: "Not Found"})
			return

		}
	})
}
