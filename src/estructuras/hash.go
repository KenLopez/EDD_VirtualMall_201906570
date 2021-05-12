package estructuras

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type NodoHash struct {
	hash           int
	Data           *Comentario
	Subcomentarios *HashTable
}

func (tabla *HashTable) ToArray() []*Comment {
	array := make([]*Comment, tabla.carga)
	pos := 0
	for i := 0; i < tabla.size; i++ {
		if tabla.datos[i] != nil {
			var subs []*Comment
			com := tabla.datos[i].Data
			if tabla.datos[i].Subcomentarios != nil {
				subs = tabla.datos[i].Subcomentarios.ToArray()
			} else {
				subs = make([]*Comment, 0)
			}
			array[pos] = &Comment{
				Comentario:     com,
				SubComentarios: subs,
			}
			pos++
		}
	}
	quicksort(&array, 0, len(array)-1)
	return array
}

func quicksort(arreglo *[]*Comment, start int, end int) {
	part := partition(arreglo, start, end)
	if (part - 1) > start {
		quicksort(arreglo, start, part-1)
	}
	if (part + 1) < end {
		quicksort(arreglo, part+1, end)
	}
}

func partition(arreglo *[]*Comment, start int, end int) int {
	pivote := (*arreglo)[end]
	panio, _ := strconv.Atoi(strings.Split(pivote.Comentario.Fecha, "/")[2])
	pmes, _ := strconv.Atoi(strings.Split(pivote.Comentario.Fecha, "/")[1])
	pdia, _ := strconv.Atoi(strings.Split(pivote.Comentario.Fecha, "/")[0])
	phora, _ := strconv.Atoi(strings.Split(pivote.Comentario.Hora, ":")[0])
	pmin, _ := strconv.Atoi(strings.Split(pivote.Comentario.Hora, ":")[1])
	psec, _ := strconv.Atoi(strings.Split(pivote.Comentario.Hora, ":")[2])
	for i := start; i < end; i++ {
		anio, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Fecha, "/")[2])
		mes, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Fecha, "/")[1])
		dia, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Fecha, "/")[0])
		hora, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Hora, ":")[0])
		min, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Hora, ":")[1])
		sec, _ := strconv.Atoi(strings.Split((*arreglo)[i].Comentario.Hora, ":")[2])
		if anio < panio {
			tmp := (*arreglo)[start]
			(*arreglo)[start] = (*arreglo)[i]
			(*arreglo)[i] = tmp
			start++
		} else if anio == panio {
			if mes < pmes {
				tmp := (*arreglo)[start]
				(*arreglo)[start] = (*arreglo)[i]
				(*arreglo)[i] = tmp
				start++
			} else if mes == pmes {
				if dia < pdia {
					tmp := (*arreglo)[start]
					(*arreglo)[start] = (*arreglo)[i]
					(*arreglo)[i] = tmp
					start++
				} else if dia == pdia {
					if hora < phora {
						tmp := (*arreglo)[start]
						(*arreglo)[start] = (*arreglo)[i]
						(*arreglo)[i] = tmp
						start++
					} else if hora == phora {
						if min < pmin {
							tmp := (*arreglo)[start]
							(*arreglo)[start] = (*arreglo)[i]
							(*arreglo)[i] = tmp
							start++
						} else if min == pmin {
							if sec < psec {
								tmp := (*arreglo)[start]
								(*arreglo)[start] = (*arreglo)[i]
								(*arreglo)[i] = tmp
								start++
							}
						}
					}
				}
			}
		}
	}
	tmp := (*arreglo)[start]
	(*arreglo)[start] = pivote
	(*arreglo)[end] = tmp
	return start
}

func NewNodoHash(key int, value *Comentario) *NodoHash {
	return &NodoHash{
		hash:           key,
		Data:           value,
		Subcomentarios: nil,
	}
}

type HashTable struct {
	size      int
	carga     int
	capacidad int
	datos     []*NodoHash
}

func NewHashTable() *HashTable {
	return &HashTable{
		size:      7,
		carga:     0,
		capacidad: 60,
		datos:     make([]*NodoHash, 7),
	}
}

func nextPrime(n int) int {
	if n < 2 {
		return 2
	} else if n == 2 {
		return 3
	}
	next := n + 1
	for i := 2; i < int(next/2); i++ {
		if (next%i == 0) && (i != next) {
			return nextPrime(next)
		}
	}
	return next
}

func (tabla *HashTable) insertar(nuevo int, value *Comentario) {
	node := NewNodoHash(nuevo, value)
	pos := tabla.position(nuevo)
	tabla.datos[pos] = node
	tabla.carga++
	if ((tabla.carga * 100) / tabla.size) > tabla.capacidad {
		old := tabla.datos
		tabla.datos = make([]*NodoHash, nextPrime(tabla.size))
		tabla.size = len(tabla.datos)
		aux := 0
		for i := 0; i < len(old); i++ {
			if old[i] != nil {
				aux = tabla.position(old[i].hash)
				tabla.datos[aux] = old[i]
			}
		}
	}
}

func (tabla *HashTable) find(key int, value *Comentario) *NodoHash {
	i, p := 0, 0
	p = tabla.hashing(key)
	for !((tabla.datos[p].hash == key) && (tabla.datos[p].Data.Fecha == value.Fecha) && (tabla.datos[p].Data.Hora == value.Hora)) {
		i++
		p = tabla.closedHashing(p, i)
		if p >= tabla.size {
			p = tabla.tableLimit(p)
		}
	}
	return tabla.datos[p]
}

func (tabla *HashTable) InsertarSub(comentario *SubComentario) {
	if comentario.Sub != nil {
		tmp := tabla.find(comentario.Comentario.Dpi, comentario.Comentario)
		if tmp.Subcomentarios == nil {
			tmp.Subcomentarios = NewHashTable()
		}
		tmp.Subcomentarios.InsertarSub(comentario.Sub)
	} else {
		comentario.Comentario.Fecha = CurrentDate()
		comentario.Comentario.Hora = CurrentTime()
		tabla.insertar(comentario.Comentario.Dpi, comentario.Comentario)
	}
	tabla.imprimir()
}

func (tabla *HashTable) hashing(key int) int {
	return int(math.Trunc(float64(tabla.size) * ((0.2520 * float64(key)) - math.Trunc(0.2520*float64(key)))))
}

func (tabla *HashTable) closedHashing(p int, i int) int {
	return p + tabla.hashing(i*i)
}

func (tabla *HashTable) tableLimit(p int) int {
	tmp := p - tabla.size
	for tmp >= tabla.size {
		tmp = tmp - tabla.size
	}
	return tmp
}

func (tabla *HashTable) position(key int) int {
	i, p := 0, 0
	p = tabla.hashing(key)
	for tabla.datos[p] != nil {
		i++
		p = tabla.closedHashing(p, i)
		if p >= tabla.size {
			p = tabla.tableLimit(p)
		}
	}
	return p
}

func (tabla *HashTable) imprimir() {
	data := make([][]string, tabla.size)
	for i := 0; i < len(tabla.datos); i++ {
		tmp := make([]string, 2)
		aux := tabla.datos[i]
		if aux != nil {
			tmp[0] = strconv.Itoa(aux.hash)
			tmp[1] = aux.Data.Mensaje
		} else {
			tmp[0] = "-"
			tmp[1] = "-"
		}
		data[i] = tmp
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Valor"})
	table.SetFooter([]string{"size", strconv.Itoa(tabla.size), "Carga", strconv.Itoa(tabla.carga)})
	table.AppendBulk(data)
	table.Render()
}
