package diccionario

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

const (
	_TABLA_LARGO_INICIAL   = 10
	_FACTOR_CAMBIO         = 2
	_FACTOR_CARGA_ACHICAR  = 0.2
	_FACTOR_CARGA_AUMENTAR = 0.7
)
const (
	_VACIO estadoCelda = iota
	_OCUPADO
	_BORRADO
)

type estadoCelda int

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoCelda
}

type hashCerrado[K comparable, V any] struct {
	tamanio  int
	cantidad int
	tabla    []celdaHash[K, V]
	borrados int
}

type iterHashCerrado[K comparable, V any] struct {
	hash   *hashCerrado[K, V]
	actual int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{tabla: make([]celdaHash[K, V], _TABLA_LARGO_INICIAL), cantidad: 0, tamanio: _TABLA_LARGO_INICIAL, borrados: 0}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	pos, encontrado := hash.buscar(clave)
	return encontrado && hash.tabla[pos].estado == _OCUPADO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion, encontrado := hash.buscar(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	factor := float64(hash.cantidad+hash.borrados+1) / float64(hash.tamanio)
	if factor >= _FACTOR_CARGA_AUMENTAR {
		hash.redimensionar(hash.tamanio * _FACTOR_CAMBIO)
	}
	posicion, encontrado := hash.buscar(clave)
	hash.tabla[posicion].clave, hash.tabla[posicion].dato = clave, dato

	if !encontrado {
		hash.tabla[posicion].estado = _OCUPADO
		hash.cantidad++
	}
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	factor := float64(hash.cantidad-1) / float64(hash.tamanio)
	if factor <= _FACTOR_CARGA_ACHICAR {
		hash.redimensionar(hash.tamanio / _FACTOR_CAMBIO)
	}
	posicion, encontrado := hash.buscar(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	hash.tabla[posicion].estado = _BORRADO
	hash.borrados++
	hash.cantidad--
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == _OCUPADO {
			if !visitar(celda.clave, celda.dato) {
				break
			}
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

	for i := 0; i < len(hash.tabla); i++ {
		if hash.tabla[i].estado == _OCUPADO {
			return &iterHashCerrado[K, V]{hash: hash, actual: i}
		}
	}
	return &iterHashCerrado[K, V]{hash: hash, actual: len(hash.tabla)}
}

func (iterador *iterHashCerrado[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.actual++
	for iterador.HaySiguiente() {
		if iterador.hash.tabla[iterador.actual].estado == _OCUPADO {
			return
		}
		iterador.actual++
	}
}

func (iterador *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.hash.tabla[iterador.actual].clave, iterador.hash.tabla[iterador.actual].dato
}

func (iterador *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iterador.actual != iterador.hash.tamanio
}

func (hash *hashCerrado[K, V]) buscar(clave K) (int, bool) {

	pos := int(hashing(convertirABytes(clave))) % hash.tamanio

	for hash.tabla[pos].estado != _VACIO {
		if hash.tabla[pos].clave == clave && hash.tabla[pos].estado == _OCUPADO {
			return pos, true
		}
		pos++
		if pos == hash.tamanio {
			pos = 0
		}
	}
	return pos, false
}

// xxhash64
func hashing(data []byte) uint64 {
	length := len(data)
	var (
		i      int
		h64    uint64
		prime1 uint64
		prime2 uint64
		prime3 uint64
		prime4 uint64
		prime5 uint64
	)
	prime1 = 11400714785074694791
	prime2 = 14029467366897019727
	prime3 = 1609587929392839161
	prime4 = 9650029242287828579
	prime5 = 2870177450012600261

	if length >= 32 {
		end := length - 32
		for i = 0; i <= end; i += 32 {
			v1 := binary.LittleEndian.Uint64(data[i:])
			v2 := binary.LittleEndian.Uint64(data[i+8:])
			v3 := binary.LittleEndian.Uint64(data[i+16:])
			v4 := binary.LittleEndian.Uint64(data[i+24:])
			h64 += v1 * prime2
			h64 = (bits.RotateLeft64(h64, 31)*prime1 + h64*prime4) + v2*prime2
			h64 = (bits.RotateLeft64(h64, 27)*prime1 + h64*prime4) + v3*prime2
			h64 = (bits.RotateLeft64(h64, 33)*prime1 + h64*prime4) + v4*prime2
		}
	}

	if i <= length-16 {
		v1 := binary.LittleEndian.Uint64(data[i:])
		v2 := binary.LittleEndian.Uint64(data[i+8:])
		h64 += v1 * prime2
		h64 = (bits.RotateLeft64(h64, 31) * prime1) + v2*prime2
		i += 16
	}

	if i <= length-8 {
		v := binary.LittleEndian.Uint64(data[i:])
		h64 ^= (bits.RotateLeft64(v*prime2, 37) * prime1)
		h64 = (h64*prime1 + prime4) * prime2
		i += 8
	}

	if i <= length-4 {
		v := binary.LittleEndian.Uint32(data[i:])
		h64 ^= uint64(v) * prime1
		h64 = (h64*prime2 + prime3)
		i += 4
	}

	for i < length {
		h64 ^= uint64(data[i]) * prime5
		h64 = (h64*prime2 + prime1)
		i++
	}

	h64 ^= uint64(length)

	h64 ^= h64 >> 33
	h64 *= prime2
	h64 ^= h64 >> 29
	h64 *= prime3
	h64 ^= h64 >> 32

	return h64 & ((1 << 63) - 1)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hashCerrado[K, V]) redimensionar(nuevoTamanio int) {
	nuevaTabla := make([]celdaHash[K, V], nuevoTamanio)

	antiguoTabla := hash.tabla
	hash.tabla = nuevaTabla
	hash.tamanio = nuevoTamanio
	hash.cantidad = 0
	hash.borrados = 0
	for _, celda := range antiguoTabla {
		if celda.estado == _OCUPADO {
			hash.Guardar(celda.clave, celda.dato)
		}
	}
}
