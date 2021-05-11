package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"./estructuras"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var indices, nombresDep []string
var vectorDatos []*estructuras.Lista
var arbolAnios *estructuras.Arbol
var arbolCuentas *estructuras.ArbolB = estructuras.NewArbolB(5)
var key string = "132115"
var grafo *estructuras.Grafo

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
	json.NewEncoder(w).Encode("Tiendas Cargadas")
}

func login(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.UserLogin
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error", Content: "No se pudo iniciar sesión."})
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	cuenta := arbolCuentas.Buscar(ms.Dpi)
	if cuenta != nil {
		if string(cuenta.(*estructuras.Usuario).Password) == string(ms.Password) {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: cuenta.(*estructuras.Usuario).Cuenta})
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error", Content: "No se pudo iniciar sesión. Credenciales incorrectas."})
		}
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error", Content: "No se pudo iniciar sesión. Credenciales incorrectas."})
	}
}

func eliminarCuenta(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.UserLogin
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error", Content: "No se pudo iniciar sesión."})
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	cuenta := arbolCuentas.Buscar(ms.Dpi)
	if cuenta != nil {
		if string(cuenta.(*estructuras.Usuario).Password) == string(ms.Password) {
			arbolCuentas.Eliminar(ms.Dpi)
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok"})
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	}
}

func registro(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.Usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error", Content: "No se pudo registrar."})
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	if arbolCuentas.Buscar(ms.Dpi) == nil {
		ms.Cuenta = "Cliente"
		arbolCuentas.Insertar(estructuras.NewKey(ms.Dpi, ms))
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok"})
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	}
}

func cargarUsuarios(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.ArchivoUsuarios
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_ :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.Usuarios); i++ {
		ms.Usuarios[i].Password = fmt.Sprintf("%x", sha256.Sum256([]byte(ms.Usuarios[i].Password)))
		arbolCuentas.Insertar(estructuras.NewKey(ms.Usuarios[i].Dpi, ms.Usuarios[i]))
	}
	json.NewEncoder(w).Encode("Usuarios Cargados")
}

func cargarGrafo(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.ArchivoGrafo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("No Jaló :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	grafo = estructuras.NewGrafo(ms)
	json.NewEncoder(w).Encode("Grafo Cargado")
}

func cargarPedidos(w http.ResponseWriter, r *http.Request) {
	var ms *estructuras.ArchivoPedido
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_ :c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &ms)
	for i := 0; i < len(ms.Pedidos); i++ {
		nodo := buscarPosicion(&estructuras.RequestFind{
			Departamento: ms.Pedidos[i].Departamento,
			Nombre:       ms.Pedidos[i].Tienda,
			Calificacion: ms.Pedidos[i].Calificacion,
		})
		destinos := estructuras.NewCola()
		if nodo == nil {
			fmt.Fprintln(w, "No_se_encontró_tienda:"+ms.Pedidos[i].Tienda+"-;")
		} else if nodo.Contenido.(*estructuras.NodoTienda).Inventario == nil {
			fmt.Fprintln(w, "La_tienda:"+ms.Pedidos[i].Tienda+" no_posee_inventario-;")
		} else {
			var prodOk []*estructuras.Codigo
			for j := 0; j < len(ms.Pedidos[i].Productos); j++ {
				nodoArbol := nodo.Contenido.(*estructuras.NodoTienda).Inventario.Buscar(ms.Pedidos[i].Productos[j].Codigo)
				if nodoArbol != nil {
					vertice := grafo.VerticeExists(nodoArbol.Contenido.(*estructuras.Producto).Almacenamiento)
					if vertice != nil {
						if destinos.Frente == nil {
							destinos.Queue(vertice)
						} else {
							aux := destinos.Frente
							insertar := true
							for aux != nil {
								if aux.Contenido.(*estructuras.Vertice).Nombre == vertice.Nombre {
									insertar = false
									break
								} else if aux.Contenido.(*estructuras.Vertice).Nombre == grafo.Inicio.Nombre {
									insertar = false
									break
								} else if aux.Contenido.(*estructuras.Vertice).Nombre == grafo.Entrega.Nombre {
									insertar = false
									break
								}
								aux = aux.Next
							}
							if insertar {
								destinos.Queue(vertice)
							}
						}
						nodoArbol.Contenido.(*estructuras.Producto).Cantidad--
						prodOk = append(prodOk, &estructuras.Codigo{Codigo: ms.Pedidos[i].Productos[j].Codigo})
					}
				}
			}
			if len(prodOk) != 0 {
				destinos.Queue(grafo.Entrega)
				destinos.Queue(grafo.Inicio)
				go grafo.RecorridoRobot(ms.Pedidos[i], destinos)
				ms.Pedidos[i].Productos = prodOk
				if arbolAnios == nil {
					arbolAnios = estructuras.NewArbol()
				}
				if arbolAnios.Buscar(estructuras.GetAnio(ms.Pedidos[i].Fecha)) == nil {
					arbolAnios.Insertar(estructuras.NewArbol(), estructuras.GetAnio(ms.Pedidos[i].Fecha))
				}
				nodoAnio := arbolAnios.Buscar(estructuras.GetAnio(ms.Pedidos[i].Fecha))
				if nodoAnio.Contenido.(*estructuras.Arbol).Buscar(estructuras.GetMes(ms.Pedidos[i].Fecha)) == nil {
					nodoAnio.Contenido.(*estructuras.Arbol).Insertar(estructuras.NewMatriz(), estructuras.GetMes(ms.Pedidos[i].Fecha))
				}
				nodoMes := nodoAnio.Contenido.(*estructuras.Arbol).Buscar(estructuras.GetMes(ms.Pedidos[i].Fecha))
				nodoMes.Contenido.(*estructuras.Matriz).NuevoPedido(ms.Pedidos[i])
			}
		}
	}
	json.NewEncoder(w).Encode("Pedidos Cargados")
}

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
	json.NewEncoder(w).Encode("Inventarios Cargados")
}

func tiendaEspecifica(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var nodo *estructuras.NodoLista = buscarPosicion(busqueda)
	if nodo == nil {
		fmt.Fprintln(w, "No_se_encontró_la_tienda_solicitada")
		json.NewEncoder(w).Encode(nil)
	} else {
		var tienda *estructuras.Tienda = nodo.Contenido.(*estructuras.NodoTienda).Tienda
		json.NewEncoder(w).Encode(tienda)
	}
}

func buscarPosicion(request *estructuras.RequestFind) *estructuras.NodoLista {
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
	var s int = columna*len(indices) + fila
	//fmt.Print(columna)
	//fmt.Print(s)
	//fmt.Print(s*5 + calificacion - 1)
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
		var aux *estructuras.NodoLista = vectorDatos[numero].First
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

func getTiendas(w http.ResponseWriter, r *http.Request) {
	var file estructuras.Archivo = estructuras.Archivo{}
	if len(vectorDatos) == 0 {
		file.Datos = make([]*estructuras.Dato, 0)
	} else {
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
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(file)
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
	if len(vectorDatos) > 0 {
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
					var aux *estructuras.NodoLista = vectorDatos[countPos+i].First
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
		cmd, _ := exec.Command(path, "-Tpng", "Grafica.dot").Output()
		_ = ioutil.WriteFile("Grafica.png", cmd, os.FileMode(0777))
		e := os.Remove("Grafica.dot")
		if e != nil {
			log.Fatal(e)
		}
		f, _ := os.Open("Grafica.png")
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	}

}

func getArbolAnio(w http.ResponseWriter, r *http.Request) {
	if arbolAnios != nil {
		data := []byte(arbolAnios.Graficar(false))
		_ = ioutil.WriteFile("Arbol-Anios.dot", data, 0644)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", "Arbol-Anios.dot").Output()
		_ = ioutil.WriteFile("Arbol-Anios.png", cmd, os.FileMode(0777))
		e := os.Remove("Arbol-Anios.dot")
		if e != nil {
			log.Fatal(e)
		}
		f, _ := os.Open("Arbol-Anios.png")
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	}

}

func getArbolCuentas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cifrado, err := strconv.Atoi(vars["cifrado"])
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		title := "Arbol-Cuentas"
		switch cifrado {
		case 0:
			title += "-(Original)"
			break
		case 1:
			title += "(Cifrado-Sensible)"
		case 2:
			title += "(Cifrado)"
		}
		data := []byte(arbolCuentas.Graficar(cifrado, key))
		_ = ioutil.WriteFile(title+".dot", data, 0644)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpng", title+".dot").Output()
		_ = ioutil.WriteFile(title+".png", cmd, os.FileMode(0777))
		e := os.Remove(title + ".dot")
		if e != nil {
			log.Fatal(e)
		}
		f, _ := os.Open(title + ".png")
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
	}
}

func updateKey(w http.ResponseWriter, r *http.Request) {
	var ms struct {
		Key string `json:Key`
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		json.Unmarshal(reqBody, &ms)
		key = ms.Key
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: ms.Key})
	}
}

func getGrafo(w http.ResponseWriter, r *http.Request) {
	if grafo == nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		title := "Grafo"
		data := []byte(grafo.Graficar())
		_ = ioutil.WriteFile(title+".dot", data, 0644)
		path, _ := exec.LookPath("neato")
		cmd, _ := exec.Command(path, "-Tpng", title+".dot").Output()
		_ = ioutil.WriteFile(title+".png", cmd, os.FileMode(0777))
		e := os.Remove(title + ".dot")
		if e != nil {
			log.Fatal(e)
		}
		f, _ := os.Open(title + ".png")
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
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
				var nodo *estructuras.NodoLista = estructuras.NewNodo(estructuras.NewNodoTienda(ms.Datos[j].Departamentos[i].Tiendas[k]))
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

func getArbolMeses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numero, err := strconv.Atoi(vars["anio"])
	if err != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else if arbolAnios == nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		a := arbolAnios.Buscar(numero)
		arbol := a.Contenido.(*estructuras.Arbol)
		if arbol != nil {
			data := []byte(arbol.Graficar(true))
			_ = ioutil.WriteFile("Arbol-Meses-"+strconv.Itoa(numero)+".dot", data, 0644)
			path, _ := exec.LookPath("dot")
			cmd, _ := exec.Command(path, "-Tpng", "Arbol-Meses-"+strconv.Itoa(numero)+".dot").Output()
			_ = ioutil.WriteFile("Arbol-Meses-"+strconv.Itoa(numero)+".png", cmd, os.FileMode(0777))
			e := os.Remove("Arbol-Meses-" + strconv.Itoa(numero) + ".dot")
			if e != nil {
				log.Fatal(e)
			}
			f, _ := os.Open("Arbol-Meses-" + strconv.Itoa(numero) + ".png")
			reader := bufio.NewReader(f)
			content, _ := ioutil.ReadAll(reader)
			encoded := base64.StdEncoding.EncodeToString(content)
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}
	}
}

func getMatriz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anio, err1 := strconv.Atoi(vars["anio"])
	mes, err2 := strconv.Atoi(vars["mes"])
	if err1 != nil || err2 != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else if arbolAnios != nil {
		arbolM := arbolAnios.Buscar(anio).Contenido.(*estructuras.Arbol)
		if arbolM != nil {
			nodoM := arbolM.Buscar(mes)
			if nodoM != nil {
				data := []byte(nodoM.Contenido.(*estructuras.Matriz).Graficar(estructuras.GetMesName(nodoM.Dato)))
				_ = ioutil.WriteFile("Pedidos-"+estructuras.GetMesName(nodoM.Dato)+"-"+strconv.Itoa(anio)+".dot", data, 0644)
				path, _ := exec.LookPath("dot")
				cmd, _ := exec.Command(path, "-Tpng", "Pedidos-"+estructuras.GetMesName(nodoM.Dato)+"-"+strconv.Itoa(anio)+".dot").Output()
				_ = ioutil.WriteFile("Pedidos-"+estructuras.GetMesName(nodoM.Dato)+"-"+strconv.Itoa(anio)+".png", cmd, os.FileMode(0777))
				e := os.Remove("Pedidos-" + estructuras.GetMesName(nodoM.Dato) + "-" + strconv.Itoa(anio) + ".dot")
				if e != nil {
					log.Fatal(e)
				}
				f, _ := os.Open("Pedidos-" + estructuras.GetMesName(nodoM.Dato) + "-" + strconv.Itoa(anio) + ".png")
				reader := bufio.NewReader(f)
				content, _ := ioutil.ReadAll(reader)
				encoded := base64.StdEncoding.EncodeToString(content)
				json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
			} else {
				json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
			}
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}
	} else {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	}
}

func getRobot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anio, err1 := strconv.Atoi(vars["anio"])
	mes, err2 := strconv.Atoi(vars["mes"])
	categoria := vars["categoria"]
	dia, err4 := strconv.Atoi(vars["dia"])
	num, err5 := strconv.Atoi(vars["num"])
	if err1 != nil || err2 != nil || err4 != nil || err5 != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else if arbolAnios == nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		arbolM := arbolAnios.Buscar(anio).Contenido.(*estructuras.Arbol)
		if arbolM != nil {
			nodoM := arbolM.Buscar(mes)
			if nodoM != nil {
				cola := nodoM.Contenido.(*estructuras.Matriz).Get(dia, categoria).Dato
				if cola != nil {
					pedido := cola.Get(num)
					if pedido != nil {
						arr := pedido.Contenido.(*estructuras.Pedido).CaminoCorto.MovToArray()
						json.NewEncoder(w).Encode(arr)
					} else {
						json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
					}
				} else {
					json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
				}
			} else {
				json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
			}
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}
	}
}

func getPedidosDia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anio, err1 := strconv.Atoi(vars["anio"])
	mes, err2 := strconv.Atoi(vars["mes"])
	categoria := vars["categoria"]
	dia, err4 := strconv.Atoi(vars["dia"])
	if err1 != nil || err2 != nil || err4 != nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else if arbolAnios == nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
	} else {
		arbolM := arbolAnios.Buscar(anio).Contenido.(*estructuras.Arbol)
		if arbolM != nil {
			nodoM := arbolM.Buscar(mes)
			if nodoM != nil {
				cola := nodoM.Contenido.(*estructuras.Matriz).Get(dia, categoria).Dato
				if cola != nil {
					data := []byte(cola.GraficarPedidos(arbolCuentas))
					_ = ioutil.WriteFile("Pedidos-"+categoria+"-"+strconv.Itoa(dia)+strconv.Itoa(mes)+strconv.Itoa(anio)+".dot", data, 0644)
					path, _ := exec.LookPath("dot")
					cmd, _ := exec.Command(path, "-Tpng", "Pedidos-"+categoria+"-"+strconv.Itoa(dia)+strconv.Itoa(mes)+strconv.Itoa(anio)+".dot").Output()
					_ = ioutil.WriteFile("Pedidos-"+categoria+"-"+strconv.Itoa(dia)+strconv.Itoa(mes)+strconv.Itoa(anio)+".png", cmd, os.FileMode(0777))
					e := os.Remove("Pedidos-" + categoria + "-" + strconv.Itoa(dia) + strconv.Itoa(mes) + strconv.Itoa(anio) + ".dot")
					if e != nil {
						log.Fatal(e)
					}
					f, _ := os.Open("Pedidos-" + categoria + "-" + strconv.Itoa(dia) + strconv.Itoa(mes) + strconv.Itoa(anio) + ".png")
					reader := bufio.NewReader(f)
					content, _ := ioutil.ReadAll(reader)
					encoded := base64.StdEncoding.EncodeToString(content)
					json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
				} else {
					json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
				}
			} else {
				json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
			}
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}
	}
}

func getInventario(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(nil)
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var nodo *estructuras.NodoLista = buscarPosicion(busqueda)
	if nodo == nil {
		fmt.Fprintln(w, "No_se_encontró_la_tienda_solicitada")
		json.NewEncoder(w).Encode(nil)
	} else {
		var inventario []*estructuras.Producto
		if nodo.Contenido.(*estructuras.NodoTienda).Inventario != nil {
			inventario = nodo.Contenido.(*estructuras.NodoTienda).Inventario.ToArrayProductos()
		} else {
			inventario = make([]*estructuras.Producto, 0)
		}
		res := struct {
			Descripcion string
			Contacto    string
			Productos   []*estructuras.Producto
		}{
			Descripcion: nodo.Contenido.(*estructuras.NodoTienda).Tienda.Descripcion,
			Contacto:    nodo.Contenido.(*estructuras.NodoTienda).Tienda.Contacto,
			Productos:   inventario,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func getArbolInventario(w http.ResponseWriter, r *http.Request) {
	var busqueda *estructuras.RequestFind
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No_Jaló_:c")
	}
	w.Header().Set("Content-Type", "application/json")
	json.Unmarshal(reqBody, &busqueda)
	var nodo *estructuras.NodoLista = buscarPosicion(busqueda)
	if nodo == nil {
		json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok"})
	} else {
		if nodo.Contenido.(*estructuras.NodoTienda).Inventario != nil {
			title := "Inventario-" + busqueda.Departamento + "-" + strings.ReplaceAll(busqueda.Nombre, " ", "_")
			data := []byte(nodo.Contenido.(*estructuras.NodoTienda).Inventario.Graficar(false))
			_ = ioutil.WriteFile(title+".dot", data, 0644)
			path, _ := exec.LookPath("dot")
			cmd, _ := exec.Command(path, "-Tpng", title+".dot").Output()
			_ = ioutil.WriteFile(title+".png", cmd, os.FileMode(0777))
			e := os.Remove(title + ".dot")
			if e != nil {
				log.Fatal(e)
			}
			f, _ := os.Open(title + ".png")
			reader := bufio.NewReader(f)
			content, _ := ioutil.ReadAll(reader)
			encoded := base64.StdEncoding.EncodeToString(content)
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Ok", Content: encoded})
		} else {
			json.NewEncoder(w).Encode(estructuras.Response{Tipo: "Error"})
		}

	}
}

func main() {
	pred := &estructuras.Usuario{
		Dpi:      1234567890101,
		Nombre:   "EDD2021",
		Correo:   "auxiliar@edd.com",
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte("1234"))),
		Cuenta:   "Admin",
	}
	pred2 := &estructuras.Usuario{
		Dpi:      3004966640101,
		Nombre:   "Kenneth López",
		Correo:   "khlopez2000@gmail.com",
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte("1230"))),
		Cuenta:   "Usuario",
	}
	arbolCuentas.Insertar(estructuras.NewKey(pred.Dpi, pred))
	arbolCuentas.Insertar(estructuras.NewKey(pred2.Dpi, pred2))
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	router.HandleFunc("/CargarPedidos", cargarPedidos).Methods("POST")
	router.HandleFunc("/CargarInventarios", cargarInventarios).Methods("POST")
	router.HandleFunc("/CargarUsuarios", cargarUsuarios).Methods("POST")
	router.HandleFunc("/CargarGrafo", cargarGrafo).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("GET")
	router.HandleFunc("/id/{numero}", id).Methods("GET")
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/EliminarCuenta", eliminarCuenta).Methods("DELETE")
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getTiendas", getTiendas).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/GetArbolAnio", getArbolAnio).Methods("GET")
	router.HandleFunc("/GetArbolMeses/{anio}", getArbolMeses).Methods("GET")
	router.HandleFunc("/GetMatriz/{anio}/{mes}", getMatriz).Methods("GET")
	router.HandleFunc("/GetPedidosDia/{anio}/{mes}/{categoria}/{dia}", getPedidosDia).Methods("GET")
	router.HandleFunc("/GetRobot/{anio}/{mes}/{categoria}/{dia}/{num}", getRobot).Methods("GET")
	router.HandleFunc("/GetArbolInventario", getArbolInventario).Methods("POST")
	router.HandleFunc("/GetArbolCuentas/{cifrado}", getArbolCuentas).Methods("GET")
	router.HandleFunc("/GetInventario", getInventario).Methods("POST")
	router.HandleFunc("/GetGrafo", getGrafo).Methods("GET")
	router.HandleFunc("/Login", login).Methods("POST")
	router.HandleFunc("/UpdateKey", updateKey).Methods("POST")
	router.HandleFunc("/Registro", registro).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "DELETE"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
