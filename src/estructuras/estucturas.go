package estructuras

type ArchivoInventario struct {
	Inventarios []*Inventario `json:Inventarios`
}

type ArchivoPedido struct {
	Pedidos []*Pedido `json:Pedidos`
}

type Pedido struct {
	Fecha        string `json:Fecha`
	Tienda       string `json:Tienda`
	Departamento string `json:Departamento`
	Calificacion int    `json:Calificacion`
	Productos    []*struct {
		Codigo string `json:Codigo`
	} `json:Productos`
}

type Inventario struct {
	Tienda       string      `json:Tienda`
	Departamento string      `json:Departamento`
	Calificacion int         `json:Calificacion`
	Productos    []*Producto `json:Productos`
}

type Producto struct {
	Nombre      string  `json:Nombre`
	Codigo      int     `json:Codigo`
	Descripcion string  `json:Descripcion`
	Precio      float32 `json:Precio`
	Cantidad    int     `json:Cantidad`
	Imagen      string  `json:Imagen`
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

func GetAscii(cadena string) int {
	var ascii int
	for i := 0; i < len(cadena); i++ {
		ascii += int(cadena[i])
	}
	return ascii
}
