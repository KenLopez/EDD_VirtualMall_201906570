package estructuras

import (
	"fmt"
	"reflect"
	"strconv"
)

type NodoCabeceraVertical struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    string
}

type NodoMatriz struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    *Cola
}

type NodoCabeceraHorizontal struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    int
}

type Matriz struct {
	CabeceraH *NodoCabeceraHorizontal
	CabeceraV *NodoCabeceraVertical
}

func NewMatriz() *Matriz {
	return &Matriz{
		CabeceraH: nil,
		CabeceraV: nil,
	}
}

func NewNodoMatriz(cola *Cola) *NodoMatriz {
	return &NodoMatriz{
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		Dato:  cola,
	}
}

func (matriz *Matriz) getVertical(dato string) interface{} {
	if matriz.CabeceraV == nil {
		return nil
	}
	var aux interface{} = matriz.CabeceraV
	for aux != nil {
		if aux.(*NodoCabeceraVertical).Dato == dato {
			return aux
		}
		aux = aux.(*NodoCabeceraVertical).Sur
	}
	return nil
}

func (matriz *Matriz) getHorizontal(dato int) interface{} {
	if matriz.CabeceraH == nil {
		return nil
	}
	var aux interface{} = matriz.CabeceraH
	for aux != nil {
		if aux.(*NodoCabeceraHorizontal).Dato == dato {
			return aux
		}
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
	return nil
}

func (matriz *Matriz) crearVertical(dato string) *NodoCabeceraVertical {
	if matriz.CabeceraV == nil {
		nueva := &NodoCabeceraVertical{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		matriz.CabeceraV = nueva
		return nueva
	}
	var aux interface{} = matriz.CabeceraV
	if dato < aux.(*NodoCabeceraVertical).Dato {
		nueva := &NodoCabeceraVertical{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		nueva.Sur = matriz.CabeceraV
		matriz.CabeceraV.Norte = nueva
		matriz.CabeceraV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).Sur != nil {
		if dato > aux.(*NodoCabeceraVertical).Dato && dato < aux.(*NodoCabeceraVertical).Sur.(*NodoCabeceraVertical).Dato {
			nueva := &NodoCabeceraVertical{
				Este:  nil,
				Oeste: nil,
				Sur:   nil,
				Norte: nil,
				Dato:  dato,
			}
			tmp := aux.(*NodoCabeceraVertical).Sur
			tmp.(*NodoCabeceraVertical).Norte = nueva
			nueva.Sur = tmp
			aux.(*NodoCabeceraVertical).Sur = nueva
			nueva.Norte = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).Sur
	}
	nueva := &NodoCabeceraVertical{
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		Dato:  dato,
	}
	aux.(*NodoCabeceraVertical).Sur = nueva
	nueva.Norte = aux
	return nueva
}

func (matriz *Matriz) crearHorizontal(dato int) *NodoCabeceraHorizontal {
	if matriz.CabeceraH == nil {
		nueva := &NodoCabeceraHorizontal{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		matriz.CabeceraH = nueva
		return nueva
	}
	var aux interface{} = matriz.CabeceraH
	if dato <= aux.(*NodoCabeceraHorizontal).Dato {
		nueva := &NodoCabeceraHorizontal{
			Este:  nil,
			Oeste: nil,
			Sur:   nil,
			Norte: nil,
			Dato:  dato,
		}
		nueva.Este = matriz.CabeceraH
		matriz.CabeceraH.Oeste = nueva
		matriz.CabeceraH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).Este != nil {
		if dato > aux.(*NodoCabeceraHorizontal).Dato && dato <= aux.(*NodoCabeceraHorizontal).Sur.(*NodoCabeceraHorizontal).Dato {
			nueva := &NodoCabeceraHorizontal{
				Este:  nil,
				Oeste: nil,
				Sur:   nil,
				Norte: nil,
				Dato:  dato,
			}
			tmp := aux.(*NodoCabeceraHorizontal).Este
			tmp.(*NodoCabeceraHorizontal).Oeste = nueva
			nueva.Este = tmp
			aux.(*NodoCabeceraHorizontal).Este = nueva
			nueva.Oeste = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).Este
	}
	nueva := &NodoCabeceraHorizontal{
		Este:  nil,
		Oeste: nil,
		Sur:   nil,
		Norte: nil,
		Dato:  dato,
	}
	aux.(*NodoCabeceraHorizontal).Este = nueva
	nueva.Oeste = aux
	return nueva
}

func (matriz *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, dato string) interface{} {
	if cabecera.Sur == nil {
		return cabecera
	}
	aux := cabecera.Sur
	if dato <= aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Departamento {
		return cabecera
	}
	for aux.(*NodoMatriz).Sur != nil {
		if dato > aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Departamento && dato <= aux.(*NodoMatriz).Sur.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Departamento {
			return aux
		}
		aux = aux.(*NodoMatriz).Sur
	}
	if dato <= aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Departamento {
		return aux.(*NodoMatriz).Norte
	}
	return aux
}

func (matriz *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dato int) interface{} {
	if cabecera.Este == nil {
		return cabecera
	}
	aux := cabecera.Este
	if dato <= GetDia(aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) {
		return cabecera
	}
	for aux.(*NodoMatriz).Este != nil {
		if dato > GetDia(aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) && dato <= GetDia(aux.(*NodoMatriz).Este.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) {
			return aux
		}
		aux = aux.(*NodoMatriz).Este
	}
	if dato <= GetDia(aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) {
		return aux.(*NodoMatriz).Oeste
	}
	return aux
}

func (matriz *Matriz) Get(dia int, categoria string) *NodoMatriz {
	cabecera := matriz.getVertical(categoria)
	if cabecera != nil {
		aux := cabecera.(*NodoCabeceraVertical).Este
		for aux != nil {
			if GetDia(aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) == dia {
				return aux.(*NodoMatriz)
			}
		}
	}
	return nil
}

func (matriz *Matriz) NuevoPedido(nuevo *Pedido) {
	cabecera := matriz.getVertical(nuevo.Departamento)
	if cabecera == nil {
		cola := NewCola()
		cola.Queue(nuevo)
		matriz.add(NewNodoMatriz(cola))
	} else {
		aux := cabecera.(*NodoCabeceraVertical).Este
		for aux != nil {
			if GetDia(aux.(*NodoMatriz).Dato.Frente.Contenido.(*Pedido).Fecha) == GetDia(nuevo.Fecha) {
				aux.(*NodoMatriz).Dato.Queue(nuevo)
				break
			}
			aux = aux.(*NodoMatriz).Este
			if aux == nil {
				cola := NewCola()
				cola.Queue(nuevo)
				matriz.add(NewNodoMatriz(cola))
			}
		}
	}
}

func (matriz *Matriz) Graficar(nombre string) string {
	if matriz.CabeceraH != nil && matriz.CabeceraV != nil {
		var archivo string = "graph Matriz{\nlayout=dot\nlabelloc = \"t\"\nnode [shape=Mrecord]\nedge [weight=1000, minlen=1.5 style=rigid color=orange]\n"
		var rank string = ""
		var nodos string = ""
		var conexionesV string = ""
		archivo += fmt.Sprintf("nodo%p", matriz) + "[color=red, label=\"PEDIDOS\\n" + nombre + "\"]\n"
		rango := fmt.Sprintf("nodo%p", matriz) + "--" + fmt.Sprintf("nodo%p", matriz.CabeceraH) + "\n"
		conexionesV += fmt.Sprintf("nodo%p", matriz) + "--" + fmt.Sprintf("nodo%p", matriz.CabeceraV) + "\n"
		var actualH interface{} = matriz.CabeceraH
		for actualH != nil {
			nodos += fmt.Sprintf("nodo%p", actualH) + "[color=purple, label=\"" + strconv.Itoa(actualH.(*NodoCabeceraHorizontal).Dato) + "\"]\n"
			if actualH.(*NodoCabeceraHorizontal).Este != nil {
				rango += fmt.Sprintf("nodo%p", actualH) + "--" + fmt.Sprintf("nodo%p", actualH.(*NodoCabeceraHorizontal).Este) + "\n"
			}
			conexionesV += fmt.Sprintf("nodo%p", actualH) + "--" + fmt.Sprintf("nodo%p", actualH.(*NodoCabeceraHorizontal).Sur) + "\n"
			actualH = actualH.(*NodoCabeceraHorizontal).Este
		}
		rank += "rank=same{\n" + rango + "}\n"
		var actualV interface{} = matriz.CabeceraV
		for actualV != nil {
			rango = fmt.Sprintf("nodo%p", actualV) + "--" + fmt.Sprintf("nodo%p", actualV.(*NodoCabeceraVertical).Este) + "\n"
			nodos += fmt.Sprintf("nodo%p", actualV) + "[color=purple, label=\"" + actualV.(*NodoCabeceraVertical).Dato + "\"]\n"
			var aux interface{} = actualV.(*NodoCabeceraVertical).Este
			for aux != nil {
				nodos += fmt.Sprintf("nodo%p", aux) + "[color=brown, style=filled, fillcolor=beige label=\"\"]\n"
				if aux.(*NodoMatriz).Este != nil {
					rango = fmt.Sprintf("nodo%p", actualV) + "--" + fmt.Sprintf("nodo%p", actualV.(*NodoCabeceraVertical).Este) + "\n"
				}
				if aux.(*NodoMatriz).Sur != nil {
					conexionesV += fmt.Sprintf("nodo%p", aux) + "--" + fmt.Sprintf("nodo%p", aux.(*NodoMatriz).Sur) + "\n"
				}
				aux = aux.(*NodoMatriz).Este
			}
			if actualV.(*NodoCabeceraVertical).Sur != nil {
				conexionesV += fmt.Sprintf("nodo%p", actualV) + "--" + fmt.Sprintf("nodo%p", actualV.(*NodoCabeceraVertical).Sur) + "\n"
			}
			rank += "rank=same{\n" + rango + "}\n"
			actualV = actualV.(*NodoCabeceraVertical).Sur
		}
		archivo += nodos + conexionesV + rank + "\n}"
		return archivo
	} else {
		return ""
	}
}

func (matriz *Matriz) add(nueva *NodoMatriz) {
	vertical := matriz.getVertical(nueva.Dato.Frente.Contenido.(*Pedido).Departamento)
	horizontal := matriz.getHorizontal(GetDia(nueva.Dato.Frente.Contenido.(*Pedido).Fecha))
	if vertical == nil {
		vertical = matriz.crearVertical(nueva.Dato.Frente.Contenido.(*Pedido).Departamento)
	}
	if horizontal == nil {
		horizontal = matriz.crearHorizontal(GetDia(nueva.Dato.Frente.Contenido.(*Pedido).Fecha))
	}
	izquierda := matriz.obtenerUltimoH(vertical.(*NodoCabeceraVertical), GetDia(nueva.Dato.Frente.Contenido.(*Pedido).Fecha))
	superior := matriz.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nueva.Dato.Frente.Contenido.(*Pedido).Departamento)
	if reflect.TypeOf(izquierda).String() == "*estructuras.NodoMatriz" {
		if izquierda.(*NodoMatriz).Este == nil {
			izquierda.(*NodoMatriz).Este = nueva
			nueva.Oeste = izquierda
		} else {
			tmp := izquierda.(*NodoMatriz).Este
			izquierda.(*NodoMatriz).Este = nueva
			nueva.Oeste = izquierda
			tmp.(*NodoMatriz).Oeste = nueva
			nueva.Este = tmp
		}
	} else {
		if izquierda.(*NodoCabeceraVertical).Este == nil {
			izquierda.(*NodoCabeceraVertical).Este = nueva
			nueva.Oeste = izquierda
		} else {
			tmp := izquierda.(*NodoCabeceraVertical).Este
			izquierda.(*NodoCabeceraVertical).Este = nueva
			nueva.Oeste = izquierda
			tmp.(*NodoMatriz).Oeste = nueva
			nueva.Este = tmp
		}
	}
	if reflect.TypeOf(superior).String() == "*estructuras.NodoMatriz" {
		if superior.(*NodoMatriz).Sur == nil {
			superior.(*NodoMatriz).Sur = nueva
			nueva.Norte = superior
		} else {
			tmp := superior.(*NodoMatriz).Sur
			superior.(*NodoMatriz).Sur = nueva
			nueva.Norte = superior
			tmp.(*NodoMatriz).Norte = nueva
			nueva.Sur = tmp
		}
	} else {
		if superior.(*NodoCabeceraHorizontal).Sur == nil {
			superior.(*NodoCabeceraHorizontal).Sur = nueva
			nueva.Norte = superior
		} else {
			tmp := superior.(*NodoCabeceraHorizontal).Sur
			superior.(*NodoCabeceraHorizontal).Sur = nueva
			nueva.Norte = superior
			tmp.(*NodoMatriz).Norte = nueva
			nueva.Sur = tmp
		}
	}
}
