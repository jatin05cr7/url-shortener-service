package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)
type URL struct{
	ID string `json:"id"`
	Orignal string `json:"orignal"`
	Short string `json:"short"`
	Clicks int `json:clicks`
}
var urlStore= make(map[string]URL)

func createShortURL(w http.ResponseWriter,r *http.Request)  {
	var url URL
	json.NewEncoder(r.Body).Decode(&url)
	url.ID=generateShortURLID()
}