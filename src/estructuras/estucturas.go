package estructuras

import (
	"strconv"
	"strings"
)

type ArchivoUsuarios struct {
	Usuarios []*Usuario `json:Usuarios`
}

type UserLogin struct {
	Dpi      int
	Password string
}

type Usuario struct {
	Dpi      int    `json:Dpi`
	Nombre   string `json:Nombre`
	Correo   string `json:Correo`
	Password string `json:Password`
	Cuenta   string `json:Cuenta`
}

type ArchivoGrafo struct {
	Nodos                []*NodoGrafo `json:Nodos`
	PosicionInicialRobot string       `json:PosicionInicialRobot`
	Entrega              string       `json:Entrega`
}

type NodoGrafo struct {
	Nombre  string           `json:Nombre`
	Enlaces []*EnlaceArchivo `json:Enlace`
}

type EnlaceArchivo struct {
	Nombre    string  `json:Nombre`
	Distancia float32 `json:Distancia`
}

type ArchivoInventario struct {
	Inventarios []*Inventario `json:Inventarios`
}

type ArchivoPedido struct {
	Pedidos []*Pedido `json:Pedidos`
}

type Pedido struct {
	Fecha        string    `json:Fecha`
	Tienda       string    `json:Tienda`
	Departamento string    `json:Departamento`
	Calificacion int       `json:Calificacion`
	Productos    []*Codigo `json:Productos`
	Cliente      int       `json:Cliente`
	CaminoCorto  *Lista
}

type Codigo struct {
	Codigo int `json:Codigo`
}

type Response struct {
	Tipo    string `json:Tipo`
	Content string `json:Content`
}

type Inventario struct {
	Tienda       string      `json:Tienda`
	Departamento string      `json:Departamento`
	Calificacion int         `json:Calificacion`
	Productos    []*Producto `json:Productos`
}

type Producto struct {
	Nombre         string  `json:Nombre`
	Codigo         int     `json:Codigo`
	Descripcion    string  `json:Descripcion`
	Precio         float32 `json:Precio`
	Cantidad       int     `json:Cantidad`
	Imagen         string  `json:Imagen`
	Almacenamiento string  `json:Almacenamiento`
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

func GetDia(fecha string) int {
	a := strings.Split(fecha, "-")[0]
	b, _ := strconv.Atoi(a)
	return b
}

func GetAnio(fecha string) int {
	a := strings.Split(fecha, "-")[2]
	b, _ := strconv.Atoi(a)
	return b
}

func GetMes(fecha string) int {
	a := strings.Split(fecha, "-")[1]
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
