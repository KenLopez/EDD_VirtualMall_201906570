package estructuras

type NodoArbol struct {
	izq, der   *NodoArbol
	dato, peso int
}

type Arbol struct {
	raiz *NodoArbol
}

func main() {

}
