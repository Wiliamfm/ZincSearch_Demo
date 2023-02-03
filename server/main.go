package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Wiliamfm/ZincSearch_Demo/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/zincSearch", func(r chi.Router) {
		r.Post("/", search)
	})
	http.ListenAndServe(":3000", r)
}

func search(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(models.SearchRequest)
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
