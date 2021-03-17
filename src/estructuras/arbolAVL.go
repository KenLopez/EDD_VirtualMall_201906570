package estructuras

import (
	"fmt"
	"reflect"
	"strconv"
)

type NodoArbol struct {
	Izq, Der  *NodoArbol
	Peso      int
	Dato      int
	Contenido interface{}
}

type Arbol struct {
	raiz *NodoArbol
}

func NewNodoArbol(contenido interface{}, dato int) *NodoArbol {
	return &NodoArbol{
		Izq:       nil,
		Der:       nil,
		Peso:      0,
		Dato:      dato,
		Contenido: contenido,
	}
}

func NewArbol() *Arbol {
	return &Arbol{
		raiz: nil,
	}
}

func rotarII(n *NodoArbol) *NodoArbol {
	n1 := n.Izq
	n.Izq = n1.Der
	n1.Der = n
	if n1.Peso == -1 {
		n.Peso = 0
		n1.Peso = 0
	} else {
		n.Peso = -1
		n1.Peso = 1
	}
	return n1
}

func rotarDD(n *NodoArbol) *NodoArbol {
	n1 := n.Der
	n.Der = n1.Izq
	n1.Izq = n
	if n1.Peso == 1 {
		n.Peso = 0
		n1.Peso = 0
	} else {
		n.Peso = 1
		n1.Peso = -1
	}
	return n1
}

func rotarDI(n *NodoArbol) *NodoArbol {
	n1 := n.Der
	n2 := n1.Izq
	n2.Izq = n
	n1.Izq = n2.Der
	n2.Der = n1
	if n2.Peso == 1 {
		n.Peso = -1
	} else {
		n.Peso = 0
	}
	if n2.Peso == -1 {
		n1.Peso = 1
	} else {
		n1.Peso = 0
	}
	n2.Peso = 0
	return n2
}

func rotarID(n *NodoArbol) *NodoArbol {
	n1 := n.Izq
	n2 := n1.Der
	n1.Izq = n2.Der
	n2.Der = n
	n1.Der = n2.Izq
	n2.Izq = n1
	if n2.Peso == 1 {
		n.Peso = -1
	} else {
		n.Peso = 0
	}
	if n2.Peso == -1 {
		n1.Peso = 1
	} else {
		n1.Peso = 0
	}
	n2.Peso = 0
	return n2
}

func insertar(raiz *NodoArbol, dato int, contenido interface{}, hc *bool) *NodoArbol {
	var n1 *NodoArbol
	if raiz == nil {
		raiz = NewNodoArbol(contenido, dato)
		*hc = true
	} else if dato < raiz.Dato {
		izq := insertar(raiz.Izq, dato, contenido, hc)
		raiz.Izq = izq
		if *hc {
			switch raiz.Peso {
			case 1:
				raiz.Peso = 0
				*hc = false
				break
			case 0:
				raiz.Peso = -1
				break
			case -1:
				n1 = raiz.Izq
				if n1.Peso == -1 {
					raiz = rotarII(raiz)
				} else {
					raiz = rotarID(raiz)
				}
				*hc = false
			}
		}
	} else if dato > raiz.Dato {
		der := insertar(raiz.Der, dato, contenido, hc)
		raiz.Der = der
		if *hc {
			switch raiz.Peso {
			case 1:
				n1 := raiz.Der
				if n1.Peso == 1 {
					raiz = rotarDD(raiz)
				} else {
					raiz = rotarDI(raiz)
				}
				*hc = false
				break
			case 0:
				raiz.Peso = 1
				break
			case -1:
				raiz.Peso = 0
				*hc = false
				break
			}
		}
	}
	return raiz
}

func (arbol *Arbol) Insertar(contenido interface{}, dato int) {
	b := false
	nodo := arbol.Buscar(dato)
	if nodo == nil {
		arbol.raiz = insertar(arbol.raiz, dato, contenido, &b)
	} else {
		if reflect.TypeOf(nodo.Contenido).String() == "*estructuras.Producto" {
			//nodo.contenido.(*Producto).Cantidad += contenido.(*Producto).Cantidad
			fmt.Println(nodo.Contenido.(*Producto).Nombre + ": " + strconv.Itoa(nodo.Contenido.(*Producto).Cantidad))
		}
	}

}

func (arbol *Arbol) Buscar(dato int) *NodoArbol {
	if arbol.raiz == nil {
		return nil
	}
	return buscar(arbol.raiz, dato)
}

func buscar(nodo *NodoArbol, dato int) *NodoArbol {
	if nodo == nil || nodo.Dato == dato {
		return nodo
	}
	if dato < nodo.Dato {
		return buscar(nodo.Izq, dato)
	}
	return buscar(nodo.Der, dato)
}
