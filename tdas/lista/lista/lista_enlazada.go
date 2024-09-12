package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func nodoCrear[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, siguiente}
}

func CrearListaEnlazada[T any]() *listaEnlazada[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := nodoCrear(dato, nil)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := nodoCrear(dato, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	if lista.primero == nil {
		lista.ultimo = nil
	}
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := nodoCrear(dato, iterador.actual)

	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}
	iterador.actual = nuevoNodo
	if nuevoNodo.siguiente == nil {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	datoEliminado := iterador.actual.dato
	iterador.lista.largo--

	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}

	if iterador.actual.siguiente == nil {
		iterador.lista.ultimo = iterador.anterior
	}

	iterador.actual = iterador.actual.siguiente

	return datoEliminado
}
