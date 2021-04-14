package estructuras

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/fernet/fernet-go"
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

func (arbol *ArbolB) Eliminar(key int) *Key {
	return eliminarB(arbol.Raiz, key)
}

func (nodo *NodoB) getMin() int {
	return (nodo.Max - 1) / 2
}

func (nodo *NodoB) getLast() *Key {
	for i := 0; i < nodo.Max; i++ {
		if nodo.Keys[i] == nil {
			return nodo.Keys[i-1]
		}
	}
	return nil
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
			if raiz.Keys[i].Der != nil {
				if raiz.Keys[i+1] == nil {
					raiz.Keys[i] = nil
				} else {
					for i := 0; i < llaves-1; i++ {
						raiz.Keys[i] = raiz.Keys[i+1]
					}
				}
				if raiz.CountKeys() < raiz.getMin() {
					rebalancear(raiz)
				}
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

func rebalancear(nodo *NodoB) {
	if nodo.NodoPadre != nil {
		llaves := nodo.NodoPadre.CountKeys()
		pos := 0
		for pos = 0; pos < llaves; pos++ {
			if nodo.NodoPadre.Keys[pos].Izq == nodo {
				break
			}
		}
		if nodo.NodoPadre.Keys[pos].Der.CountKeys() > nodo.getMin() {
			mover := eliminarB(nodo, nodo.NodoPadre.Keys[pos].Der.Keys[0].Valor)
			mover.Der = nodo.NodoPadre.Keys[pos].Der
			mover.Izq = nodo.NodoPadre.Keys[pos].Izq
			nodo.NodoPadre.Keys[pos].Izq = nil
			nodo.NodoPadre.Keys[pos].Der = nil
			nodo.insertKey(nodo.NodoPadre.Keys[pos])
			nodo.NodoPadre.Keys[pos] = mover
		} else if pos > 0 && nodo.NodoPadre.Keys[pos-1].Izq.CountKeys() > nodo.getMin() {
			mover := eliminarB(nodo, nodo.NodoPadre.Keys[pos-1].Izq.getLast().Valor)
			nodo.insertKey(mover)
		}
	}
}

func (nodo *NodoB) getTag(cifrado int) string {
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
			switch cifrado {
			case 0:
				dpi = strconv.Itoa(nodo.Keys[i].Valor)
				correo = nodo.Keys[i].Dato.(*Usuario).Correo
				nombre = nodo.Keys[i].Dato.(*Usuario).Nombre
				password = nodo.Keys[i].Dato.(*Usuario).Password
				cuenta = nodo.Keys[i].Dato.(*Usuario).Cuenta
				break
			case 1:
				tok, _ := fernet.EncryptAndSign([]byte(strconv.Itoa(nodo.Keys[i].Valor)), &fernet.Key{byte(nodo.Keys[i].Valor)})
				dpi = string(tok)[0:25]
				tok, _ = fernet.EncryptAndSign([]byte(nodo.Keys[i].Dato.(*Usuario).Correo), &fernet.Key{byte(nodo.Keys[i].Valor)})
				correo = string(tok)[0:25]
				nombre = nodo.Keys[i].Dato.(*Usuario).Nombre
				cryptoPass := sha256.New()
				cryptoPass.Write([]byte(nodo.Keys[i].Dato.(*Usuario).Password))
				password = strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass.Sum(nil)), "\"", "\\\"")[0:15]
				password = strings.ReplaceAll(password, "}", "\\")
				cuenta = nodo.Keys[i].Dato.(*Usuario).Cuenta
				break
			case 2:
				tok, _ := fernet.EncryptAndSign([]byte(strconv.Itoa(nodo.Keys[i].Valor)), &fernet.Key{byte(nodo.Keys[i].Valor)})
				dpi = string(tok)[0:25]
				tok, _ = fernet.EncryptAndSign([]byte(nodo.Keys[i].Dato.(*Usuario).Correo), &fernet.Key{byte(nodo.Keys[i].Valor)})
				correo = string(tok)[0:25]
				tok, _ = fernet.EncryptAndSign([]byte(nodo.Keys[i].Dato.(*Usuario).Nombre), &fernet.Key{byte(nodo.Keys[i].Valor)})
				nombre = string(tok)[0:25]
				cryptoPass := sha256.New()
				cryptoPass.Write([]byte(nodo.Keys[i].Dato.(*Usuario).Password))
				password = strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass.Sum(nil)), "\"", "\\\"")[0:15]
				password = strings.ReplaceAll(password, "}", "\\")
				tok, _ = fernet.EncryptAndSign([]byte(nodo.Keys[i].Dato.(*Usuario).Cuenta), &fernet.Key{byte(nodo.Keys[i].Valor)})
				cuenta = string(tok)[0:25]
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
				hijos += nodo.Keys[i].Izq.getTag(cifrado)
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
				hijos += nodo.Keys[i-1].Der.getTag(cifrado)
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

func (arbol *ArbolB) Graficar(cifrado int) string {
	return "digraph G{\nnode [shape=Mrecord, color=purple]\n" + arbol.Raiz.getTag(cifrado) + "}"
}
