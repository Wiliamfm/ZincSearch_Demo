package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Wiliamfm/ZincSearch_Demo/server/models"
	indexer "github.com/Wiliamfm/ZincSearch_Demo/server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/zincsearch", func(r chi.Router) {
		r.Post("/", search)
	})
	path := os.Args[1]
	if !indexer.Index(path) {
		log.Fatal("Could not load emails to ZincSearch")
	}
	log.Printf("Emails loaded\nServer running on %d\n", 3000)
	http.ListenAndServe(":3000", r)
}

func search(w http.ResponseWriter, r *http.Request) {
	data := models.ClientRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(err, w, 500)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		sendError(err, w, 500)
		return
	}
	reqSearch := models.SearchRequest{SearchType: "match", Query: models.Query{Term: data.Search, StartTime: "", EndTime: ""}, SortFields: "", From: 0, MaxResults: 20, Source: make([]string, 0)}
	jsonData, err := json.Marshal(reqSearch)
	if err != nil {
		sendError(err, w, 500)
		return
	}
	log.Println(string(jsonData))
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		sendError(err, w, 500)
		return
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func sendError(err error, res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	data, err := json.Marshal(map[string]string{"error": err.Error()})
	log.Println(string(data))
	if err != nil {
		log.Fatal(err)
	}
	res.Write(data)
}
