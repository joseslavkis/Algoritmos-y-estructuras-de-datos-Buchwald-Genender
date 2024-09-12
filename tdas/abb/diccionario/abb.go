package diccionario

import (
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterAbb[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

// bien
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cmp: funcion_cmp}
}

// bien
func nodoCrear[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

// bien
func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	referenciaANodo := arbol.buscar(clave)
	return *referenciaANodo != nil
}

// bien
func (arbol *abb[K, V]) Obtener(clave K) V {
	referenciaANodo := arbol.buscar(clave)
	if *referenciaANodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*referenciaANodo).dato
}

// bien
func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	referenciaANodo := arbol.buscar(clave)
	if *referenciaANodo == nil {
		*referenciaANodo = nodoCrear(clave, dato)
		arbol.cantidad++
		return
	}
	(*referenciaANodo).dato = dato
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	referenciaANodo := arbol.buscar(clave)
	if *referenciaANodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return arbol.borrar(referenciaANodo)
}

// in-order(se visitan izq-yo-der) esto es para visitar los elementos de forma ordenada
func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(arbol.raiz, desde, hasta, visitar, arbol.cmp)
}

// si o si esta bien
func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

// lo mismo
func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func (arbol *abb[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{arbol: arbol, pila: TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](), desde: desde, hasta: hasta}
	iter.apilarHastaMinimo(arbol.raiz)

	return iter
}

func (iter *iterAbb[K, V]) apilarHastaMinimo(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if (iter.desde == nil || iter.arbol.cmp(nodo.clave, *iter.desde) >= 0) && (iter.hasta == nil || iter.arbol.cmp(nodo.clave, *iter.hasta) <= 0) {
			iter.pila.Apilar(nodo)
			nodo = nodo.izq
		} else if iter.desde != nil && iter.arbol.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
		} else {
			nodo = nodo.izq
		}
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.Desapilar()
	if actual.der != nil {
		iter.apilarHastaMinimo(actual.der)
	}
}

// bien
func (arbol *abb[K, V]) buscar(clave K) **nodoAbb[K, V] {
	return arbol.buscarInterno(clave, arbol.raiz, &(arbol.raiz))
}

func (arbol *abb[K, V]) buscarInterno(clave K, nodo *nodoAbb[K, V], referenciaANodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if nodo == nil || arbol.cmp(clave, nodo.clave) == 0 {
		return referenciaANodo
	} else if arbol.cmp(clave, nodo.clave) > 0 {
		return arbol.buscarInterno(clave, nodo.der, &nodo.der)
	} else {
		return arbol.buscarInterno(clave, nodo.izq, &nodo.izq)
	}
}

func (arbol *abb[K, V]) borrar(referenciaANodo **nodoAbb[K, V]) V {
	nodo := *referenciaANodo
	elementoBorrado := nodo.dato

	if nodo.izq != nil && nodo.der != nil {
		sustituto := sustituir(nodo.der)
		nodo.clave = sustituto.clave
		nodo.dato = sustituto.dato
		arbol.borrar(&sustituto)
		arbol.cantidad++
	} else if nodo.izq != nil {
		*referenciaANodo = nodo.izq
	} else if nodo.der != nil {
		*referenciaANodo = nodo.der
	} else {
		*referenciaANodo = nil
	}
	arbol.cantidad--
	return elementoBorrado
}
func iterarRango[K comparable, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || cmp(nodo.clave, *desde) >= 0 {
		if !iterarRango(nodo.izq, desde, hasta, visitar, cmp) {
			return false
		}
	}
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || cmp(nodo.clave, *hasta) <= 0 {
		if !iterarRango(nodo.der, desde, hasta, visitar, cmp) {
			return false
		}
	}
	return true
}

// bien
func sustituir[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izq == nil {
		return nodo
	}
	return sustituir(nodo.izq)
}
