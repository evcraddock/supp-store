package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evcraddock/supp-store/models"
	"github.com/evcraddock/supp-store/services"
)

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	entry := models.SupplementEntry{}

	json.NewDecoder(r.Body).Decode(&entry)

	s := services.GetSession()

	entryService := services.InitSupplmentEntryService(s)

	err := entryService.Create(entry)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(201)
}

func GetEntries(w http.ResponseWriter, r *http.Request) {

	s := services.GetSession()

	entryService := services.InitSupplmentEntryService(s)

	entries, err := entryService.GetAll()

	if err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(entries)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)

}
