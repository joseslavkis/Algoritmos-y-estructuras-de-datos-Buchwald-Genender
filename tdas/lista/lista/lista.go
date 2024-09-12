package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista tiene elementos, false en caso contrario.
	EstaVacia() bool

	//InsertarPrimero coloca un elemento dado en la primer posicion de la lista
	InsertarPrimero(T)

	//InsertarUltimo coloca un elemento dado en la ultima posicion de la lista
	InsertarUltimo(T)

	//BorrarPrimero elimina el primero elemento de la lista y lo devuelve. Panic si la lista esta vacia
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista sin eliminarlo. Panic si la lista esta vacia
	VerPrimero() T

	//VerUlitmo devuelve el ultimo elemento de la lista sin eliminarlo. Panic si la lista esta vacia
	VerUltimo() T

	//Largo devuelve la cantidad de elementos que posee la lista
	Largo() int

	//Iterar recorre y aplica una funcion a los elementos de la lista hasta que esta ultima devuelva false
	Iterar(visitar func(T) bool)

	//Iterador inicializa un iterador
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//Devuelve el elemento actual del iterador. Panic si el iterador ya habia terminado de iterar
	VerActual() T

	//Devuelve true si el iterador puede seguir iterando sobre los elementos de la lista
	HaySiguiente() bool

	//Avanza hacia el siguiente elemento de la lista. Panic si el iterador termino de iterar
	Siguiente()

	//Insertar coloca un elemento en la lista
	Insertar(T)

	//Borrar elimina de la lista el elemento actual del iterador y devuelve dicho elemento.
	// Panic si el iterador termino de iterar
	Borrar() T
}
