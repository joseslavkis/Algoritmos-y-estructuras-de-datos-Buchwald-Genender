package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila0 := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila0.EstaVacia(), "La pila debería estar vacía luego de crearla, ocurrencia con int")
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.Desapilar()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.VerTope()
	})
}

func TestPilaVaciaDespuesDeDesapilar(t *testing.T) {

	pila0 := TDAPila.CrearPilaDinamica[string]()
	pila0.Apilar("h")
	pila0.Desapilar()
	require.True(t, pila0.EstaVacia(), "La pila debería estar vacía luego de ser vaciada(desapilada), ocurrencia con pila0/string")
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.Desapilar()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.VerTope()
	})
	pila1 := TDAPila.CrearPilaDinamica[float64]()
	elementos := []float64{1, 2.78, 3.5, 4, 5.999}
	for i := 0; i < len(elementos); i++ {

		pila1.Apilar(elementos[i])
	}
	for i := 0; i < len(elementos); i++ {
		pila1.Desapilar()
	}
	require.True(t, pila1.EstaVacia(), "La pila deberia estar vacía luego de desapilarle todos sus elementos, ocurrencia en pila1")
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.Desapilar()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.VerTope()
	})
}

func TestPilaVaciaDespuesDeApilar(t *testing.T) {
	pila0 := TDAPila.CrearPilaDinamica[string]()
	elementos := []string{"1", "2", "3", "4", "5", "6", "7", "8", "11"}
	for i := 0; i < len(elementos); i++ {
		pila0.Apilar(elementos[i])
	}
	require.False(t, pila0.EstaVacia(), "La pila no deberia estar vacía luego de crearla y apilar una cantidad de elementos x, ocurrencia en pila0")

	pila1 := TDAPila.CrearPilaDinamica[int]()
	pila1.Apilar(1)
	require.False(t, pila1.EstaVacia(), "La pila no deberia estar vacía luego de crearla y apilar una cantidad de elementos x, ocurrencia en pila1")
}

func TestVerTope(t *testing.T) {
	pila0 := TDAPila.CrearPilaDinamica[string]()
	elementos0 := []string{"h", "j", "7", "hola"}
	for i := 0; i < len(elementos0); i++ {
		pila0.Apilar(elementos0[i])
		require.Equal(t, elementos0[i], pila0.VerTope(), "El tope no se actualiza correctamente al apilar")
	}
	pila1 := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila1.VerTope()
	})
}

func TestDesapilar(t *testing.T) {
	pila0 := TDAPila.CrearPilaDinamica[int]()
	elementos0 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(elementos0); i++ {
		pila0.Apilar(elementos0[i])
	}
	for i := len(elementos0) - 1; i >= 0; i-- {
		require.Equal(t, elementos0[i], pila0.VerTope(), "El tope es incorrecto(pila0)")
		require.Equal(t, elementos0[i], pila0.Desapilar(), "Al desapilar la pila, no se desapila el último elemento(pila0)")
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.VerTope()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila0.Desapilar()
	})
}

func PruebasVolumen(t *testing.T) {
	volumen := 100000
	pila0 := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < volumen; i++ {
		pila0.Apilar(i)
		require.Equal(t, i, pila0.VerTope(), "El tope de la pila es incorrecto al apilar, pruebas de volumen(pila0)")
	}
	require.False(t, pila0.EstaVacia(), "La pila no debería estar vacía después de apilar x elementos(pila0)")

	for i := volumen - 1; i >= 0; i-- {
		require.Equal(t, i, pila0.VerTope(), "El tope de la pila es incorrecto, pruebas de volumen(pila0)")
		require.Equal(t, i, pila0.Desapilar(), "Al desapilar la pila no se desapila el último elemento, pruebas de volumen(pila0)")
	}
	require.True(t, pila0.EstaVacia(), "La pila debería estar vacía luego de ser desapilada por completo, prueba de volumen(pila0)")
}
