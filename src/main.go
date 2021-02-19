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
var vectorDatos []*estructuras.Lista

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A_tus_órdenes,_capitán... :D")
}

func cargartienda(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.Archivo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_ :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	linealizar(ms)
	fmt.Fprintln(w, "Datos_Guardados")
}

func tiendaEspecifica(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var tienda *estructuras.Tienda = buscarPosicion(busqueda)
	if tienda == nil {
		fmt.Fprintln(w, "No_se_encontró_la_tienda_solicitada")
	} else {
		json.NewEncoder(w).Encode(tienda)
	}
}

func buscarPosicion(request *estructuras.RequestFind) *estructuras.Tienda {
	indice := string(request.Nombre[0])
	var fila, columna int
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
	return vectorDatos[calcularPos(fila, columna, request.Calificacion)].Buscar(request.Nombre)
}

func calcularPos(fila int, columna int, calificacion int) int {
	var s int
	s = columna*len(indices) + fila
	fmt.Print(columna)
	fmt.Print(s)
	fmt.Print(s*5 + calificacion - 1)
	return s*5 + calificacion - 1
}

func id(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numero, err := strconv.Atoi(vars["numero"])
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	if numero >= len(vectorDatos) || len(vectorDatos) == 0 || vectorDatos[numero].Size == 0 {
		fmt.Fprintf(w, "No_Encontrado")
	} else {
		var lista []*estructuras.Tienda
		var aux *estructuras.Nodo = vectorDatos[numero].First
		for i := 0; i < vectorDatos[numero].Size; i++ {
			lista = append(lista, aux.Tienda)
			aux = aux.Next
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lista)
	}
}

func eliminar(w http.ResponseWriter, r *http.Request) {
	var eliminar *estructuras.DeleteReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &eliminar)
	var tienda *estructuras.Tienda = eliminarPosicion(eliminar)
	if tienda == nil {
		fmt.Fprintln(w, "No_se_encontró_la_tienda_solicitada")
	} else {
		fmt.Fprintln(w, "Eliminado.")
	}
}

func eliminarPosicion(request *estructuras.DeleteReq) *estructuras.Tienda {
	indice := string(request.Nombre[0])
	var fila, columna int
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
		if nombresDep[j] == request.Categoria {
			columna = j
			break
		}
		if j == len(nombresDep)-1 {
			return nil
		}
	}
	return vectorDatos[calcularPos(fila, columna, request.Calificacion)].Eliminar(request.Nombre)
}

func guardar(w http.ResponseWriter, r *http.Request) {
	if len(vectorDatos) == 0 {
		fmt.Fprintf(w, "No_Existen_Datos_Cargados")
	} else {
		var file estructuras.Archivo = estructuras.Archivo{}
		var pos int = 0
		for i := 0; i < len(indices); i++ {
			file.Datos = append(file.Datos, &estructuras.Dato{Indice: indices[i]})
			for j := 0; j < len(nombresDep); j++ {
				file.Datos[i].Departamentos = append(file.Datos[i].Departamentos, &estructuras.Departamento{Nombre: nombresDep[j]})
			}
		}
		for i := 0; i < len(file.Datos[0].Departamentos); i++ {
			for j := 0; j < len(file.Datos); j++ {
				for k := 0; k < 5; k++ {
					file.Datos[j].Departamentos[i].Tiendas = append(file.Datos[j].Departamentos[i].Tiendas, *vectorDatos[pos+k].ToArray()...)
				}
				pos += 5
				if len(file.Datos[j].Departamentos[i].Tiendas) == 0 {
					file.Datos[j].Departamentos[i].Tiendas = make([]*estructuras.Tienda, 0)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(file)
	}
}

func linealizar(ms *estructuras.Archivo) {
	var vector []*estructuras.Lista
	var letras []string
	var nombres []string
	for i := 0; i < len(ms.Datos); i++ {
		letras = append(letras, ms.Datos[i].Indice)
	}
	for i := 0; i < len(ms.Datos[0].Departamentos); i++ {
		nombres = append(nombres, ms.Datos[0].Departamentos[i].Nombre)
	}
	indices = letras
	nombresDep = nombres
	for i := 0; i < len(ms.Datos[0].Departamentos); i++ {
		for j := 0; j < len(ms.Datos); j++ {
			for l := 0; l < 5; l++ {
				vector = append(vector, estructuras.NewLista())
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
	vectorDatos = vector
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("POST")
	router.HandleFunc("/id/{numero}", id).Methods("GET")
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
