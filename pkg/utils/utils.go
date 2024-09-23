package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonReturn(f func(w http.ResponseWriter, r *http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := f(w, r)
		if err != nil {
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}