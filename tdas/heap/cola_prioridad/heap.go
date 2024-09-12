package cola_prioridad

type funcionCmp[T any] func(T, T) int

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   funcionCmp[T]
}

const (
	_LARGO_INICIAL      = 10
	_FACTOR_REDIMENSION = 2
	_FACTOR_ACHICAR     = 4
)

func CrearHeap[T any](funcion_cmp funcionCmp[T]) ColaPrioridad[T] {
	return &colaConPrioridad[T]{datos: make([]T, _LARGO_INICIAL), cmp: funcion_cmp}
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.Cantidad() == 0
}

func (heap *colaConPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Encolar(dato T) {
	if heap.cant == cap(heap.datos) {
		heap.redimensionar(cap(heap.datos) * _FACTOR_REDIMENSION)
	}
	heap.datos[heap.cant] = dato
	heap.upHeap(heap.cant)
	heap.cant++
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	if heap.cant <= cap(heap.datos)/_FACTOR_ACHICAR && cap(heap.datos) > _LARGO_INICIAL {
		heap.redimensionar(cap(heap.datos) / _FACTOR_REDIMENSION)
	}
	elementoBorrado := heap.datos[0]
	heap.cant--
	swap(&heap.datos[heap.cant], &heap.datos[0])
	heap.downHeap(0)
	return elementoBorrado
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp funcionCmp[T]) ColaPrioridad[T] {
	var capacidad int
	if len(arreglo) < _LARGO_INICIAL {
		capacidad = _LARGO_INICIAL
	} else {
		capacidad = len(arreglo) * _FACTOR_REDIMENSION
	}
	heap := &colaConPrioridad[T]{datos: make([]T, capacidad), cmp: funcion_cmp, cant: len(arreglo)}
	copy(heap.datos, arreglo)
	heap.heapify()
	return heap
}

func HeapSort[T any](elementos []T, funcion_cmp funcionCmp[T]) {
	heap := &colaConPrioridad[T]{datos: elementos, cmp: funcion_cmp, cant: len(elementos)}
	heap.heapify()

	for i := heap.cant - 1; i > 0; i-- {
		swap(&heap.datos[0], &heap.datos[i])
		heap.cant--
		heap.downHeap(0)
	}
	heap.cant = len(elementos)
}

func (heap *colaConPrioridad[T]) upHeap(posicion int) {
	if posicion == 0 {
		return
	}
	padre := (posicion - 1) / 2
	if heap.cmp(heap.datos[posicion], heap.datos[padre]) > 0 {
		swap(&heap.datos[posicion], &heap.datos[padre])
		heap.upHeap(padre)
	}
}

func (heap *colaConPrioridad[T]) max(posicion int) int {
	hijoIzq, hijoDer := 2*posicion+1, 2*posicion+2

	if hijoIzq >= heap.cant && hijoDer >= heap.cant {
		return posicion
	}

	if hijoDer >= heap.cant {
		if heap.cmp(heap.datos[hijoIzq], heap.datos[posicion]) > 0 {
			return hijoIzq
		}
		return posicion
	}

	maxPos := posicion
	if heap.cmp(heap.datos[hijoIzq], heap.datos[maxPos]) > 0 {
		maxPos = hijoIzq
	}
	if heap.cmp(heap.datos[hijoDer], heap.datos[maxPos]) > 0 {
		maxPos = hijoDer
	}

	return maxPos
}

func (heap *colaConPrioridad[T]) downHeap(posicion int) {
	for {
		posMayorHijo := heap.max(posicion)
		if posMayorHijo == posicion {
			break
		}
		swap(&heap.datos[posicion], &heap.datos[posMayorHijo])
		posicion = posMayorHijo
	}
}

func (heap *colaConPrioridad[T]) heapify() {
	for i := (heap.cant / 2) - 1; i >= 0; i-- {
		heap.downHeap(i)
	}
}

func (heap *colaConPrioridad[T]) redimensionar(nuevoTamanio int) {
	nuevosDatos := make([]T, nuevoTamanio)
	copy(nuevosDatos, heap.datos[:heap.cant])
	heap.datos = nuevosDatos
}

func swap[T any](elem1, elem2 *T) {
	*elem1, *elem2 = *elem2, *elem1
}
