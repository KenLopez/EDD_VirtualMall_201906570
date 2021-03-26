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
	n.Der = n2.Izq
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
	n.Izq = n2.Der
	n2.Der = n
	n1.Der = n2.Izq
	n2.Izq = n1
	if n2.Peso == 1 {
		n1.Peso = -1
	} else {
		n1.Peso = 0
	}
	if n2.Peso == -1 {
		n.Peso = 1
	} else {
		n.Peso = 0
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
	} else if dato >= raiz.Dato {
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
			nodo.Contenido.(*Producto).Cantidad = contenido.(*Producto).Cantidad
			//fmt.Println(nodo.Contenido.(*Producto).Nombre + ": " + strconv.Itoa(nodo.Contenido.(*Producto).Cantidad))
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

func (arbol *Arbol) GetTag(n *NodoArbol, meses bool) string {
	var cadena string = fmt.Sprintf("nodo%p", n) + "[label=\"<f0>|{" + "Peso: " + strconv.Itoa(n.Peso) + "|" + n.GetDato(meses) + "}|<f1>\"];\n"
	if n.Der == nil && n.Izq == nil {
		return cadena
	}
	if n.Izq != nil {
		cadena += arbol.GetTag(n.Izq, meses)
		cadena += fmt.Sprintf("nodo%p", n) + ": <f0>" + "->" + fmt.Sprintf("nodo%p", n.Izq) + "[arrowhead=rvee, color=brown];\n"
	}
	if n.Der != nil {
		cadena += arbol.GetTag(n.Der, meses)
		cadena += fmt.Sprintf("nodo%p", n) + ": <f1>" + "->" + fmt.Sprintf("nodo%p", n.Der) + "[arrowhead=lvee, color=brown];\n"
	}
	return cadena
}

func (arbol *Arbol) TiendaTag(n *NodoArbol) string {
	var cadena string = fmt.Sprintf("nodo%p", n) + "[label=\"<f0>|{" + "Peso: " + strconv.Itoa(n.Peso) + "|{" +
		strconv.Itoa(n.Dato) + "|" + n.Contenido.(*Producto).Nombre + "|Q " + fmt.Sprintf("%.2f", n.Contenido.(*Producto).Precio) +
		"}|Cant.:" + strconv.Itoa(n.Contenido.(*Producto).Cantidad) + "}|<f1>\"];\n"
	if n.Der == nil && n.Izq == nil {
		return cadena
	}
	if n.Izq != nil {
		cadena += arbol.TiendaTag(n.Izq)
		cadena += fmt.Sprintf("nodo%p", n) + ": <f0>" + "->" + fmt.Sprintf("nodo%p", n.Izq) + "[arrowhead=rvee, color=brown];\n"
	}
	if n.Der != nil {
		cadena += arbol.TiendaTag(n.Der)
		cadena += fmt.Sprintf("nodo%p", n) + ": <f1>" + "->" + fmt.Sprintf("nodo%p", n.Der) + "[arrowhead=lvee, color=brown];\n"
	}
	return cadena
}

func (nodo *NodoArbol) GetDato(meses bool) string {
	if meses {
		return GetMesName(nodo.Dato)
	}
	return strconv.Itoa(nodo.Dato)
}

func (arbol *Arbol) Graficar(meses bool) string {
	if arbol.raiz != nil {
		archivo := "digraph G{\nnode[shape=Mrecord, color=\"#00bf0d\"];\nrankdir=TD;\n"
		if reflect.TypeOf(arbol.raiz.Contenido).String() == "*estructuras.Producto" {
			archivo += arbol.TiendaTag(arbol.raiz)
		} else {
			archivo += arbol.GetTag(arbol.raiz, meses)
		}
		archivo += "}"
		return archivo
	} else {
		return ""
	}
}

func (arbol *Arbol) ToArrayProductos() []*Producto {
	return toArrayProductos(arbol.raiz)
}

func toArrayProductos(raiz *NodoArbol) []*Producto {
	var array []*Producto
	if raiz == nil {
		array = make([]*Producto, 0)
	} else {
		array = append(array, toArrayProductos(raiz.Der)...)
		array = append(array, raiz.Contenido.(*Producto))
		array = append(array, toArrayProductos(raiz.Izq)...)
	}
	return array
}
