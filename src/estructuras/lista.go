package estructuras

import (
	"reflect"
)

type NodoLista struct {
	Contenido  interface{}
	Next, Prev *NodoLista
}

type NodoTienda struct {
	Tienda     *Tienda
	Inventario *Arbol
}

type Lista struct {
	First, Last *NodoLista
	Size        int
}

func NewLista() *Lista {
	return &Lista{nil, nil, 0}
}

func NewNodoTienda(tienda *Tienda) *NodoTienda {
	return &NodoTienda{tienda, nil}
}

func NewNodo(contenido interface{}) *NodoLista {
	return &NodoLista{contenido, nil, nil}
}

func (lista *Lista) Buscar(dato string) *NodoLista {
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

func (lista *Lista) InsertarInicio(nuevo *NodoLista) {
	nuevo.Next = lista.First
	lista.First.Prev = nuevo
	lista.First = nuevo
}

func (lista *Lista) recount() {
	if lista.First != nil {
		aux := lista.First
		count := 0
		for aux != nil {
			count++
			aux = aux.Next
		}
		lista.Size = count
	}
}

func (lista *Lista) InsertarSimple(nuevo *NodoLista) {
	if nuevo != nil {
		if lista.First == nil {
			lista.First = nuevo
		} else {
			aux := lista.First
			for aux.Next != nil {
				aux = aux.Next
			}
			aux.Next = nuevo
		}
		if nuevo.Next != nil {
			lista.recount()
		} else {
			lista.Size++
		}
	}
}

func (lista *Lista) InsertarFinal(nuevo *NodoLista) {
	lista.Last.Next = nuevo
	nuevo.Prev = lista.Last
	lista.Last = nuevo
}

func (lista *Lista) InsertarEntre(nuevo *NodoLista, aux *NodoLista) {
	nuevo.Next = aux
	nuevo.Prev = aux.Prev
	aux.Prev.Next = nuevo
	aux.Prev = nuevo
}

func (nodo *NodoLista) GetDatoString() string {
	if reflect.TypeOf(nodo.Contenido).String() == "*estructuras.NodoTienda" {
		return nodo.Contenido.(*NodoTienda).Tienda.Nombre
	}
	return ""
}

func (lista *Lista) Insertar(nuevo *NodoLista) {
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
	var aux *NodoLista = lista.First
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
		var aux *NodoLista = lista.First
		for i := 0; i < lista.Size; i++ {
			array = append(array, aux.Contenido.(*NodoTienda).Tienda)
			aux = aux.Next
		}
	}
	return &array
}

func (lista *Lista) MovToArray() *[]struct {
	Origen  string
	Peso    float32
	Destino string
} {
	array := make([]struct {
		Origen  string
		Peso    float32
		Destino string
	}, lista.Size)
	aux := lista.First
	i := 0
	for aux != nil {
		array[i] = struct {
			Origen  string
			Peso    float32
			Destino string
		}{
			Origen:  aux.Contenido.(*Movimiento).Origen.Nombre,
			Peso:    aux.Contenido.(*Movimiento).Peso,
			Destino: aux.Contenido.(*Movimiento).Destino.Nombre,
		}
		i++
		aux = aux.Next
	}
	return &array
}
