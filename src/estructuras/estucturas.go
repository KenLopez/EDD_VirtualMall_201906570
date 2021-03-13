package estructuras

type producto struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float32
	cantidad    int
	imagen      string
}

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
	Logo         string `json:Logo`
}

func NewTienda() *Tienda {
	return &Tienda{"", "", "", 0, ""}
}

func (this *Tienda) GetAscii() int {
	var ascii int
	for i := 0; i < len(this.Nombre); i++ {
		ascii += int(this.Nombre[i])
	}
	return ascii
}
