package estructuras

import (
	"reflect"
)

type Nodo struct {
	Contenido  interface{}
	Next, Prev *Nodo
}

type NodoTienda struct {
	Tienda     *Tienda
	Inventario *Arbol
}

type Lista struct {
	First, Last *Nodo
	Size        int
}

func NewLista() *Lista {
	return &Lista{nil, nil, 0}
}

func NewNodoTienda(tienda *Tienda) *NodoTienda {
	return &NodoTienda{tienda, nil}
}

func NewNodo(contenido interface{}) *Nodo {
	return &Nodo{contenido, nil, nil}
}

func (lista *Lista) Buscar(dato string) *Nodo {
	if lista.Size == 0 {
		return nil
	} else {
		aux := lista.First
		for i := 0; i < lista.Size; i++ {
			if aux.GetDatoString() == dato {
				return aux
			}
			aux = aux.Next
		}
		return nil
	}
}

func (lista *Lista) Eliminar(dato string) *Tienda {
	if lista.Size == 0 {
		return nil
	} else {
		aux := lista.First
		for i := 0; i < lista.Size; i++ {
			if reflect.TypeOf(aux.Contenido).String() == "*estructuras.NodoTienda" {
				if aux.Contenido.(*NodoTienda).Tienda.Nombre == dato {
					if i == 0 {
						if lista.Size == 1 {
							lista.First = nil
							lista.Last = nil
							lista.Size--
							return aux.Contenido.(*NodoTienda).Tienda
						} else {
							lista.First.Next.Prev = nil
							lista.First = lista.First.Next
							lista.Size--
							return aux.Contenido.(*NodoTienda).Tienda
						}
					} else if i == lista.Size-1 {
						aux.Prev.Next = nil
						lista.Last = aux.Prev
						lista.Size--
						return aux.Contenido.(*NodoTienda).Tienda
					} else {
						aux.Prev.Next = aux.Next
						aux.Next.Prev = aux.Prev
						lista.Size--
						return aux.Contenido.(*NodoTienda).Tienda
					}
				}
			}
			aux = aux.Next
		}
		return nil
	}
}

func (lista *Lista) InsertarInicio(nuevo *Nodo) {
	nuevo.Next = lista.First
	lista.First.Prev = nuevo
	lista.First = nuevo
}

func (lista *Lista) InsertarFinal(nuevo *Nodo) {
	lista.Last.Next = nuevo
	nuevo.Prev = lista.Last
	lista.Last = nuevo
}

func (lista *Lista) InsertarEntre(nuevo *Nodo, aux *Nodo) {
	nuevo.Next = aux
	nuevo.Prev = aux.Prev
	aux.Prev.Next = nuevo
	aux.Prev = nuevo
}

func (nodo *Nodo) GetDatoString() string {
	if reflect.TypeOf(nodo.Contenido).String() == "*estructuras.NodoTienda" {
		return nodo.Contenido.(*NodoTienda).Tienda.Nombre
	}
	return ""
}

func (lista *Lista) Insertar(nuevo *Nodo) {
	lista.Size++
	ascii1 := GetAscii(nuevo.GetDatoString())
	if lista.Size-1 == 0 {
		lista.First = nuevo
		lista.Last = nuevo
		return
	}
	ascii2 := GetAscii(lista.First.GetDatoString())
	if lista.Size-1 == 1 {
		if ascii1 < ascii2 {
			lista.InsertarInicio(nuevo)
			return
		}
		lista.InsertarFinal(nuevo)
		return
	}
	var aux *Nodo = lista.First
	for i := 0; i < lista.Size-1; i++ {
		ascii2 = GetAscii(aux.GetDatoString())
		if ascii1 < ascii2 {
			if i == 0 {
				lista.InsertarInicio(nuevo)
				return
			}
			lista.InsertarEntre(nuevo, aux)
			return
		}
		if i == lista.Size-2 {
			lista.InsertarFinal(nuevo)
			return
		}
		aux = aux.Next
	}
}

func (lista *Lista) ToString() string {
	var cadena string
	aux := lista.First
	for i := 0; i < lista.Size; i++ {
		if reflect.TypeOf(aux.Contenido).String() == "*estructuras.NodoTienda" {
			cadena += aux.Contenido.(*NodoTienda).Tienda.Nombre + "\n"
		}
		aux = aux.Next
	}
	return cadena
}

func (lista *Lista) ToArray() *[]*Tienda {
	var array []*Tienda
	if lista.Size != 0 {
		var aux *Nodo = lista.First
		for i := 0; i < lista.Size; i++ {
			array = append(array, aux.Contenido.(*NodoTienda).Tienda)
			aux = aux.Next
		}
	}
	return &array
}
