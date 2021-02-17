package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./estructuras"
	"github.com/gorilla/mux"
)

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A tus órdenes, capitán... :D")

}

func cargartienda(w http.ResponseWriter, r *http.Request) {
	var ms estructuras.Archivo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No jaló :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	fmt.Fprintln(w, ms.Datos[0].Departamentos[0].Tiendas.First)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
