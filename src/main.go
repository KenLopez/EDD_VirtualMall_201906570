package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	linealizar()
	fmt.Fprintln(w, "Listo")
}

func tiendaEspecifica(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No jaló :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var tienda *estructuras.Tienda = buscarPosicion(busqueda)
	if tienda == nil {
		fmt.Fprintln(w, "No se encontró la tienda solicitada")
	} else {
		json.NewEncoder(w).Encode(tienda)
	}
}

func buscarPosicion(request *estructuras.RequestFind) *estructuras.Tienda {
	indice := string(request.Nombre[0])
	var fila, columna, p, s, t int
	for i := 0; i < len(indices); i++ {
		if indices[i] == indice {
			fila = i
			break
		}
		if i == len(indices)-1 {
			return nil
		}
	}
	for j := 0; j < len(nombresDep); j++ {
		if nombresDep[j] == request.Departamento {
			columna = j
			break
		}
		if j == len(nombresDep)-1 {
			return nil
		}
	}
	p = columna
	s = p*len(indices) + fila
	t = s*5 + request.Calificacion - 1

	return vector[t].Buscar(request.Nombre)
}

func id(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	if numero >= len(vector) || len(vector) == 0 || vector[numero].Size == 0 {
		fmt.Fprintf(w, "NoEncontrado")
	} else {
		var lista []*estructuras.Tienda
		var aux *estructuras.Nodo = vector[numero].First
		for i := 0; i < vector[numero].Size; i++ {
			lista = append(lista, aux.Tienda)
			aux = aux.Next
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lista)
	}

}

func linealizar() {
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
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("POST")
	router.HandleFunc("/id/{numero}", id).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
