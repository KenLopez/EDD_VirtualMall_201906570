package estructuras

import "fmt"

type Grafo struct {
	Nodos     *Lista
	Recorrido *Lista
	Inicio    *Vertice
	Entrega   *Vertice
}

type Vertice struct {
	Nombre  string
	Enlaces *Lista
}

type Enlace struct {
	Peso    float32
	Destino *Vertice
}

type Movimiento struct {
	Origen  *Vertice
	Peso    float32
	Destino *Vertice
}

func NewMovimiento(origen *Vertice, peso float32, destino *Vertice) *Movimiento {
	return &Movimiento{
		Origen:  origen,
		Peso:    peso,
		Destino: destino,
	}
}

func NewVertice(nombre string) *Vertice {
	return &Vertice{
		Nombre:  nombre,
		Enlaces: NewLista(),
	}
}

func NewEnlace(peso float32, destino *Vertice) *Enlace {
	return &Enlace{
		Peso:    peso,
		Destino: destino,
	}
}

func (g *Grafo) VerticeExists(nombre string) *Vertice {
	if g.Nodos.First != nil {
		aux := g.Nodos.First
		for aux != nil {
			if aux.Contenido.(*Vertice).Nombre == nombre {
				return aux.Contenido.(*Vertice)
			}
			aux = aux.Next
		}
	}
	return nil
}

func EnlaceExists(enlaces *Lista, enlace *Enlace) *Enlace {
	if enlaces.First != nil {
		aux := enlaces.First
		for aux != nil {
			if aux.Contenido.(*Enlace).Peso == enlace.Peso && aux.Contenido.(*Enlace).Destino == enlace.Destino {
				return aux.Contenido.(*Enlace)
			}
			aux = aux.Next
		}
	}
	return nil
}

func NewGrafo(archivo *ArchivoGrafo) *Grafo {
	graph := &Grafo{
		Nodos:     NewLista(),
		Recorrido: NewLista(),
	}
	for i := 0; i < len(archivo.Nodos); i++ {
		if graph.VerticeExists(archivo.Nodos[i].Nombre) == nil {
			graph.Nodos.InsertarSimple(NewNodo(NewVertice(archivo.Nodos[i].Nombre)))
		}
	}
	for i := 0; i < len(archivo.Nodos); i++ {
		nodo := graph.VerticeExists(archivo.Nodos[i].Nombre)
		for j := 0; j < len(archivo.Nodos[i].Enlaces); j++ {
			tmp := graph.VerticeExists(archivo.Nodos[i].Enlaces[j].Nombre)
			if tmp != nil {
				enlace := NewEnlace(archivo.Nodos[i].Enlaces[j].Distancia, tmp)
				if EnlaceExists(nodo.Enlaces, enlace) == nil {
					nodo.Enlaces.InsertarSimple(NewNodo(NewEnlace(archivo.Nodos[i].Enlaces[j].Distancia, tmp)))
					tmp.Enlaces.InsertarSimple(NewNodo(NewEnlace(archivo.Nodos[i].Enlaces[j].Distancia, nodo)))
				}
			}
		}
	}
	graph.Inicio = graph.VerticeExists(archivo.PosicionInicialRobot)
	graph.Entrega = graph.VerticeExists(archivo.Entrega)
	return graph
}

func MovimientoExists(lista *Lista, origen string, peso float32, destino string) *Movimiento {
	aux := lista.First
	for aux != nil {
		if aux.Contenido.(*Movimiento).Origen.Nombre == origen && aux.Contenido.(*Movimiento).Peso == peso &&
			aux.Contenido.(*Movimiento).Destino.Nombre == destino {
			return aux.Contenido.(*Movimiento)
		} else if aux.Contenido.(*Movimiento).Origen.Nombre == destino && aux.Contenido.(*Movimiento).Peso == peso &&
			aux.Contenido.(*Movimiento).Destino.Nombre == origen {
			return aux.Contenido.(*Movimiento)
		}
		aux = aux.Next
	}
	return nil
}

func (grafo *Grafo) Graficar() string {
	nodos := ""
	conexiones := ""
	enlaces := NewLista()
	aux := grafo.Nodos.First
	extra := ""
	for aux != nil {
		if aux.Contenido.(*Vertice) == grafo.Inicio {
			extra = "style=filled, fillcolor=\"#fcba03\", color=\"#8a2301\""
		} else if aux.Contenido.(*Vertice) == grafo.Entrega {
			extra = "style=filled, fillcolor=\"#bb0ec4\""
		} else {
			extra = ""
		}
		nodos += "nodo" + fmt.Sprintf("%p", aux.Contenido.(*Vertice)) + "[shape=oval, label=\"" + aux.Contenido.(*Vertice).Nombre + "\"" + extra + "]\n"
		aux2 := aux.Contenido.(*Vertice).Enlaces.First
		for aux2 != nil {
			if MovimientoExists(enlaces, aux.Contenido.(*Vertice).Nombre, aux2.Contenido.(*Enlace).Peso, aux2.Contenido.(*Enlace).Destino.Nombre) == nil {
				conexiones += "nodo" + fmt.Sprintf("%p", aux.Contenido.(*Vertice)) + "--" + "nodo" + fmt.Sprintf("%p", aux2.Contenido.(*Enlace).Destino) +
					"[label=" + fmt.Sprintf("%.2f", aux2.Contenido.(*Enlace).Peso) + "]\n"
				enlaces.InsertarSimple(NewNodo(NewMovimiento(aux.Contenido.(*Vertice), aux2.Contenido.(*Enlace).Peso, aux2.Contenido.(*Enlace).Destino)))
			}
			aux2 = aux2.Next
		}
		aux = aux.Next
	}

	return "graph G{\nedge[len=2.5]" + nodos + conexiones + "}"
}
