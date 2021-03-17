package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

/*func cargarPedidos(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.ArchivoPedido
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_ :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.Pedidos); i++ {
		nodo := buscarPosicion(&estructuras.RequestFind{
			Departamento: ms.Inventarios[i].Departamento,
			Nombre:       ms.Inventarios[i].Tienda,
			Calificacion: ms.Inventarios[i].Calificacion,
		})
		if nodo == nil {
			fmt.Fprintln(w, "No_se_encontró_tienda:"+ms.Inventarios[i].Tienda+"-;")
		} else {
			if nodo.Inventario == nil {
				nodo.Inventario = estructuras.NewArbol()
			}
			for j := 0; j < len(ms.Inventarios[i].Productos); j++ {
				//fmt.Println(ms.Inventarios[i].Productos[j].Nombre)
				nodo.Inventario.Insertar(ms.Inventarios[i].Productos[j], ms.Inventarios[i].Productos[j].Codigo)
			}
		}
	}
	fmt.Fprintln(w, "Productos_Cargados")
}*/

func cargarInventarios(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.ArchivoInventario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_ :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.Inventarios); i++ {
		nodo := buscarPosicion(&estructuras.RequestFind{
			Departamento: ms.Inventarios[i].Departamento,
			Nombre:       ms.Inventarios[i].Tienda,
			Calificacion: ms.Inventarios[i].Calificacion,
		})
		if nodo == nil {
			fmt.Fprintln(w, "No_se_encontró_tienda:"+ms.Inventarios[i].Tienda+"-;")
		} else {
			if nodo.Contenido.(*estructuras.NodoTienda).Inventario == nil {
				nodo.Contenido.(*estructuras.NodoTienda).Inventario = estructuras.NewArbol()
			}
			for j := 0; j < len(ms.Inventarios[i].Productos); j++ {
				//fmt.Println(ms.Inventarios[i].Productos[j].Nombre)
				nodo.Contenido.(*estructuras.NodoTienda).Inventario.Insertar(ms.Inventarios[i].Productos[j], ms.Inventarios[i].Productos[j].Codigo)
			}
		}
	}
	fmt.Fprintln(w, "Productos_Cargados")
}

func tiendaEspecifica(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var nodo *estructuras.Nodo = buscarPosicion(busqueda)
	if nodo == nil {
		fmt.Fprintln(w, "No_se_encontró_la_tienda_solicitada")
	} else {
		var tienda *estructuras.Tienda = nodo.Contenido.(*estructuras.NodoTienda).Tienda
		json.NewEncoder(w).Encode(tienda)
	}
}

func buscarPosicion(request *estructuras.RequestFind) *estructuras.Nodo {
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
	pos := calcularPos(fila, columna, request.Calificacion)
	if pos >= len(vectorDatos) {
		return nil
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
			lista = append(lista, aux.Contenido.(*estructuras.NodoTienda).Tienda)
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

		jsonData, _ := json.MarshalIndent(file, "", "	")
		_ = ioutil.WriteFile("Datos.json", jsonData, 0644)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(file)
	}
}

func getArreglo(w http.ResponseWriter, r *http.Request) {
	var listas, posiciones, conexionesV, conexionesL string
	var countPos, countFila, countColumna, countList, numCluster int
	posiciones = "digraph G{\ncompound=true;\nsubgraph cluster0{" +
		"style=invis;\nedge[minlen=0.1, dir=fordware]\n"
	listas = ""
	countColumna = 0
	numCluster = 1
	for i := 0; i < len(vectorDatos)-1; i++ {
		conexionesV += "struct" + strconv.Itoa(i) + "->struct" + strconv.Itoa(i+1) +
			"[arrowhead=box, color=\"#9100d4\"];\n"
	}
	conexionesV += "}\n"
	for countPos < len(vectorDatos) {
		if countPos == 5*len(indices)+5*len(indices)*countColumna {
			countColumna++
		}
		var calificacion int = 0
		countFila = 0
		for i := 0; i < 10; i++ {
			if countPos == len(vectorDatos)-5 {
				break
			}
			calificacion++
			if i == 5 {
				countFila++
				calificacion = 1
			}
			posiciones += "struct" + strconv.Itoa(countPos+i) + "[shape=Mrecord,color" +
				"=blue, label=\"" + indices[countFila] + "|" + nombresDep[countColumna] +
				"|{Pos: " + strconv.Itoa(countPos+i) + "|Calif.: " + strconv.Itoa(calificacion) +
				"*}\"];\n"
			if vectorDatos[countPos+i].Size > 0 {
				var aux *estructuras.Nodo = vectorDatos[countPos+i].First
				conexionesL += "struct" + strconv.Itoa(countPos+i) + "->nodo" + strconv.Itoa(countList) +
					"[arrowhead=dot, color=\"#b8002b\"];\n"
				listas += "subgraph cluster" + strconv.Itoa(numCluster) + "{\nstyle=invis;\nedge[dir=both]\n"
				for j := 0; j < vectorDatos[countPos+i].Size; j++ {
					listas += "nodo" + strconv.Itoa(countList) + "[shape=Mrecord, color=" +
						"\"#00bf0d\",label=\"{{" + strconv.Itoa(estructuras.GetAscii(aux.GetDatoString())) + "|" +
						aux.Contenido.(*estructuras.NodoTienda).Tienda.Nombre + "}|" + aux.Contenido.(*estructuras.NodoTienda).Tienda.Descripcion + "}\"];\n"
					if j != vectorDatos[countPos+i].Size-1 {
						aux = aux.Next
					}

					if j >= 1 {
						conexionesL += "nodo" + strconv.Itoa(countList-1) + "->" +
							"nodo" + strconv.Itoa(countList) + "[arrowhead=rvee, color=orange];\n" +
							"nodo" + strconv.Itoa(countList) + "->" +
							"nodo" + strconv.Itoa(countList-1) + "[arrowhead=rvee, color=yellow];\n"
					}
					countList++
				}
				numCluster++
				listas += "}\n"
			}
		}
		countPos += 10
	}
	conexionesL += "}"
	data := []byte(posiciones + conexionesV + listas + conexionesL)
	_ = ioutil.WriteFile("Grafica.dot", data, 0644)
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "Grafica.dot").Output()
	_ = ioutil.WriteFile("Grafica.pdf", cmd, os.FileMode(0777))

	fmt.Fprintf(w, "Ya_Está_La_Gráfica")

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
				var nodo *estructuras.Nodo = estructuras.NewNodo(estructuras.NewNodoTienda(ms.Datos[j].Departamentos[i].Tiendas[k]))
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
	//router.HandleFunc("/CargarPedidos", cargarPedidos).Methods("POST")
	router.HandleFunc("/CargarInventarios", cargarInventarios).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("POST")
	router.HandleFunc("/id/{numero}", id).Methods("GET")
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
