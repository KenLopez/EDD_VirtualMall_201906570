package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./estructuras"
)

var matriz []estructuras.Lista
var ms estructuras.Archivo

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A tus órdenes, capitán... :D")

}

func cargartienda(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Hola")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	fmt.Fprintln(w, "Funciona")
}

func crearMatriz(data *estructuras.Archivo) {

}

func main() {
	lista := estructuras.NewLista()
	nodo := estructuras.NewNodo()
	nodo.Tienda.Nombre = "Coca"
	lista.Insertar(nodo)
	nodo2 := estructuras.NewNodo()
	nodo2.Tienda.Nombre = "Mario"
	lista.Insertar(nodo2)
	nodo3 := estructuras.NewNodo()
	nodo3.Tienda.Nombre = "Kenneth"
	lista.Insertar(nodo3)
	nodo4 := estructuras.NewNodo()
	nodo4.Tienda.Nombre = "Ana"
	lista.Insertar(nodo4)
	nodo5 := estructuras.NewNodo()
	nodo5.Tienda.Nombre = "maca"
	lista.Insertar(nodo5)
	aux := lista.First
	for i := 0; i < lista.Size; i++ {
		fmt.Println(aux.Tienda.Nombre)
		aux = aux.Next
	}
	/*router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))*/
}
