package estructuras

import "reflect"

type NodoCabeceraVertical struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    interface{}
}

type NodoMatriz struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    interface{}
}

type NodoCabeceraHorizontal struct {
	Este, Oeste, Sur, Norte interface{}
	Dato                    interface{}
}

type Matriz struct {
	CabeceraH *NodoCabeceraHorizontal
	CabeceraV *NodoCabeceraVertical
}

func (matriz *Matriz) getVertical(dato int) interface{} {
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

func (matriz *Matriz) crearVertical(dato int) *NodoCabeceraVertical {
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
	if dato < aux.(*NodoCabeceraVertical).Dato.(int) {
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
		if dato > aux.(*NodoCabeceraVertical).Dato.(int) && dato < aux.(*NodoCabeceraVertical).Sur.(*NodoCabeceraVertical).Dato.(int) {
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
	if dato < aux.(*NodoCabeceraHorizontal).Dato.(int) {
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
		if dato > aux.(*NodoCabeceraHorizontal).Dato.(int) && dato < aux.(*NodoCabeceraHorizontal).Sur.(*NodoCabeceraHorizontal).Dato.(int) {
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

func (matriz *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, dato int) interface{} {
	if cabecera.Sur == nil {
		return cabecera
	}
	aux := cabecera.Sur
	if dato <= aux.(*NodoMatriz).Dato.(int) {
		return cabecera
	}
	for aux.(*NodoMatriz).Sur != nil {
		if dato > aux.(*NodoMatriz).Dato.(int) && dato <= aux.(*NodoMatriz).Sur.(*NodoMatriz).Dato.(int) {
			return aux
		}
		aux = aux.(*NodoMatriz).Sur
	}
	if dato <= aux.(*NodoMatriz).Dato.(int) {
		return aux.(*NodoMatriz).Norte
	}
	return aux
}

func (matriz *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dato int) interface{} {
	if cabecera.Este == nil {
		return cabecera
	}
	aux := cabecera.Este
	if dato <= aux.(*NodoMatriz).Dato.(int) {
		return cabecera
	}
	for aux.(*NodoMatriz).Este != nil {
		if dato > aux.(*NodoMatriz).Dato.(int) && dato <= aux.(*NodoMatriz).Este.(*NodoMatriz).Dato.(int) {
			return aux
		}
		aux = aux.(*NodoMatriz).Este
	}
	if dato <= aux.(*NodoMatriz).Dato.(int) {
		return aux.(*NodoMatriz).Oeste
	}
	return aux
}

func (matriz *Matriz) Add(nueva *NodoMatriz) {
	vertical := matriz.getVertical(nueva.Dato.(int))
	horizontal := matriz.getHorizontal(nueva.Dato.(int))
	if vertical == nil {
		vertical = matriz.crearVertical(nueva.Dato.(int))
	}
	if horizontal == nil {
		horizontal = matriz.crearHorizontal(nueva.Dato.(int))
	}
	izquierda := matriz.obtenerUltimoH(vertical.(*NodoCabeceraVertical), nueva.Dato.(int))
	superior := matriz.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal), nueva.Dato.(int))
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
