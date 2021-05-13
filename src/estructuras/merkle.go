package estructuras

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoMerkle struct {
	valor     string
	contenido string
	derecha   *NodoMerkle
	izquierda *NodoMerkle
}

type ArbolMerkle struct {
	raiz *NodoMerkle
}

func newNodoMerkle(valor string, contenido string, derecha *NodoMerkle, izquierda *NodoMerkle) *NodoMerkle {
	return &NodoMerkle{valor, contenido, derecha, izquierda}
}

func NewArbolMerkle() *ArbolMerkle {
	return &ArbolMerkle{}
}

func (this *ArbolMerkle) Insertar(valor string) {
	n := newNodoMerkle(valor, fmt.Sprintf("%x", sha256.Sum256([]byte(valor))), nil, nil)

	if this.raiz == nil {
		list := list.New()

		list.PushBack(n)
		list.PushBack(newNodoMerkle(strconv.Itoa(-1), fmt.Sprintf("%x", sha256.Sum256([]byte("-1"))), nil, nil))
		this.construirArbol(list)

	} else {
		lis := this.ObtenerLista()

		lis.PushBack(n)
		this.construirArbol(lis)
	}
}

func (this *ArbolMerkle) ObtenerLista() *list.List {
	lis := list.New()
	obtenerLista(lis, this.raiz.izquierda)
	obtenerLista(lis, this.raiz.derecha)
	return lis

}

func obtenerLista(lista *list.List, actual *NodoMerkle) {
	if actual != nil {
		obtenerLista(lista, actual.izquierda)
		if actual.derecha == nil && actual.valor != strconv.Itoa(-1) {
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.derecha)

	}
}

func (this *ArbolMerkle) construirArbol(lista *list.List) {
	size := float64(lista.Len())

	cantmerkle := 1

	for (size / 2) > 1 {
		cantmerkle++
		size = size / 2
	}

	nodostot := math.Pow(2, float64(cantmerkle))

	for lista.Len() < int(nodostot) {
		lista.PushBack(newNodoMerkle(strconv.Itoa(-1), fmt.Sprintf("%x", sha256.Sum256([]byte("-1"))), nil, nil))
	}

	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)

		nodo1 := primero.Value.(*NodoMerkle)

		nodo2 := segundo.Value.(*NodoMerkle)

		concatenado := nodo1.contenido + nodo2.contenido
		nuevo := newNodoMerkle(concatenado, fmt.Sprintf("%x", sha256.Sum256([]byte(concatenado))), nodo2, nodo1)

		lista.PushBack(nuevo)

	}

	this.raiz = lista.Front().Value.(*NodoMerkle)

}

func (this *ArbolMerkle) Codigo() string {
	var cadena strings.Builder

	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if this.raiz != nil {
		valor := this.raiz.valor
		contenido := this.raiz.contenido
		if len(valor) > 15 {
			valor = valor[0:15] + "..."
		}
		if len(contenido) > 15 {
			contenido = contenido[0:15] + "..."
		}
		fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\", color=brown, style=filled, fillcolor=green];\n", &(*this.raiz), contenido, valor)

		this.generar(&cadena, (this.raiz), this.raiz.izquierda, true)
		this.generar(&cadena, this.raiz, this.raiz.derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	return cadena.String()
}

func (this *ArbolMerkle) generar(cadena *strings.Builder, padre *NodoMerkle, actual *NodoMerkle, izquierda bool) {
	if actual != nil {
		valor := actual.valor
		contenido := actual.contenido
		if len(valor) > 15 {
			valor = valor[0:15] + "..."
		}
		if len(contenido) > 15 {
			contenido = contenido[0:15] + "..."
		}
		fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\", color=brown, style=filled, fillcolor=green];\n", &(*actual), contenido, valor)

		if izquierda {
			fmt.Fprintf(cadena, "node%p:f0->node%p:f1[arrowhead=rvee, color=orange]\n", &(*padre), &(*actual))
		} else {
			fmt.Fprintf(cadena, "node%p:f2->node%p:f1[arrowhead=lvee, color=orange]\n", &(*padre), &(*actual))
		}

		this.generar(cadena, actual, actual.izquierda, true)
		this.generar(cadena, actual, actual.derecha, false)
	}
}
