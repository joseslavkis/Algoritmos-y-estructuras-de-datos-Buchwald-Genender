package cola

type nodoCola[T any] struct {
	valor     T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato, nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.valor
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	Desencolado := cola.VerPrimero()
	cola.primero = cola.primero.siguiente
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return Desencolado
}

func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := nodoCrear(valor)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}
