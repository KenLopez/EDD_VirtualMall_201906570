package estructuras

import (
	"reflect"
	"strconv"
	"strings"
)

type NodoCabeceraVertical struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    string
}

type NodoMatriz struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    *Lista
}

type NodoCabeceraHorizontal struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    int
}

type Matriz struct {
	CabeceraH *NodoCabeceraHorizontal
	CabeceraV *NodoCabeceraVertical
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
	if matriz.CabeceraV == nil {
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
	if dato < aux.(*NodoCabeceraHorizontal).Dato {
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
		if dato > aux.(*NodoCabeceraHorizontal).Dato && dato < aux.(*NodoCabeceraHorizontal).Sur.(*NodoCabeceraHorizontal).Dato {
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
	if dato <= aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Departamento {
		return cabecera
	}
	for aux.(*NodoMatriz).Sur != nil {
		if dato > aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Departamento && dato <= aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Departamento {
			return aux
		}
		aux = aux.(*NodoMatriz).Sur
	}
	if dato <= aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Departamento {
		return aux.(*NodoMatriz).Norte
	}
	return aux
}

func (matriz *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dato int) interface{} {
	if cabecera.Este == nil {
		return cabecera
	}
	aux := cabecera.Este
	if dato <= GetDia(aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Fecha) {
		return cabecera
	}
	for aux.(*NodoMatriz).Este != nil {
		if dato > GetDia(aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Fecha) && dato <= GetDia(aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Fecha) {
			return aux
		}
		aux = aux.(*NodoMatriz).Este
	}
	if dato <= GetDia(aux.(*NodoMatriz).Dato.First.Contenido.(*Pedido).Fecha) {
		return aux.(*NodoMatriz).Oeste
	}
	return aux
}

func (matriz *Matriz) add(nueva *NodoMatriz) {
	vertical := matriz.getVertical(nueva.Dato.First.Contenido.(*Pedido).Departamento)
	horizontal := matriz.getHorizontal(GetDia(nueva.Dato.First.Contenido.(*Pedido).Fecha))
	if vertical == nil {
		vertical = matriz.crearVertical(nueva.Dato.First.Contenido.(*Pedido).Departamento)
	}
	if horizontal == nil {
		horizontal = matriz.crearHorizontal(GetDia(nueva.Dato.First.Contenido.(*Pedido).Fecha))
	}
	izquierda := matriz.obtenerUltimoH(vertical.(*NodoCabeceraVertical), GetDia(nueva.Dato.First.Contenido.(*Pedido).Fecha))
	superior := matriz.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nueva.Dato.First.Contenido.(*Pedido).Departamento)
	if reflect.TypeOf(izquierda).String() == "estructuras.NodoMatriz" {
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
	if reflect.TypeOf(superior).String() == "estructuras.NodoMatriz" {
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
			tmp := izquierda.(*NodoCabeceraVertical).Este
			izquierda.(*NodoCabeceraVertical).Este = nueva
			nueva.Oeste = izquierda
			tmp.(*NodoMatriz).Oeste = nueva
			nueva.Este = tmp
		}
	}
}

func GetDia(fecha string) int {
	a := strings.Split(fecha, "-")[0]
	b, _ := strconv.Atoi(a)
	return b
}

func GetMesName(numero int) string {
	switch numero {
	case 1:
		return "ENERO"
	case 2:
		return "FEBRERO"
	case 3:
		return "MARZO"
	case 4:
		return "ABRIL"
	case 5:
		return "MAYO"
	case 6:
		return "JUNIO"
	case 7:
		return "JULIO"
	case 8:
		return "AGOSTO"
	case 9:
		return "SEPTIEMBRE"
	case 10:
		return "OCTUBRE"
	case 11:
		return "NOVIEMBRE"
	case 12:
		return "DICIEMBRE"
	}
	return "ERROR"
}
