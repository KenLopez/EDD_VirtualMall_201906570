package estructuras

type DeleteReq struct {
	Nombre       string `json:Nombre`
	Categoria    string `json:Categoria`
	Calificacion int    `json:Calificacion`
}

type RequestFind struct {
	Departamento string `json:Departamento`
	Nombre       string `json:Nombre`
	Calificacion int    `json:Calificacion`
}

type Archivo struct {
	Datos []*Dato `json:Datos`
}

type Dato struct {
	Indice        string          `json:Indice`
	Departamentos []*Departamento `json:Departamentos`
}

type Departamento struct {
	Nombre  string    `json:Nombre`
	Tiendas []*Tienda `json:Tiendas`
}

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int    `json:Calificacion`
}

type Nodo struct {
	Tienda     *Tienda
	Next, Prev *Nodo
}

type Lista struct {
	First, Last *Nodo
	Size        int
}

func NewTienda() *Tienda {
	return &Tienda{"", "", "", 0}
}

func NewNodo(tienda *Tienda) *Nodo {
	return &Nodo{tienda, nil, nil}
}

func NewLista() *Lista {
	return &Lista{nil, nil, 0}
}

func (this *Lista) Buscar(tienda string) *Tienda {
	if this.Size == 0 {
		return nil
	} else {
		aux := this.First
		for i := 0; i < this.Size; i++ {
			if aux.Tienda.Nombre == tienda {
				return aux.Tienda
			}
			aux = aux.Next
		}
		return nil
	}
}

func (this *Lista) Eliminar(tienda string) *Tienda {
	if this.Size == 0 {
		return nil
	} else {
		aux := this.First
		for i := 0; i < this.Size; i++ {
			if aux.Tienda.Nombre == tienda {
				if i == 0 {
					if this.Size == 1 {
						this.First = nil
						this.Last = nil
						this.Size--
						return aux.Tienda
					} else {
						this.First.Next.Prev = nil
						this.First = this.First.Next
						this.Size--
						return aux.Tienda
					}
				} else if i == this.Size-1 {
					aux.Prev.Next = nil
					this.Last = aux.Prev
					this.Size--
					return aux.Tienda
				} else {
					aux.Prev.Next = aux.Next
					aux.Next.Prev = aux.Prev
					this.Size--
					return aux.Tienda
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

func (this *Tienda) GetAscii() int {
	var ascii int
	for i := 0; i < len(this.Nombre); i++ {
		ascii += int(this.Nombre[i])
	}
	return ascii
}

func (this *Lista) Insertar(nuevo *Nodo) {
	this.Size++
	ascii1 := nuevo.Tienda.GetAscii()
	if this.Size-1 == 0 {
		this.First = nuevo
		this.Last = nuevo
		return
	}
	ascii2 := this.First.Tienda.GetAscii()
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
		ascii2 = 0
		runes2 := []rune(aux.Tienda.Nombre)
		for j := 0; j < len(runes2); j++ {
			ascii2 += int(runes2[j])
		}
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
		cadena += aux.Tienda.Nombre + "\n"
		aux = aux.Next
	}
	return cadena
}

func (this *Lista) ToArray() *[]*Tienda {
	var array []*Tienda
	if this.Size != 0 {
		var aux *Nodo = this.First
		for i := 0; i < this.Size; i++ {
			array = append(array, aux.Tienda)
			aux = aux.Next
		}
	}
	return &array
}
