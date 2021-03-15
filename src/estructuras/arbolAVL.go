package estructuras

type NodoArbol struct {
	izq, der *NodoArbol
	peso     int
	producto *Producto
}

type Arbol struct {
	raiz *NodoArbol
}

func NewNodoArbol(producto *Producto) *NodoArbol {
	return &NodoArbol{
		izq:      nil,
		der:      nil,
		peso:     0,
		producto: producto,
	}
}

func NewArbol() *Arbol {
	return &Arbol{
		raiz: nil,
	}
}

func rotarII(n *NodoArbol) *NodoArbol {
	n1 := n.izq
	n.izq = n1.der
	n1.der = n
	if n1.peso == -1 {
		n.peso = 0
		n1.peso = 0
	} else {
		n.peso = -1
		n1.peso = 1
	}
	return n1
}

func rotarDD(n *NodoArbol) *NodoArbol {
	n1 := n.der
	n.der = n1.izq
	n1.izq = n
	if n1.peso == 1 {
		n.peso = 0
		n1.peso = 0
	} else {
		n.peso = 1
		n1.peso = -1
	}
	return n1
}

func rotarDI(n *NodoArbol) *NodoArbol {
	n1 := n.der
	n2 := n1.izq
	n2.izq = n
	n1.izq = n2.der
	n2.der = n1
	if n2.peso == 1 {
		n.peso = -1
	} else {
		n.peso = 0
	}
	if n2.peso == -1 {
		n1.peso = 1
	} else {
		n1.peso = 0
	}
	n2.peso = 0
	return n2
}

func rotarID(n *NodoArbol) *NodoArbol {
	n1 := n.izq
	n2 := n1.der
	n1.izq = n2.der
	n2.der = n
	n1.der = n2.izq
	n2.izq = n1
	if n2.peso == 1 {
		n.peso = -1
	} else {
		n.peso = 0
	}
	if n2.peso == -1 {
		n1.peso = 1
	} else {
		n1.peso = 0
	}
	n2.peso = 0
	return n2
}

func Insertar(raiz *NodoArbol, dato *Producto, hc *bool) *NodoArbol {
	var n1 *NodoArbol
	if raiz == nil {
		raiz = NewNodoArbol(dato)
		*hc = true
	} else if dato.Codigo < raiz.producto.Codigo {
		izq := Insertar(raiz.izq, dato, hc)
		raiz.izq = izq
		if *hc {
			switch raiz.peso {
			case 1:
				raiz.peso = 0
				*hc = false
				break
			case 0:
				raiz.peso = -1
				break
			case -1:
				n1 = raiz.izq
				if n1.peso == -1 {
					raiz = rotarII(raiz)
				} else {
					raiz = rotarID(raiz)
				}
				*hc = false
			}
		}
	} else if dato.Codigo > raiz.producto.Codigo {
		der := Insertar(raiz.der, dato, hc)
		raiz.der = der
		if *hc {
			switch raiz.peso {
			case 1:
				n1 := raiz.der
				if n1.peso == 1 {
					raiz = rotarDD(raiz)
				} else {
					raiz = rotarDI(raiz)
				}
				*hc = false
				break
			case 0:
				raiz.peso = 1
				break
			case -1:
				raiz.peso = 0
				*hc = false
				break
			}
		}
	}
	return raiz
}

func (arbol *Arbol) Insertar(dato *Producto) {
	b := false
	arbol.raiz = Insertar(arbol.raiz, dato, &b)
}
