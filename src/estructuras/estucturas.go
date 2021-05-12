package estructuras

import (
	"strconv"
	"strings"
	"time"
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
	Producto       *HashTable
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

type ComentarioTienda struct {
	Tienda     *RequestFind   `json:Tieda`
	Comentario *SubComentario `json:Comentario`
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
	Comentarios  *HashTable
}

func NewTienda() *Tienda {
	return &Tienda{"", "", "", 0, "", nil}
}

type Comment struct {
	Comentario     *Comentario `json:Comentario`
	SubComentarios []*Comment  `json:Subcomentarios`
}

type Comentario struct {
	Dpi     int    `json:Dpi`
	Mensaje string `json:Mensaje`
	Fecha   string `json:Fecha`
	Hora    string `json:Hora`
}

type SubComentario struct {
	Comentario *Comentario    `json:Comentario`
	Sub        *SubComentario `json:Sub`
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

func CurrentDate() string {
	date := strconv.Itoa(time.Now().Day())
	switch time.Now().Month() {
	case time.January:
		date += "/01/"
		break
	case time.February:
		date += "/02/"
		break
	case time.March:
		date += "/03/"
		break
	case time.April:
		date += "/04/"
		break
	case time.May:
		date += "/05/"
		break
	case time.June:
		date += "/06/"
		break
	case time.July:
		date += "/07/"
		break
	case time.August:
		date += "/08/"
		break
	case time.September:
		date += "/09/"
		break
	case time.October:
		date += "/10/"
		break
	case time.November:
		date += "/11/"
		break
	case time.December:
		date += "/12/"
	}
	date += strconv.Itoa(time.Now().Year())
	return date
}

func CurrentTime() string {
	ctime := strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute()) + ":" + strconv.Itoa(time.Now().Second())
	return ctime
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
