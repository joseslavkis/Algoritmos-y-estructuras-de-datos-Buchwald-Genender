package ejercicios_test

import (
	"strings"
	"testing"
	"tp0/ejercicios"

	"github.com/stretchr/testify/require"
)

func TestSwap(t *testing.T) {
	a := 10
	b := 5
	ejercicios.Swap(&a, &b)
	require.EqualValues(t, 5, a)
	require.EqualValues(t, 10, b)

	ejercicios.Swap(&a, &b)
	require.EqualValues(t, 10, a)
	require.EqualValues(t, 5, b)

	ejercicios.Swap(&a, &a)
	require.EqualValues(t, 10, a)

	a = 3
	b = 3
	ejercicios.Swap(&a, &b)
	require.EqualValues(t, 3, a)
	require.EqualValues(t, 3, b)

	a = 0
	b = 10
	ejercicios.Swap(&a, &b)
	require.EqualValues(t, 10, a)
	require.EqualValues(t, 0, b)

	ejercicios.Swap(&a, &b)
	require.EqualValues(t, 0, a)
	require.EqualValues(t, 10, b)
}

func TestMaximo(t *testing.T) {
	var (
		vacio             = []int{}
		unico             = []int{1}
		secuenciales      = []int{1, 2, 3, 4, 5}
		negativos         = []int{-2000, -1500, -1000, -3000}
		algunosNegativos  = []int{8, -10, 27, 3, -50}
		algunosNegUltimo  = []int{8, -10, 27}
		algunosNegPrimero = []int{100, -10, 27}
		repetidos         = []int{1, 2, 3, 4, 3, 4, 2, 1}
		todosCero         = []int{0, 0, 0, 0, 0}
	)

	require.Equal(t, -1, ejercicios.Maximo(vacio), "Si el arreglo tiene largo 0, devuelve -1")
	require.Equal(t, 0, ejercicios.Maximo(unico), "El maximo de un arreglo de un unico elemento es ese unico")
	require.Equal(t, 4, ejercicios.Maximo(secuenciales), "El maximo de un arreglo creciente es el ultimo")
	require.Equal(t, 2, ejercicios.Maximo(negativos), "Se encuentra maximo en arreglo de negativos")
	require.Equal(t, 2, ejercicios.Maximo(algunosNegativos), "Se encuentra el maximo en un arreglo con positivos y negativos")
	require.Equal(t, ejercicios.Maximo(algunosNegativos), ejercicios.Maximo(algunosNegUltimo), "Se encuentra correctamente el maximo si acortemos largo")
	require.Equal(t, 0, ejercicios.Maximo(algunosNegPrimero), "Maximo al inicio")
	require.Equal(t, 3, ejercicios.Maximo(repetidos), "El maximo de un vector con maximo repetido es la primera ocurrencia")
	require.Equal(t, 0, ejercicios.Maximo(todosCero), "El maximo de un vector con todos ceros es el primero")

	volumen := make([]int, 100000)
	for i := range volumen {
		volumen[i] = i + 1
	}
	require.Equal(t, 99999, ejercicios.Maximo(volumen), "Si el arreglo es muy grande, igualmente funcioa y no demora mucho en calcular")
}

func TestComparar(t *testing.T) {
	{
		a := []int{10}
		b := []int{10}
		require.Equal(t, 0, ejercicios.Comparar(a, b))
	}
	{
		a := []int{10, 20}
		b := []int{10, 20}
		require.Equal(t, 0, ejercicios.Comparar(a, b))
	}
	{
		a := []int{}
		b := []int{}
		require.Equal(t, 0, ejercicios.Comparar(a, b))
	}
	{
		a := []int{}
		b := []int{1, 2, 3}
		require.Equal(t, -1, ejercicios.Comparar(a, b))
		require.Equal(t, 1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{0, 2, 3}
		require.Equal(t, 1, ejercicios.Comparar(a, b))
		require.Equal(t, -1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{0, 2, 4}
		require.Equal(t, 1, ejercicios.Comparar(a, b))
		require.Equal(t, -1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{1, 2}
		require.Equal(t, 1, ejercicios.Comparar(a, b))
		require.Equal(t, -1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{1, 2, 2, 4}
		require.Equal(t, 1, ejercicios.Comparar(a, b))
		require.Equal(t, -1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 3, 3, 2}
		require.Equal(t, -1, ejercicios.Comparar(a, b))
		require.Equal(t, 1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{3, 2, 1, 0}
		require.Equal(t, -1, ejercicios.Comparar(a, b))
		require.Equal(t, 1, ejercicios.Comparar(b, a))
	}
	{
		a := []int{1, 2, 3}
		b := []int{3, 2, 1}
		require.Equal(t, -1, ejercicios.Comparar(a, b))
		require.Equal(t, 1, ejercicios.Comparar(b, a))
	}
}

func TestSeleccion(t *testing.T) {
	var (
		vacio = []int{}
		unico = []int{8}
		vec1  = []int{3, 5, 4, 2, 1}
		vec2  = []int{4, 8, 15, 16, 23, 42}
		vec3  = []int{-38, -46, -65, -78}
	)

	ejercicios.Seleccion(vacio)
	require.Equal(t, []int{}, vacio, "No debe romperse por no tener elemenots")
	ejercicios.Seleccion(unico)
	require.Equal(t, []int{8}, unico, "El arreglo con un solo elemento debe quedar igual")
	ejercicios.Seleccion(vec1)
	require.Equal(t, []int{1, 2, 3, 4, 5}, vec1, "Se ordena correctamente un arreglo")
	ejercicios.Seleccion(vec2)
	require.Equal(t, []int{4, 8, 15, 16, 23, 42}, vec2, "Un arreglo ya ordenado no cambia su orden")
	ejercicios.Seleccion(vec3)
	require.Equal(t, []int{-78, -65, -46, -38}, vec3, "El algoritmo funciona con números negativos")
}

func TestSuma(t *testing.T) {
	var (
		vacio = []int{}
		unico = []int{8}
		vec1  = []int{3, 5, 4, 2, 1}
		vec2  = []int{4, 8, 15, 16, 23, 42}
		vec3  = []int{-38, -46, -65, -78}
		vec4  = []int{10, 9, -15, 0, 7, -12, 1}
	)

	require.Equal(t, 0, ejercicios.Suma(vacio))
	require.Equal(t, 8, ejercicios.Suma(unico))
	require.Equal(t, 15, ejercicios.Suma(vec1))
	require.Equal(t, 108, ejercicios.Suma(vec2))
	require.Equal(t, -227, ejercicios.Suma(vec3))
	require.Equal(t, 0, ejercicios.Suma(vec4))
}

func TestEsCadenaCapicua(t *testing.T) {
	require.True(t, ejercicios.EsCadenaCapicua(""))
	require.True(t, ejercicios.EsCadenaCapicua("a"))
	require.True(t, ejercicios.EsCadenaCapicua("ana"))
	require.True(t, ejercicios.EsCadenaCapicua("neuquen"))

	// Agarramos palíndromos, les sacamos los espacios para que queden capicua
	capicua := strings.Replace("anita lava la tina", " ", "", -1)
	require.True(t, ejercicios.EsCadenaCapicua(capicua))
	capicua = strings.Replace("son robos o sobornos", " ", "", -1)
	require.True(t, ejercicios.EsCadenaCapicua(capicua))

	require.False(t, ejercicios.EsCadenaCapicua("EstoNoEsCapicua"))
	require.False(t, ejercicios.EsCadenaCapicua("Neuquen"))
	require.False(t, ejercicios.EsCadenaCapicua("palijlap"))
	require.False(t, ejercicios.EsCadenaCapicua("ab"))

	require.True(t, ejercicios.EsCadenaCapicua(" Espacios soicapsE "))
	require.False(t, ejercicios.EsCadenaCapicua(" EE"))
}
