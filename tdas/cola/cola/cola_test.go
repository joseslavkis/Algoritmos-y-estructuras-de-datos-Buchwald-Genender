package cola_test

import (
	"testing"

	TDACola "tdas/cola"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola0 := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola0.EstaVacia(), "La cola deberia estar vacia luego de crearla, ocurrencia con int")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.VerPrimero()
	})
}

func TestColaVaciaDespuesDeDesencolar(t *testing.T) {
	cola0 := TDACola.CrearColaEnlazada[string]()
	cola0.Encolar("h")
	cola0.Desencolar()
	require.True(t, cola0.EstaVacia(), "La cola deberia estar vacia luego de ser vaciada(desencolada), ocurrencia con cola0/string")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.VerPrimero()
	})

	cola1 := TDACola.CrearColaEnlazada[float64]()
	elementos := []float64{1, 2.78, 3.5, 4, 5.999}
	for _, elem := range elementos {
		cola1.Encolar(elem)
	}
	for range elementos {
		cola1.Desencolar()
	}
	require.True(t, cola1.EstaVacia(), "La cola deberia estar vacia luego de desencolarle todos sus elementos, ocurrencia en cola1")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.VerPrimero()
	})
}

func TestColaVaciaDespuesDeEncolar(t *testing.T) {
	cola0 := TDACola.CrearColaEnlazada[string]()
	elementos := []string{"1", "2", "3", "4", "5", "6", "7", "8", "11"}
	for _, elem := range elementos {
		cola0.Encolar(elem)
	}
	require.False(t, cola0.EstaVacia(), "La cola no deberia estar vacia luego de crearla y encolar una cantidad de elementos x, ocurrencia en cola0")

	cola1 := TDACola.CrearColaEnlazada[int]()
	cola1.Encolar(1)
	require.False(t, cola1.EstaVacia(), "La cola no deberia estar vacia luego de crearla y encolar una cantidad de elementos x, ocurrencia en cola1")
}

func TestVerPrimero(t *testing.T) {
	cola0 := TDACola.CrearColaEnlazada[string]()
	elementos0 := []string{"h", "j", "7", "hola"}
	for _, elem := range elementos0 {
		cola0.Encolar(elem)
	}
	require.Equal(t, elementos0[0], cola0.VerPrimero(), "El primer elemento es incorrecto")

	cola1 := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola1.VerPrimero()
	})
}

func TestDesencolar(t *testing.T) {
	cola0 := TDACola.CrearColaEnlazada[int]()
	elementos0 := []int{1, 2, 3, 4, 5}
	for _, elem := range elementos0 {
		cola0.Encolar(elem)
	}
	for i := 0; i < len(elementos0); i++ {
		require.Equal(t, elementos0[i], cola0.VerPrimero(), "El primero es incorrecto(cola0)")
		require.Equal(t, elementos0[i], cola0.Desencolar(), "Al desencolar la cola, no se desencola el primer elemento(cola0)")
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.VerPrimero()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola0.Desencolar()
	})
}

func PruebasVolumen(t *testing.T) {
	volumen := 100000
	cola0 := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < volumen; i++ {
		cola0.Encolar(i)
		require.Equal(t, 0, cola0.VerPrimero(), "El primero de la cola es incorrecto al encolar, pruebas de volumen(cola0)")
	}
	require.False(t, cola0.EstaVacia(), "La cola no deberia estar vacia despues de encolar x elementos(cola0)")

	for i := 0; i < volumen; i++ {
		require.Equal(t, i, cola0.VerPrimero(), "El primero de la cola es incorrecto, pruebas de volumen(cola0)")
		require.Equal(t, i, cola0.Desencolar(), "Al desencolar la cola, no se desencola el primer elemento, pruebas de volumen(cola0)")
	}
	require.True(t, cola0.EstaVacia(), "La cola deberia estar vacia despues de ser desencolada por completo, prueba de volumen(cola0)")
}
