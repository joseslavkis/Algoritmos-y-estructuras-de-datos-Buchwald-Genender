package pila

/* Definición del struct pila proporcionado por la cátedra. */
const _LARGOINICIALPILA = 15
const _PROPORCIONCANTIDADCAPACIDAD = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _LARGOINICIALPILA), cantidad: 0}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		nuevaCapacidad := pila.cantidad * 2
		redimensionarPila(pila, nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad*_PROPORCIONCANTIDADCAPACIDAD <= cap(pila.datos) && cap(pila.datos) > _LARGOINICIALPILA {
		nuevaCapacidad := int(cap(pila.datos) / 2)
		redimensionarPila(pila, nuevaCapacidad)
	}
	elemento := pila.datos[pila.cantidad-1]
	pila.cantidad--
	return elemento
}

func redimensionarPila[T any](pila *pilaDinamica[T], capacidad int) {
	nuevosDatos := make([]T, capacidad)
	copy(nuevosDatos, pila.datos[:pila.cantidad])
	pila.datos = nuevosDatos
}
