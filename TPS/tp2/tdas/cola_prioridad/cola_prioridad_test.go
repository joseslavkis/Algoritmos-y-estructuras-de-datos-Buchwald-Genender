package cola_prioridad_test

import (
	"math/rand"
	"sort"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 100000

type IntSliceDesc []int

func (s IntSliceDesc) Len() int           { return len(s) }
func (s IntSliceDesc) Less(i, j int) bool { return s[i] > s[j] }
func (s IntSliceDesc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func TestHeapVacio(t *testing.T) {
	t.Log("comprueba que al crearse el heap este efectivamente esté vacío")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia(), "El heap no actua como un heap vacío al crearse")
	require.Equal(t, heap.Cantidad(), 0, "La cantidad del heap es incorrecta al crearse")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heap.VerMax()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heap.Desencolar()
	})
}

func TestColaVaciaDespuesDeDesencolar(t *testing.T) {
	t.Log("Verifica que la cola mantenga su estado de vacia luego de desapilar todos sus elementos")
	heap := TDAHeap.CrearHeap(func(a, b float64) int { return int(a) - int(b) })
	elementos := []float64{3.5, 4, 3.5, 2.78, 0.999}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}
	for i := 0; i < len(elementos); i++ {
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia(), "El heap no actua como un heap vacío al crearse")
	require.Equal(t, heap.Cantidad(), 0, "La cantidad del heap es incorrecta al crearse")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heap.VerMax()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heap.Desencolar()
	})
}

func TestColaVaciaDespuesDeEncolar(t *testing.T) {
	t.Log("Verifica que el heap no tenga estado vacío luego de encolar elementos")
	heap := TDAHeap.CrearHeap(func(a, b string) int { return strings.Compare(a, b) })
	elementos := []string{"hola", "como", "andas", "ingeniero", "informático"}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}
	require.False(t, heap.EstaVacia(), "El heap no debería estar vacío luego de encolar elementos")
	require.Equal(t, heap.Cantidad(), 5, "El heap no actualiza correectamente su cantidad al ser encolado")
}

func TestMaxCorrecto(t *testing.T) {
	t.Log("Verifica que el heap no tenga estado vacío luego de encolar elementos")
	heap1 := TDAHeap.CrearHeap(func(a, b string) int { return strings.Compare(a, b) }) //heap de maximos
	elementos := []string{"hola", "como", "andas", "ingeniero", "informático"}
	for _, elem := range elementos {
		heap1.Encolar(elem)
	}
	require.Equal(t, "ingeniero", heap1.VerMax(), "El maximo es incorrecto")
	heap1.Encolar("ingenieros")
	require.Equal(t, "ingenieros", heap1.VerMax(), "El maximo es incorrecto")

	compararStringsMinimo := func(a, b string) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	}
	heap2 := TDAHeap.CrearHeap(compararStringsMinimo)
	for _, elem := range elementos {
		heap2.Encolar(elem)
	}
	require.Equal(t, "andas", heap2.VerMax(), "El minimo es incorrecto")
	heap2.Encolar("aandas")
	require.Equal(t, "aandas", heap2.VerMax(), "El minimo es incorrecto")
}

func TestHeapify(t *testing.T) {
	t.Log("Verifica que el heap se construye correctamente a partir de un array inicial")
	elementos := []int{3, 5, 1, 10, 2}
	heap := TDAHeap.CrearHeapArr(elementos, func(a, b int) int { return a - b })
	require.Equal(t, 10, heap.VerMax(), "El máximo es incorrecto después de heapify")
	require.Equal(t, 5, heap.Cantidad(), "La cantidad del heap es incorrecta después de heapify")
}

func TestRedimensionarHeapEncolar(t *testing.T) {
	t.Log("Verifica que el heap se redimensiona correctamente al encolar algunos elementos")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	elementos := make([]int, 20)
	for i := 0; i < 20; i++ {
		elementos[i] = i
		heap.Encolar(elementos[i])
	}
	require.Equal(t, 20, heap.Cantidad(), "La cantidad del heap es incorrecta después de encolar")
	require.Equal(t, 19, heap.VerMax(), "El máximo es incorrecto después de encolar muchos elementos")
}

func TestRedimensionarHeapDesencolar(t *testing.T) {
	t.Log("Verifica que el heap se redimensiona correctamente al desencolar algunos elementos")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	for i := 0; i < 20; i++ {
		heap.Encolar(i)
	}
	for i := 0; i < 18; i++ {
		heap.Desencolar()
	}
	require.Equal(t, 2, heap.Cantidad(), "La cantidad del heap es incorrecta después de desencolar")
	require.Equal(t, 1, heap.VerMax(), "El máximo es incorrecto después de desencolar muchos elementos")
}

func TestHeapDiferentesTipos(t *testing.T) {
	t.Log("Verifica que el heap funciona correctamente con diferentes tipos de datos")
	heapInt := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	heapInt.Encolar(5)
	heapInt.Encolar(3)
	require.Equal(t, 5, heapInt.VerMax(), "El máximo es incorrecto para enteros")

	heapStr := TDAHeap.CrearHeap(func(a, b string) int { return strings.Compare(a, b) })
	heapStr.Encolar("beta")
	heapStr.Encolar("alpha")
	require.Equal(t, "beta", heapStr.VerMax(), "El máximo es incorrecto para cadenas")

	type Persona struct {
		nombre string
		edad   int
	}
	heapPersona := TDAHeap.CrearHeap(func(a, b Persona) int { return a.edad - b.edad })
	heapPersona.Encolar(Persona{nombre: "Juan", edad: 30})
	heapPersona.Encolar(Persona{nombre: "Ana", edad: 25})
	require.Equal(t, Persona{nombre: "Juan", edad: 30}, heapPersona.VerMax(), "El máximo es incorrecto para estructuras")
}

func TestHeapSortAscendente(t *testing.T) {
	t.Log("Verifica que la primitiva HeapSort ordena correctamente un arreglo")
	arreglo := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	esperado := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	TDAHeap.HeapSort(arreglo, func(a, b int) int { return a - b })
	require.Equal(t, esperado, arreglo, "El arreglo no está ordenado correctamente")
}

func TestHeapsortDescendente(t *testing.T) {
	t.Log("Verifica que HeapSort ordena correctamente un arreglo en orden inverso")
	arreglo := []int{1, 2, 3, 4, 5, 9, 7, 8, 6}
	esperado := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	TDAHeap.HeapSort(arreglo, func(a, b int) int { return b - a })
	require.Equal(t, esperado, arreglo, "El arreglo no está ordenado correctamente")
}

func TestVolumenMax(t *testing.T) {
	t.Log("Verifica el comportamiento del heap con un gran número de elementos (heap de máximos)")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	for i := 0; i < _VOLUMEN; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, _VOLUMEN, heap.Cantidad(), "La cantidad del heap es incorrecta después de encolar muchos elementos")
	require.Equal(t, _VOLUMEN-1, heap.VerMax(), "El máximo es incorrecto después de encolar muchos elementos")

	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar(), "El elemento desencolado es incorrecto")
	}
	require.True(t, heap.EstaVacia(), "El heap no está vacío después de desencolar todos los elementos")

	vec_aux := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		num := rand.Intn(10000) + 1
		vec_aux[i] = num
		heap.Encolar(num)
	}

	sort.Ints(vec_aux)

	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, vec_aux[i], heap.Desencolar(), "El elemento desencolado es incorrecto")
	}
	require.True(t, heap.EstaVacia(), "El heap no está vacío después de desencolar todos los elementos")
}

func TestVolumenMin(t *testing.T) {
	t.Log("Verifica el comportamiento del heap con un gran número de elementos (heap de máximos)")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return b - a })
	for i := 0; i < _VOLUMEN; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, _VOLUMEN, heap.Cantidad(), "La cantidad del heap es incorrecta después de encolar muchos elementos")
	require.Equal(t, 0, heap.VerMax(), "El máximo es incorrecto después de encolar muchos elementos")

	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, i, heap.Desencolar(), "El elemento desencolado es incorrecto")
	}
	require.True(t, heap.EstaVacia(), "El heap no está vacío después de desencolar todos los elementos")

	vec_aux := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		num := rand.Intn(10000) + 1
		vec_aux[i] = num
		heap.Encolar(num)
	}

	sort.Ints(vec_aux)

	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, vec_aux[i], heap.Desencolar(), "El elemento desencolado es incorrecto")
	}
	require.True(t, heap.EstaVacia(), "El heap no está vacío después de desencolar todos los elementos")
}

func TestVolumenHeapsortAscendente(t *testing.T) {
	t.Log("Verifica el comportamiento de heapsort (heap de máximos), se verifica tanto si se recibe un vector ya ordenado como también desordenado")
	vector_a_ordenar := make([]int, _VOLUMEN)
	vector_aux := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		vector_a_ordenar[i] = i
		vector_aux[i] = i
	}
	TDAHeap.HeapSort(vector_a_ordenar, func(a, b int) int { return a - b })
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, vector_aux[i], vector_a_ordenar[i], "El elemento ordenado es incorrecto luego de hacer heapsort")
	}

	vector_a_ordenar1 := make([]int, _VOLUMEN)
	vector_aux1 := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		num := rand.Intn(10000) + 1
		vector_a_ordenar1[i] = num
		vector_aux1[i] = num
	}
	TDAHeap.HeapSort(vector_a_ordenar1, func(a, b int) int { return a - b })
	sort.Ints(vector_aux1)
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, vector_aux1[i], vector_a_ordenar1[i], "El elemento ordenado es incorrecto luego de hacer heapsort")
	}
}

func TestVolumenHeapsortDescendente(t *testing.T) {
	t.Log("Verifica el comportamiento de heapsort (heap de máximos), se verifica tanto si se recibe un vector ya ordenado como también desordenado")
	vector_a_ordenar := make([]int, _VOLUMEN)
	vector_aux := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		vector_a_ordenar[i] = i
		vector_aux[i] = i
	}
	TDAHeap.HeapSort(vector_a_ordenar, func(a, b int) int { return b - a })
	for i := _VOLUMEN - 1; i < 0; i++ {
		require.Equal(t, vector_aux[i], vector_a_ordenar[i], "El elemento ordenado es incorrecto luego de hacer heapsort")
	}

	vector_a_ordenar1 := make([]int, _VOLUMEN)
	vector_aux1 := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		num := rand.Intn(10000) + 1
		vector_a_ordenar1[i] = num
		vector_aux1[i] = num
	}
	TDAHeap.HeapSort(vector_a_ordenar1, func(a, b int) int { return b - a })
	sort.Sort(IntSliceDesc(vector_aux1))
	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, vector_aux1[i], vector_a_ordenar1[i], "El elemento ordenado es incorrecto luego de hacer heapsort")
	}
}

func TestDuplicados(t *testing.T) {
	t.Log("Verifica que el heap maneja correctamente los elementos duplicados")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	elementos := []int{5, 1, 3, 5, 2, 5, 4}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}
	require.Equal(t, 7, heap.Cantidad(), "La cantidad del heap es incorrecta después de encolar elementos duplicados")
	require.Equal(t, 5, heap.VerMax(), "El máximo es incorrecto con elementos duplicados")
}

func TestHeapStruct(t *testing.T) {
	t.Log("Verifica que el heap maneja correctamente estructuras complejas")
	type Persona struct {
		nombre string
		edad   int
	}
	heap := TDAHeap.CrearHeap(func(a, b Persona) int { return a.edad - b.edad })
	personas := []Persona{
		{"Juan", 30},
		{"Ana", 25},
		{"Pedro", 35},
		{"Maria", 28},
	}
	for _, persona := range personas {
		heap.Encolar(persona)
	}
	require.Equal(t, Persona{"Pedro", 35}, heap.VerMax(), "El máximo es incorrecto para estructuras complejas")
}

func TestMultiplesRedimensionamientos(t *testing.T) {
	t.Log("Verifica el comportamiento del heap después de múltiples redimensionamientos")
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	for i := 0; i < 50; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, 50, heap.Cantidad(), "La cantidad del heap es incorrecta después de encolar muchos elementos")
	require.Equal(t, 49, heap.VerMax(), "El máximo es incorrecto después de encolar muchos elementos")

	for i := 0; i < 45; i++ {
		heap.Desencolar()
	}
	require.Equal(t, 5, heap.Cantidad(), "La cantidad del heap es incorrecta después de desencolar muchos elementos")
	require.Equal(t, 4, heap.VerMax(), "El máximo es incorrecto después de desencolar muchos elementos")
}
