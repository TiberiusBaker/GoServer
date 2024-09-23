package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TiberiusBaker/GoServer/pkg/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	http.Handle("/", r)
	fmt.Println("Server running on port 8000")
	config.GetDB()
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
