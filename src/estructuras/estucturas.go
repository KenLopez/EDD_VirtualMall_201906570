package estructuras

type Archivo struct {
	Datos []Dato
}

type Dato struct {
	Indice        string
	Departamentos []Departamento
}

type Departamento struct {
	Nombre  string
	Tiendas *Lista
}

type Tienda struct {
	Nombre, Descripcion, Contacto string
	Calificacion                  int
}

type Nodo struct {
	Tienda     *Tienda
	Next, Prev *Nodo
}

type Lista struct {
	First, Last *Nodo
	Size        int
}

func NuevaLista() *Lista {
	return &Lista{nil, nil, 0}
}

func (this *Lista) Insertar(nuevo *Nodo) {
	if this.First == nil {
		this.First = nuevo
		this.Last = nuevo
	} else {
		this.Last.Next = nuevo
		nuevo.Prev = this.Last
		this.Last = nuevo
	}
}

/*func (this *Lista) ToString() string {
	var texto string
	aux := this.first
	for aux != nil {
		texto += aux.ToString()
		aux = aux.Next
	}
	return texto
}

func (this *Nodo) ToString() string {
	var texto string
	texto = "Origen: " + this.Info.Origen + "\n" +
		"Destino: " + this.Info.Destino + "\n" + "Mensajes:\n"
	for i := 0; i < len(this.Info.Msgs); i++ {
		texto += "[" + this.Info.Msgs[i].Fecha + "] " + this.Info.Msgs[i].Texto + "\n\n"
	}
	return texto
}

func (this *Lista) Print() {
	fmt.Println("---LISTA DE MENSAJES---\n")
	this.ToString()
}*/
