package estructuras

type Cola struct {
	Frente *Nodo
}

func NewCola() *Cola {
	return &Cola{Frente: nil}
}

func (cola *Cola) Queue(dato interface{}) {
	nodo := &Nodo{
		Contenido: dato,
		Next:      nil,
	}
	if cola.Frente == nil {
		cola.Frente = nodo
	} else {
		aux := cola.Frente
		for aux.Next != nil {
			aux = aux.Next
		}
		aux.Next = nodo
	}
}

func (cola *Cola) Dequeue() *Nodo {
	if cola.Frente == nil {
		return nil
	}
	nodo := cola.Frente
	cola.Frente = cola.Frente.Next
	return nodo
}
