package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main()  {
	router :=mux.NewRouter()
	router.HandleFunc("/shorten",createShortURL).Methods("POST")
	router.HandleFunc("/{shortURL}",redirectURL).Methods("GET")
	router.HandleFunc("stats/{shortURL}",getURLStats).Methods("GET")
	http.Handle("/",router)
	log.Println("Server runinng on port 8080")
	log.Fatal(http.ListenAndServe(":8080",router))

}

