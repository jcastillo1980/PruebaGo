package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func funcGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "valor: %s", params["valor"])
}

func main() {
	fmt.Println("Inicio del programa")
	MiMux := mux.NewRouter().StrictSlash(true)
	MiMux.HandleFunc("/api/{valor}", funcGet).Methods("GET")
	MiMux.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	MiServer := &http.Server{
		Addr:           ":5050",
		Handler:        MiMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(MiServer.ListenAndServe())
}
