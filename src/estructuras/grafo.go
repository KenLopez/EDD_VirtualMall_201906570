package estructuras

import (
	"fmt"
	"strings"
)

type Grafo struct {
	Nodos   *Lista
	Inicio  *Vertice
	Entrega *Vertice
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

func (grafo *Grafo) RecorridoRobot(pedido *Pedido, destinos *Cola) {
	pedido.CaminoCorto = NewLista()
	aux := destinos.Dequeue()
	actual := grafo.Inicio
	for aux != nil {
		new := grafo.RecorridoCorto(actual, aux.Contenido.(*Vertice))
		if new == nil {
			break
		} else if new.First == nil {
			break
		}
		pedido.CaminoCorto.InsertarSimple(new.First)
		actual = aux.Contenido.(*Vertice)
		aux = destinos.Dequeue()
	}
}

func (grafo *Grafo) RecorridoCorto(inicio *Vertice, destino *Vertice) *Lista {
	movimientos := NewLista()
	visitados := NewLista()
	visitados.InsertarSimple(NewNodo(inicio))
	actual := inicio
	var transicion *Movimiento
	for actual.Nombre != destino.Nombre {
		transicion = grafo.aristaCorta(actual, visitados)
		if transicion == nil {
			return nil
		}
		visitados.InsertarSimple(NewNodo(transicion.Destino))
		movimientos.InsertarSimple(NewNodo(transicion))
		actual = transicion.Destino
	}
	return movimientos
}

func verticeInList(lista *Lista, nombre string) bool {
	if lista.First != nil {
		aux := lista.First
		for aux != nil {
			if aux.Contenido.(*Vertice).Nombre == nombre {
				return true
			}
			aux = aux.Next
		}
		return false
	} else {
		return false
	}
}

func (grafo *Grafo) aristaCorta(inicio *Vertice, visitados *Lista) *Movimiento {
	var best *Movimiento
	aux := inicio.Enlaces.First
	for aux != nil {
		if !verticeInList(visitados, aux.Contenido.(*Enlace).Destino.Nombre) {
			best = NewMovimiento(inicio, aux.Contenido.(*Enlace).Peso, aux.Contenido.(*Enlace).Destino)
			break
		}
		aux = aux.Next
	}
	if best == nil {
		return nil
	}
	aux = inicio.Enlaces.First
	for aux != nil {
		if aux.Contenido.(*Enlace).Peso < best.Peso {
			if !verticeInList(visitados, aux.Contenido.(*Enlace).Destino.Nombre) {
				best.Peso = aux.Contenido.(*Enlace).Peso
				best.Destino = aux.Contenido.(*Enlace).Destino
			}
		}
		aux = aux.Next
	}
	return best
}

func NewGrafo(archivo *ArchivoGrafo) *Grafo {
	graph := &Grafo{
		Nodos: NewLista(),
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
	name := ""
	for aux != nil {
		if aux.Contenido.(*Vertice) == grafo.Inicio {
			extra = "style=filled, fillcolor=\"#fcba03\", color=\"#8a2301\""
			name = aux.Contenido.(*Vertice).Nombre + "\\n" + "(INICIO)"
		} else if aux.Contenido.(*Vertice) == grafo.Entrega {
			extra = "style=filled, fillcolor=\"#dc87e6\", color=\"#bb0ec4\""
			name = aux.Contenido.(*Vertice).Nombre + "\\n" + "(ENTREGA)"
		} else {
			extra = ""
			name = aux.Contenido.(*Vertice).Nombre
		}
		newName := strings.Split(name, " ")
		if len(newName) > 2 {
			name = newName[0]
			for i := 1; i < len(newName); i++ {
				if i%2 == 0 && i > 0 {
					name += "\\n"
				}
				name += " " + newName[i]
			}
		}
		nodos += "nodo" + fmt.Sprintf("%p", aux.Contenido.(*Vertice)) + "[shape=oval, label=\"" + name + "\"" + extra + "]\n"
		aux2 := aux.Contenido.(*Vertice).Enlaces.First
		for aux2 != nil {
			if MovimientoExists(enlaces, aux.Contenido.(*Vertice).Nombre, aux2.Contenido.(*Enlace).Peso, aux2.Contenido.(*Enlace).Destino.Nombre) == nil {
				conexiones += "nodo" + fmt.Sprintf("%p", aux.Contenido.(*Vertice)) + "--" + "nodo" + fmt.Sprintf("%p", aux2.Contenido.(*Enlace).Destino) +
					"[label=" + fmt.Sprintf("%.2f", aux2.Contenido.(*Enlace).Peso) + ", labeldistance=0]\n"
				enlaces.InsertarSimple(NewNodo(NewMovimiento(aux.Contenido.(*Vertice), aux2.Contenido.(*Enlace).Peso, aux2.Contenido.(*Enlace).Destino)))
			}
			aux2 = aux2.Next
		}
		aux = aux.Next
	}

	return "graph G{\nedge[len=2.5]" + nodos + conexiones + "}"
}
