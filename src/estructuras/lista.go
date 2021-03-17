package estructuras

import "reflect"

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

func (this *Lista) Buscar(dato string) *Nodo {
	if this.Size == 0 {
		return nil
	} else {
		aux := this.First
		for i := 0; i < this.Size; i++ {
			if reflect.TypeOf(aux.Contenido).String() == "*estructuras.NodoTienda" {
				if aux.Contenido.(*NodoTienda).Tienda.Nombre == dato {
					return aux
				}
			}
			aux = aux.Next
		}
		return nil
	}
}

func (this *Lista) Eliminar(dato string) *Tienda {
	if this.Size == 0 {
		return nil
	} else {
		aux := this.First
		for i := 0; i < this.Size; i++ {
			if reflect.TypeOf(aux.Contenido).String() == "*estructuras.NodoTienda" {
				if aux.Contenido.(*NodoTienda).Tienda.Nombre == dato {
					if i == 0 {
						if this.Size == 1 {
							this.First = nil
							this.Last = nil
							this.Size--
							return aux.Contenido.(*NodoTienda).Tienda
						} else {
							this.First.Next.Prev = nil
							this.First = this.First.Next
							this.Size--
							return aux.Contenido.(*NodoTienda).Tienda
						}
					} else if i == this.Size-1 {
						aux.Prev.Next = nil
						this.Last = aux.Prev
						this.Size--
						return aux.Contenido.(*NodoTienda).Tienda
					} else {
						aux.Prev.Next = aux.Next
						aux.Next.Prev = aux.Prev
						this.Size--
						return aux.Contenido.(*NodoTienda).Tienda
					}
				}
			}
			aux = aux.Next
		}
		return nil
	}
}

func (this *Lista) InsertarInicio(nuevo *Nodo) {
	nuevo.Next = this.First
	this.First.Prev = nuevo
	this.First = nuevo
}

func (this *Lista) InsertarFinal(nuevo *Nodo) {
	this.Last.Next = nuevo
	nuevo.Prev = this.Last
	this.Last = nuevo
}

func (this *Lista) InsertarEntre(nuevo *Nodo, aux *Nodo) {
	nuevo.Next = aux
	nuevo.Prev = aux.Prev
	aux.Prev.Next = nuevo
	aux.Prev = nuevo
}

func (this *Nodo) GetDatoString() string {
	if reflect.TypeOf(this.Contenido).String() == "*estructuras.NodoTienda" {
		return this.Contenido.(*NodoTienda).Tienda.Nombre
	}
	return ""
}

func (this *Lista) Insertar(nuevo *Nodo) {
	this.Size++
	ascii1 := GetAscii(nuevo.GetDatoString())
	if this.Size-1 == 0 {
		this.First = nuevo
		this.Last = nuevo
		return
	}
	ascii2 := GetAscii(this.First.GetDatoString())
	if this.Size-1 == 1 {
		if ascii1 < ascii2 {
			this.InsertarInicio(nuevo)
			return
		}
		this.InsertarFinal(nuevo)
		return
	}
	var aux *Nodo = this.First
	for i := 0; i < this.Size-1; i++ {
		ascii2 = GetAscii(aux.GetDatoString())
		if ascii1 < ascii2 {
			if i == 0 {
				this.InsertarInicio(nuevo)
				return
			}
			this.InsertarEntre(nuevo, aux)
			return
		}
		if i == this.Size-2 {
			this.InsertarFinal(nuevo)
			return
		}
		aux = aux.Next
	}
}

func (this *Lista) ToString() string {
	var cadena string
	aux := this.First
	for i := 0; i < this.Size; i++ {
		if reflect.TypeOf(aux.Contenido).String() == "*estructuras.NodoTienda" {
			cadena += aux.Contenido.(*NodoTienda).Tienda.Nombre + "\n"
		}
		aux = aux.Next
	}
	return cadena
}

func (this *Lista) ToArray() *[]*Tienda {
	var array []*Tienda
	if this.Size != 0 {
		var aux *Nodo = this.First
		for i := 0; i < this.Size; i++ {
			array = append(array, aux.Contenido.(*NodoTienda).Tienda)
			aux = aux.Next
		}
	}
	return &array
}
