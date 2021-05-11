package estructuras

import "math"

type NodoHash struct {
	hash           int
	Data           *Comentario
	Subcomentarios *HashTable
}

func NewNodoHash(key int, value *Comentario) *NodoHash {
	return &NodoHash{
		hash:           key,
		Data:           value,
		Subcomentarios: NewHashTable(),
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
	for i := 2; i < next; i++ {
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
	for tabla.datos[p].hash != key && tabla.datos[p].Data.Fecha != value.Fecha && tabla.datos[p].Data.Hora != value.Hora {
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
		tmp.Subcomentarios.InsertarSub(comentario.Sub)
	} else {
		tabla.insertar(comentario.Comentario.Dpi, comentario.Comentario)
	}
}

func (tabla *HashTable) hashing(key int) int {
	return int(math.Trunc(float64(tabla.size) * ((0.2520 * float64(key)) - math.Trunc(0.2520*float64(key)))))
}

func (tabla *HashTable) closedHashing(p int, i int) int {
	return p + tabla.hashing(i)
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
	for tabla.datos[p] != nil && tabla.datos[p].hash != key {
		i++
		p = tabla.closedHashing(p, i)
		if p >= tabla.size {
			p = tabla.tableLimit(p)
		}
	}
	return p
}
