package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/TiberiusBaker/GoServer/pkg/models"
	"github.com/go-chi/chi/v5"
)

func JsonReturn(f func(w http.ResponseWriter, r *http.Request) (interface{}, error), success int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Failed to marshal JSON: " + err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		success = http.StatusOK
		w.WriteHeader(success)
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

func ParseId(r *http.Request, paramName models.ContextKey) (string, error) {
	id := chi.URLParam(r, string(paramName))
	_, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		return "", fmt.Errorf("invalid id: %v", err)
	}
	return id, nil
}
