package estructuras

import (
	"fmt"
	"strconv"
)

type Cola struct {
	Frente *Nodo
}

func NewCola() *Cola {
	return &Cola{Frente: nil}
}

func (cola *Cola) Queue(dato interface{}) {
	nodo := &Nodo{
		Contenido: dato,
		Next:      nil,
	}
	if cola.Frente == nil {
		cola.Frente = nodo
	} else {
		aux := cola.Frente
		for aux.Next != nil {
			aux = aux.Next
		}
		aux.Next = nodo
	}
}

func (cola *Cola) Dequeue() *Nodo {
	if cola.Frente == nil {
		return nil
	}
	nodo := cola.Frente
	cola.Frente = cola.Frente.Next
	return nodo
}

func (cola *Cola) GraficarPedidos(arbol *ArbolB) string {
	var listas, nodos, conexionesC, conexionesP string
	var numCluster int = 1
	nodos = "digraph G{\ncompound=true;\nsubgraph cluster0{" +
		"style=invis;\nedge[minlen=0.1, dir=fordware]\n"
	nodos += "inicio[shape=Mrecord,color=blue, label=\"PEDIDOS\\n" + cola.Frente.Contenido.(*Pedido).Fecha + "\\n" + cola.Frente.Contenido.(*Pedido).Departamento + "\"]\n"
	conexionesC = "inicio->" + fmt.Sprintf("nodo%p", cola.Frente) + "[arrowhead=vee, color=\"#9100d4\"]\n"
	aux := cola.Frente
	listas = ""
	conexionesP = ""
	for aux != nil {
		nodos += fmt.Sprintf("nodo%p", aux) + "[shape=Mrecord, color=blue,label=\"{"
		cliente := arbol.Buscar(aux.Contenido.(*Pedido).Cliente)
		if cliente != nil {
			nodos += "{" + strconv.Itoa(cliente.(*Usuario).Dpi) + "|" + cliente.(*Usuario).Nombre + "}"
		} else {
			nodos += "Anónimo"
		}
		nodos += "|{" + aux.Contenido.(*Pedido).Tienda + "|" + "Calif.: " + strconv.Itoa(aux.Contenido.(*Pedido).Calificacion) + "*}}\"]\n"
		conexionesP += fmt.Sprintf("nodo%p", aux) + "->" + fmt.Sprintf("nodo%p", aux.Contenido.(*Pedido).Productos[0]) + "[arrowhead=dot, color=\"#b8002b\"]\n"
		listas += "subgraph cluster" + strconv.Itoa(numCluster) + "{\nstyle=invis\n"
		numCluster++
		for i := 0; i < len(aux.Contenido.(*Pedido).Productos); i++ {
			listas += fmt.Sprintf("nodo%p", aux.Contenido.(*Pedido).Productos[i]) + "[shape=Mrecord, color=\"#00bf0d\", label=\"" + strconv.Itoa(aux.Contenido.(*Pedido).Productos[i].Codigo) + "\"]\n"
			if i+1 < len(aux.Contenido.(*Pedido).Productos) {
				listas += fmt.Sprintf("nodo%p", aux.Contenido.(*Pedido).Productos[i]) + "->" + fmt.Sprintf("nodo%p", aux.Contenido.(*Pedido).Productos[i+1]) + "[arrowhead=box, color=orange]\n"
			}
		}
		listas += "}\n"
		if aux.Next != nil {
			conexionesC += fmt.Sprintf("nodo%p", aux) + "->" + fmt.Sprintf("nodo%p", aux.Next) + "[arrowhead=vee, color=\"#9100d4\"]\n"
		}
		aux = aux.Next
	}
	return nodos + conexionesC + "}\n" + listas + conexionesP + "}"
}
