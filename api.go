package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func host(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal([]string{})
	if err != nil {
		fmt.Println("Error json stringifying stageRes : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(bytes)
}

func join(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal([]string{})
	if err != nil {
		fmt.Println("Error json stringifying stageRes : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(bytes)
}

func create(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal([]string{})
	if err != nil {
		fmt.Println("Error json stringifying stageRes : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(bytes)
}

func apiRouter() http.Handler {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
  }))

	r.Get("/join", join)
	r.Get("/create", create)
	r.Put("/host/{action}", host)
	return r
}
