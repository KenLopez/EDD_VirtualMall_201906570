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

var indices, nombresDep []string
var vector []estructuras.Lista
var ms estructuras.Archivo

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A tus órdenes, capitán... :D")

}

func cargartienda(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No jaló :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	crearMatriz()
	fmt.Fprintln(w, vector[4].First.Tienda.Nombre)
}

func crearMatriz() {
	for i := 0; i < len(ms.Datos); i++ {
		indices = append(indices, ms.Datos[i].Indice)
	}
	//fmt.Println(filas)
	for i := 0; i < len(ms.Datos[0].Departamentos); i++ {
		nombresDep = append(nombresDep, ms.Datos[0].Departamentos[i].Nombre)
	}
	//fmt.Println(columnas)
	for i := 0; i < len(ms.Datos[0].Departamentos); i++ {
		for j := 0; j < len(ms.Datos); j++ {
			for l := 0; l < 5; l++ {
				vector = append(vector, *estructuras.NewLista())
			}
			for k := 0; k < len(ms.Datos[j].Departamentos[i].Tiendas); k++ {
				var nodo *estructuras.Nodo = estructuras.NewNodo(ms.Datos[j].Departamentos[i].Tiendas[k])
				if ms.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 1 {
					vector[len(vector)-5].Insertar(nodo)
				} else if ms.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 2 {
					vector[len(vector)-4].Insertar(nodo)
				} else if ms.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 3 {
					vector[len(vector)-3].Insertar(nodo)
				} else if ms.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 4 {
					vector[len(vector)-2].Insertar(nodo)
				} else if ms.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 5 {
					vector[len(vector)-1].Insertar(nodo)
				}
			}
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}