package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./nocache"

	"github.com/gorilla/mux"
)

func funcGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")

	fmt.Fprintf(w, "Valor: {%s}", params["valor"])
}

func funcGetOS(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	salida := "Variables:\r\n"
	salida = salida + "--------------------\r\n"
	for i, v := range os.Environ() {
		salida = salida + fmt.Sprintf("%d : ", i) + v + "\r\n"
	}
	fmt.Fprint(w, salida)
}

func main() {
	rutaStatic := flag.String("v", "./public", "Ruta de la carpeta fichero estaticos")
	puerto := flag.Int("p", 5050, "Puerto TCP servidor")
	flag.Parse()

	fmt.Println("Inicio del Servidor:")
	fmt.Printf("http://127.0.0.1:%d\r\n", *puerto)
	fmt.Printf("Ruta: [%s]\r\n", *rutaStatic)
	MiMux := mux.NewRouter().StrictSlash(false)
	MiMux.HandleFunc("/api/repite/{valor}", funcGet).Methods("GET")
	MiMux.HandleFunc("/api/os", funcGetOS).Methods("GET")

	MiMux.PathPrefix("/").Handler(nocache.NoCache(http.FileServer(http.Dir(fmt.Sprintf("%s", *rutaStatic)))))
	MiServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", *puerto),
		Handler:        MiMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       log.New(os.Stdout, "SERVER", 0),
	}

	log.Fatalln(MiServer.ListenAndServe())
}
