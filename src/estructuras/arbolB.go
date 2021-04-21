package estructuras

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
)

type Key struct {
	Valor int
	Izq   *NodoB
	Der   *NodoB
	Dato  interface{}
}

func NewKey(valor int, dato interface{}) *Key {
	return &Key{
		Valor: valor,
		Izq:   nil,
		Der:   nil,
		Dato:  dato,
	}
}

type NodoB struct {
	Max       int
	NodoPadre *NodoB
	Keys      []*Key
}

func NewNodoB(max int) *NodoB {
	return &NodoB{
		Max:       max,
		NodoPadre: nil,
		Keys:      make([]*Key, max),
	}
}

func (nodo *NodoB) CountKeys() int {
	count := 0
	for nodo.Keys[count] != nil {
		count++
		if count == nodo.Max {
			break
		}
	}
	return count
}

type ArbolB struct {
	k    int
	Raiz *NodoB
}

func NewArbolB(m int) *ArbolB {
	return &ArbolB{
		k:    m,
		Raiz: NewNodoB(m),
	}
}

func (arbol *ArbolB) Insertar(key *Key) {
	if arbol.Raiz.Keys[0] != nil {
		insertarB(arbol.Raiz, key)
	} else {
		arbol.Raiz.Keys[0] = key
	}
}

func (nodo *NodoB) insertKey(key *Key) {
	for pos := 0; pos < nodo.Max; pos++ {
		if nodo.Keys[pos] == nil {
			nodo.Keys[pos] = key
			nodo.Keys[pos-1].Der = nodo.Keys[pos].Izq
			break
		} else if nodo.Keys[pos].Valor > key.Valor {
			for i := nodo.CountKeys(); i > pos; i-- {
				nodo.Keys[i] = nodo.Keys[i-1]
			}
			nodo.Keys[pos] = key
			if pos > 0 {
				nodo.Keys[pos-1].Der = nodo.Keys[pos].Izq
			}
			if pos < nodo.Max-1 && nodo.Keys[pos+1] != nil {
				nodo.Keys[pos+1].Izq = nodo.Keys[pos].Der
			}
			break
		}
	}
	if nodo.CountKeys() == nodo.Max {
		mid := (nodo.Max - 1) / 2
		separador := nodo.Keys[mid]
		separador.Izq = NewNodoB(nodo.Max)
		separador.Der = NewNodoB(nodo.Max)
		for i := 0; i < mid; i++ {
			if nodo.Keys[i].Izq != nil {
				nodo.Keys[i].Izq.NodoPadre = separador.Izq
			}
			if nodo.Keys[i].Der != nil {
				nodo.Keys[i].Der.NodoPadre = separador.Izq
			}
			separador.Izq.Keys[i] = nodo.Keys[i]
			if nodo.Keys[mid+1+i].Izq != nil {
				nodo.Keys[mid+1+i].Izq.NodoPadre = separador.Der
			}
			if nodo.Keys[mid+1+i].Der != nil {
				nodo.Keys[mid+1+i].Der.NodoPadre = separador.Der
			}
			separador.Der.Keys[i] = nodo.Keys[mid+1+i]
		}
		if nodo.NodoPadre == nil {
			separador.Izq.NodoPadre = nodo
			separador.Der.NodoPadre = nodo
			newKeys := make([]*Key, 5)
			newKeys[0] = separador
			nodo.Keys = newKeys
		} else {
			separador.Izq.NodoPadre = nodo.NodoPadre
			separador.Der.NodoPadre = nodo.NodoPadre
			nodo.NodoPadre.insertKey(separador)
		}
	}
}

func insertarB(raiz *NodoB, key *Key) {
	llaves := raiz.CountKeys()
	for i := 0; i < llaves; i++ {
		if key.Valor < raiz.Keys[i].Valor {
			if raiz.Keys[i].Izq != nil {
				insertarB(raiz.Keys[i].Izq, key)
			} else {
				raiz.insertKey(key)
			}
			break
		} else if i == llaves-1 {
			if raiz.Keys[i].Der != nil {
				insertarB(raiz.Keys[i].Der, key)
			} else {
				raiz.insertKey(key)
			}
		}
	}
}

func (arbol *ArbolB) Buscar(key int) interface{} {
	return buscarB(arbol.Raiz, key)
}

func buscarB(raiz *NodoB, key int) interface{} {
	for i := 0; i < raiz.CountKeys(); i++ {
		if key < raiz.Keys[i].Valor {
			if raiz.Keys[i].Izq != nil {
				return buscarB(raiz.Keys[i].Izq, key)
			} else {
				return nil
			}
		} else if key == raiz.Keys[i].Valor {
			return raiz.Keys[i].Dato
		}
		if i == raiz.CountKeys()-1 {
			if raiz.Keys[i].Der != nil {
				return buscarB(raiz.Keys[i].Der, key)
			} else {
				return nil
			}
		}
	}
	return nil
}

func encrypt(mensaje string, keyStr string) []byte {
	text := []byte(mensaje)
	key := make([]byte, 32)
	for i := 0; i < len(keyStr); i++ {
		key[i] = byte(keyStr[i])
	}
	c, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	return gcm.Seal(nonce, nonce, text, nil)

}

func (nodo *NodoB) getTag(cifrado int, key string) string {
	str := fmt.Sprintf("nodo%p", nodo) + "[label=\""
	hijos := ""
	links := ""
	mid := ""
	dpi := ""
	correo := ""
	nombre := ""
	password := ""
	cuenta := ""
	llaves := nodo.CountKeys()
	for i := 0; i <= llaves; i++ {
		if i < llaves {
			password = nodo.Keys[i].Dato.(*Usuario).Password[0:10]
			switch cifrado {
			case 0:
				dpi = strconv.Itoa(nodo.Keys[i].Valor)
				correo = nodo.Keys[i].Dato.(*Usuario).Correo
				nombre = nodo.Keys[i].Dato.(*Usuario).Nombre
				cuenta = nodo.Keys[i].Dato.(*Usuario).Cuenta
				break
			case 1:
				dpi = base64.StdEncoding.EncodeToString(encrypt(strconv.Itoa(nodo.Keys[i].Valor), key))[15:25]
				correo = base64.StdEncoding.EncodeToString(encrypt(nodo.Keys[i].Dato.(*Usuario).Correo, key)[15:25])
				nombre = nodo.Keys[i].Dato.(*Usuario).Nombre
				cuenta = nodo.Keys[i].Dato.(*Usuario).Cuenta
				break
			case 2:
				dpi = base64.StdEncoding.EncodeToString(encrypt(strconv.Itoa(nodo.Keys[i].Valor), key))[15:25]
				correo = base64.StdEncoding.EncodeToString(encrypt(nodo.Keys[i].Dato.(*Usuario).Correo, key))[15:25]
				nombre = base64.StdEncoding.EncodeToString(encrypt(nodo.Keys[i].Dato.(*Usuario).Nombre, key))[15:25]
				cuenta = base64.StdEncoding.EncodeToString(encrypt(nodo.Keys[i].Dato.(*Usuario).Cuenta, key))[15:25]
				break
			}
			if i == llaves/2 {
				mid = "<mid>"
			} else {
				mid = ""
			}
			str += "<f" + strconv.Itoa(i) + ">|{{" + mid + dpi + "}|{" + correo + "}|{" + nombre + "|" +
				password + "}|" + cuenta + "}|"
			if nodo.Keys[i].Izq != nil {
				hijos += nodo.Keys[i].Izq.getTag(cifrado, key)
				llavesH := nodo.Keys[i].Izq.CountKeys()
				if llavesH%2 != 0 {
					links += fmt.Sprintf("nodo%p", nodo) + ":" + "f" + strconv.Itoa(i) + "->" + fmt.Sprintf("nodo%p", nodo.Keys[i].Izq) + ":<mid>\n"
				} else {
					links += fmt.Sprintf("nodo%p", nodo) + ":" + "f" + strconv.Itoa(i) + "->" + fmt.Sprintf("nodo%p", nodo.Keys[i].Izq) + ":" + "<f" + strconv.Itoa(llavesH/2) + ">\n"
				}
			}
		} else {
			str += "<f" + strconv.Itoa(i) + ">\"]\n"
			if nodo.Keys[i-1].Der != nil {
				hijos += nodo.Keys[i-1].Der.getTag(cifrado, key)
				llavesH := nodo.Keys[i-1].Der.CountKeys()
				if llavesH%2 != 0 {
					links += fmt.Sprintf("nodo%p", nodo) + ":" + "f" + strconv.Itoa(i) + "->" + fmt.Sprintf("nodo%p", nodo.Keys[i-1].Der) + ":<mid>\n"
				} else {
					links += fmt.Sprintf("nodo%p", nodo) + ":" + "f" + strconv.Itoa(i) + "->" + fmt.Sprintf("nodo%p", nodo.Keys[i-1].Der) + ":" + "<f" + strconv.Itoa(llavesH/2) + ">\n"
				}

			}
		}
	}
	return str + hijos + links
}

func (arbol *ArbolB) Graficar(cifrado int, key string) string {
	return "digraph G{\nnode [shape=Mrecord, color=purple]\n" + arbol.Raiz.getTag(cifrado, key) + "}"
}

func (arbol *ArbolB) Eliminar(key int) *Key {
	return eliminarB(arbol.Raiz, key)
}

func (nodo *NodoB) getLast() *Key {
	return nodo.Keys[nodo.CountKeys()-1]
}

func (nodo *NodoB) prestarIzq(pos int) {
	nodo.insertKey(NewKey(nodo.NodoPadre.Keys[pos].Valor, nodo.NodoPadre.Keys[pos].Dato))
	prestado := nodo.NodoPadre.Keys[pos].Izq.getLast()
	nodo.NodoPadre.Keys[pos].Dato = prestado.Dato
	nodo.NodoPadre.Keys[pos].Valor = prestado.Valor
	nodo.NodoPadre.Keys[pos].Izq.elimKey(nodo.NodoPadre.Keys[pos].Izq.CountKeys() - 1)
}

func (nodo *NodoB) prestarDer(pos int) {
	nodo.insertKey(NewKey(nodo.NodoPadre.Keys[pos].Valor, nodo.NodoPadre.Keys[pos].Dato))
	prestado := nodo.NodoPadre.Keys[pos].Der.Keys[0]
	nodo.NodoPadre.Keys[pos].Dato = prestado.Dato
	nodo.NodoPadre.Keys[pos].Valor = prestado.Valor
	nodo.NodoPadre.Keys[pos].Der.elimKey(0)
}

func (nodo *NodoB) rebalancear() {
	pos := 0
	llavesPadre := nodo.NodoPadre.CountKeys()
	for i := 0; i < llavesPadre; i++ {
		if nodo.NodoPadre.Keys[i].Izq == nodo {
			pos = i
			break
		}
		if i == llavesPadre-1 {
			pos = i + 1
		}
	}
	if pos == llavesPadre && nodo.NodoPadre.Keys[pos-1].Izq.CountKeys() > (nodo.Max-1)/2 {
		nodo.prestarIzq(pos - 1)
	} else if pos == 0 && nodo.NodoPadre.Keys[pos].Der.CountKeys() > (nodo.Max-1)/2 {
		nodo.prestarDer(pos)
	} else {
		if pos < llavesPadre && nodo.NodoPadre.Keys[pos].Der.CountKeys() > (nodo.Max-1)/2 {
			nodo.prestarIzq(pos)
		} else if pos > 0 && nodo.NodoPadre.Keys[pos-1].Izq.CountKeys() > (nodo.Max-1)/2 {
			nodo.prestarDer(pos - 1)
		} else {
			newHijo := NewNodoB(nodo.Max)
			var sep *Key
			var hermano *NodoB
			if pos == llavesPadre {
				sep = nodo.NodoPadre.Keys[pos-1]
				hermano = nodo.NodoPadre.Keys[pos-1].Izq
			} else {
				sep = nodo.NodoPadre.Keys[pos]
				hermano = nodo.NodoPadre.Keys[pos].Der
			}
			newHijo.Keys = nodo.Keys
			sepKey := NewKey(sep.Valor, sep.Dato)
			sepKey.Izq = newHijo.Keys[newHijo.CountKeys()-1].Der
			newHijo.insertKey(sepKey)
			llavesHermano := hermano.CountKeys()
			for i := 0; i < llavesHermano; i++ {
				newHijo.insertKey(hermano.Keys[i])
			}
			if nodo.NodoPadre.NodoPadre == nil {
				for i := 0; i < newHijo.CountKeys(); i++ {
					if newHijo.Keys[i].Izq != nil {
						newHijo.Keys[i].Izq.NodoPadre = nodo.NodoPadre
						newHijo.Keys[i].Der.NodoPadre = nodo.NodoPadre
					}
				}
				nodo.NodoPadre.Keys = newHijo.Keys
			} else {
				nodo.Keys = newHijo.Keys
				if pos > 0 {
					nodo.NodoPadre.Keys[pos-1].Der = nodo
				}
				if pos < llavesPadre-1 {
					nodo.NodoPadre.Keys[pos+1].Izq = nodo
				}
				for i := pos; i < llavesPadre; i++ {
					nodo.NodoPadre.Keys[i] = nodo.NodoPadre.Keys[i+1]
				}
				if nodo.NodoPadre.CountKeys() < (nodo.Max-1)/2 {
					nodo.NodoPadre.rebalancear()
				}
			}
		}
	}
}

func (nodo *NodoB) elimKey(pos int) {
	llaves := nodo.CountKeys()
	if pos == llaves-1 {
		nodo.Keys[pos] = nil
	} else {
		for i := pos; i < nodo.Max-1; i++ {
			nodo.Keys[i] = nodo.Keys[i+1]
		}
	}
	if nodo.NodoPadre != nil && nodo.CountKeys() < (nodo.Max-1)/2 {
		nodo.rebalancear()
	}
}

func (nodo *NodoB) hijoMin() *NodoB {
	llaves := nodo.CountKeys() - 1
	if nodo.Keys[llaves].Der == nil {
		return nodo
	} else {
		return nodo.Keys[llaves].Der.hijoMin()
	}
}

func eliminarB(raiz *NodoB, key int) *Key {
	llaves := raiz.CountKeys()
	for i := 0; i < llaves; i++ {
		if key < raiz.Keys[i].Valor {
			if raiz.Keys[i].Izq != nil {
				return eliminarB(raiz.Keys[i].Izq, key)
			} else {
				return nil
			}
		} else if key == raiz.Keys[i].Valor {
			elim := raiz.Keys[i]
			if raiz.Keys[i].Izq == nil {
				raiz.elimKey(i)
			} else {
				min := raiz.Keys[i].Izq.hijoMin()
				minKeys := min.CountKeys()
				raiz.Keys[i].Valor = min.Keys[minKeys-1].Valor
				raiz.Keys[i].Dato = min.Keys[minKeys-1].Dato
				min.elimKey(minKeys - 1)
			}
			return elim
		}
		if i == raiz.CountKeys()-1 {
			if raiz.Keys[i].Der != nil {
				return eliminarB(raiz.Keys[i].Der, key)
			} else {
				return nil
			}
		}
	}
	return nil
}
