package lista_test

import (
	"fmt"
	TDALISTA "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	listaPrueba := TDALISTA.CrearListaEnlazada[int]()
	require.True(t, listaPrueba.EstaVacia(), "La lista esta vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaPrueba.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaPrueba.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaPrueba.BorrarPrimero() }, "hay que agregarle elementos a la lista")
	require.Equal(t, 0, listaPrueba.Largo(), "la lista deberia ser de largo 0 luego de ser creada")
}

func TestListaConUnElemento(t *testing.T) {
	listanumero := TDALISTA.CrearListaEnlazada[int]()
	listanumero.InsertarPrimero(35)
	require.Equal(t, 1, listanumero.Largo(), " insertar el numero correcto de elementos")
	require.Equal(t, 35, listanumero.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 35, listanumero.VerUltimo(), "no se inserto correctamente")
	require.False(t, listanumero.EstaVacia(), "la lista esta vacia")
	require.Equal(t, 35, listanumero.BorrarPrimero(), "No se borro correctamente")
	require.True(t, listanumero.EstaVacia(), "La lista esta vacia")
	require.Equal(t, 0, listanumero.Largo(), " insertar el numero correcto de elementos")
}

func TestAgregarElementosLista(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[string]()

	lista.InsertarPrimero("Crepusculo")
	lista.InsertarUltimo("Game Of Thrones")
	require.Equal(t, "Crepusculo", lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, "Game Of Thrones", lista.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, 2, lista.Largo(), " insertar el numero correcto de elementos")
	require.Equal(t, "Crepusculo", lista.BorrarPrimero(), "No se borro correctamente")
	require.Equal(t, "Game Of Thrones", lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, "Game Of Thrones", lista.BorrarPrimero(), "No se borro correctamente")
	require.True(t, lista.EstaVacia(), "La lista esta vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "hay que agregarle elementos a la lista")

}

func TestPruebasInt(t *testing.T) {
	listaInt := TDALISTA.CrearListaEnlazada[int]()

	listaInt.InsertarPrimero(1)
	listaInt.InsertarUltimo(2)
	listaInt.InsertarUltimo(3)
	require.Equal(t, 1, listaInt.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 3, listaInt.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, 3, listaInt.Largo(), " insertar el numero correcto de elementos")
}

func TestPruebaString(t *testing.T) {
	listaGrupo := TDALISTA.CrearListaEnlazada[string]()

	listaGrupo.InsertarPrimero("Jose")
	listaGrupo.InsertarUltimo("Slavkis")
	listaGrupo.InsertarUltimo("Mapu")
	listaGrupo.InsertarUltimo("Canquis")
	require.Equal(t, "Jose", listaGrupo.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, "Canquis", listaGrupo.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, 4, listaGrupo.Largo(), " insertar el numero correcto de elementos")

}

func TestPruebasFloat(t *testing.T) {
	listaValores := TDALISTA.CrearListaEnlazada[float64]()

	listaValores.InsertarPrimero(3.19)
	listaValores.InsertarUltimo(2.14)
	listaValores.InsertarUltimo(34.600)
	listaValores.InsertarUltimo(2.71828)
	require.Equal(t, 3.19, listaValores.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 2.71828, listaValores.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, 4, listaValores.Largo(), " insertar el numero correcto de elementos")

}

func TestPruebasRune(t *testing.T) {
	listaRune := TDALISTA.CrearListaEnlazada[rune]()

	listaRune.InsertarPrimero('\u0061')
	listaRune.InsertarUltimo('\u0062')
	listaRune.InsertarUltimo('\u0063')
	listaRune.InsertarUltimo('\u0064')
	listaRune.InsertarUltimo('\u0065')
	listaRune.InsertarUltimo('\u0066')
	listaRune.InsertarUltimo('\u0067')
	require.Equal(t, '\u0061', listaRune.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, '\u0067', listaRune.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, 7, listaRune.Largo(), " insertar el numero correcto de elementos")

}

func TestPruebasVolumen(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	tamaño := 100000
	lista.InsertarPrimero(0)

	for i := 1; i <= tamaño; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo(), "no se inserto correctamente")
	}
	require.Equal(t, 0, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, tamaño, lista.VerUltimo(), "no se inserto correctamente")
	require.Equal(t, tamaño+1, lista.Largo(), " insertar el numero correcto de elementos")

	for j := 0; j <= tamaño; j++ {
		require.Equal(t, j, lista.VerPrimero(), "No se inserto correctamente")
		lista.BorrarPrimero()
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "hay que agregarle elementos a la lista")
	require.Equal(t, 0, lista.Largo(), " insertar el numero correcto de elementos")

	lista.InsertarUltimo(0)

	for i := 1; i <= tamaño; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero(), "no se inserto correctamente")
	}
	require.Equal(t, 0, lista.VerUltimo(), "No se inserto correctamente")
	require.Equal(t, tamaño, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, tamaño+1, lista.Largo(), " insertar el numero correcto de elementos")
	require.Equal(t, 0, lista.VerUltimo(), "No se inserto correctamente")
	for !lista.EstaVacia() {
		lista.BorrarPrimero()
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "hay que agregarle elementos a la lista")
	require.Equal(t, 0, lista.Largo(), " insertar el numero correcto de elementos")
}

func TestMultiplicarListaConIterar(t *testing.T) {
	listaOriginal := TDALISTA.CrearListaEnlazada[int]()
	listaModificada := TDALISTA.CrearListaEnlazada[int]()

	listaOriginal.InsertarUltimo(1)
	listaOriginal.InsertarUltimo(2)
	listaOriginal.InsertarUltimo(3)
	listaOriginal.InsertarUltimo(4)
	listaOriginal.InsertarUltimo(5)

	iteradorOriginal := listaOriginal.Iterador()
	for iteradorOriginal.HaySiguiente() {
		elem := iteradorOriginal.VerActual()
		iteradorOriginal.Siguiente()

		listaModificada.InsertarUltimo(elem * 2)
	}
	iteradorOriginal = listaOriginal.Iterador()
	iteradorModificado := listaModificada.Iterador()
	for iteradorOriginal.HaySiguiente() && iteradorModificado.HaySiguiente() {
		elemOriginal := iteradorOriginal.VerActual()
		elemModificado := iteradorModificado.VerActual()
		require.Equal(t, elemOriginal*2, elemModificado,
			"El elemento modificado no coincide con el elemento original multiplicado por 2")
		iteradorOriginal.Siguiente()
		iteradorModificado.Siguiente()
	}
}

func TestIterarListaSuma(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	suma := 0
	sumar := func(dato int) bool {
		suma += dato
		return true
	}
	lista.Iterar(sumar)
	sumaEsperada := 6
	require.Equal(t, sumaEsperada, suma, "La suma de los elementos no coincide al utilizar el iterador interno")
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	suma := 0
	visitar := func(dato int) bool {
		suma += dato
		return dato != 3
	}
	lista.Iterar(visitar)
	sumaEsperada := 1 + 2 + 3
	require.Equal(t, sumaEsperada, suma, "La suma de los elementos visitados no coincide")
}
func TestIteradorInsertar(t *testing.T) {

	lista := TDALISTA.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(1)
	require.Equal(t, 1, lista.VerPrimero(), "No se inserto correctamente")

}

func TestIteradorFinal(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[string]()
	lista.InsertarPrimero("harry potter y la camara secreta")
	lista.InsertarUltimo("el señor de los anillos")
	lista.InsertarUltimo("percy jackson y el ladron del rayo")
	lista.InsertarUltimo("divergente")
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar("Twilight")
	require.Equal(t, "Twilight", lista.VerUltimo(), "no se inserto correctamente")
}

func TestIterarMedio(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	i := 0
	x := 0
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		i++
		if iterador.VerActual() == 2 {
			iterador.Insertar(9)
			break
		}
		iterador.Siguiente()
	}

	iterador1 := lista.Iterador()
	for iterador1.HaySiguiente() {
		x++
		if iterador1.VerActual() == 9 {
			require.Equal(t, i, x, "No se insertó en la posición correcta")
			break
		}
		iterador1.Siguiente()
	}

}

func TestRemoverelemento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[float64]()

	lista.InsertarPrimero(8.16)
	lista.InsertarUltimo(44.9)
	lista.InsertarUltimo(95.9)

	iterador := lista.Iterador()

	iterador.Borrar()

	require.Equal(t, 44.9, lista.VerPrimero(), "No se inserto correctamente")

}

func TestRemoverElementoMedioNoEsta(t *testing.T) {

	lista := TDALISTA.CrearListaEnlazada[string]()
	lista.InsertarPrimero("harry potter y la camara secreta")
	lista.InsertarUltimo("el señor de los anillos")
	lista.InsertarUltimo("percy jackson y el ladron del rayo")

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		dato := iterador.VerActual()
		if dato == "el señor de los anillos" {
			iterador.Borrar()
		}
		iterador.Siguiente()
	}

	iterador2 := lista.Iterador()
	i := 0
	for iterador2.HaySiguiente() {
		i++
		if iterador2.VerActual() == "percy jackson y el ladron del rayo" {
			require.Equal(t, "percy jackson y el ladron del rayo", iterador2.VerActual(),
				"No se insertó en la posición correcta")
			require.Equal(t, 2, i, "No se insertó en la posición correcta")
		} else if iterador2.VerActual() == "harry potter y la camara secreta" {
			require.Equal(t, "harry potter y la camara secreta", iterador2.VerActual(),
				"No se insertó en la posición correcta")
			require.Equal(t, 1, i, "No se insertó en la posición correcta")
		}
		iterador2.Siguiente()
	}
	require.Equal(t, 2, lista.Largo(), "El largo de la lista es incorrecto luego de borrar un elemento")

}

func TestIterarListaNumeros(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista1 := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero(),
		"Al insertar un elemento en una lista vacia, no se lo trata como el primer elemento de la misma")
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(8)

	suma := 0
	visitar := func(num int) bool {
		suma += num
		lista1.InsertarUltimo(num)
		return num != 5
	}
	lista.Iterar(visitar)
	require.Equal(t, 17, suma, "La suma de los números menores a 10 debería ser 12")
	require.Equal(t, 5, lista1.VerUltimo(), "El ultimo elemento no es correcto")
}

func TestInsertarAlFinalConIterador(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista1 := TDALISTA.CrearListaEnlazada[int]()
	lista1.InsertarUltimo(1)
	lista1.InsertarUltimo(2)
	lista1.InsertarUltimo(3)

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	elementoInsertado := 4
	iterador.Insertar(elementoInsertado)
	lista1.InsertarUltimo(elementoInsertado)

	require.Equal(t, lista1.VerUltimo(), lista.VerUltimo(),
		"El elemento insertado usando el iterador no coincide con el último elemento de la lista")
}
func TestRemoverUltimoConIterador(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	var ultimo int
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(10)

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		ultimo = iterador.VerActual()
		if ultimo == 10 {
			iterador.Borrar()
			break
		}
		iterador.Siguiente()
	}

	nuevoUltimo := lista.VerUltimo()
	require.Equal(t, 8, nuevoUltimo, "El último elemento no se actualizó correctamente después de removerlo con el iterador")
}

func TestIteradorBorraEnElMedioLuegoIteraConOtroIterador(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Valor 1")
	lista.InsertarUltimo("Valor 2")
	lista.InsertarUltimo("Valor 3")
	lista.InsertarUltimo("Valor 4")
	lista.InsertarUltimo("Valor 5")
	lista.InsertarUltimo("Valor 6")
	lista.InsertarUltimo("Valor 7")

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {

		ultimo := iterador.VerActual()

		if ultimo == "Valor 4" {
			iterador.Borrar()

		}
		iterador.Siguiente()
	}

	require.Equal(t, 6, lista.Largo(), "El largo de la lista es incorrecto luego de borrar un elemento")

	iterador2 := lista.Iterador()

	i := 1
	for iterador2.HaySiguiente() {
		if i == 4 {
			i++
		}
		esperado := fmt.Sprintf("Valor %d", i)
		elemento := iterador2.VerActual()
		require.Equal(t, esperado, elemento)
		i++
		iterador2.Siguiente()
	}

}

func TestIteradorBorraAlFinal(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Valor 1")
	lista.InsertarUltimo("Valor 2")
	lista.InsertarUltimo("Valor 3")
	lista.InsertarUltimo("Valor 4")
	lista.InsertarUltimo("Valor 5")

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual() == "Valor 5" {
			break
		}
		iterador.Siguiente()
	}

	iterador.Borrar()

	require.Equal(t, "Valor 4", lista.VerUltimo(), "El último elemento no es el esperado después de borrar")

}

func TestIteradorBorraListaConUnicoElemento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Valor único")

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Borrar()
	}
	require.True(t, lista.EstaVacia(), "La lista esta vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "hay que agregarle elementos a la lista")

	lista.InsertarPrimero("valor2")
	lista.InsertarUltimo("valor3")
	require.Equal(t, "valor2", lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 2, lista.Largo(), " insertar el numero correcto de elementos")
	require.Equal(t, "valor3", lista.VerUltimo(), "No se inserto correctamente")

}

func TestIteradorBorrarTodos(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	i := 4

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		i--
		iterador.Borrar()
		require.Equal(t, i, lista.Largo(), " insertar el numero correcto de elementos")
	}
	require.True(t, lista.EstaVacia(), "La lista esta vacia")

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "hay que agregarle elementos a la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "hay que agregarle elementos a la lista")

	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	lista.InsertarUltimo(40)
	require.Equal(t, 40, lista.VerUltimo(), "No se inserto correctamente")
	require.Equal(t, 10, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 4, lista.Largo(), " insertar el numero correcto de elementos")
}

func TestEstIteradorBorraEnElMedio(t *testing.T) {

	l := TDALISTA.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	l.InsertarUltimo(4)
	l.InsertarUltimo(5)

	iter := l.Iterador()
	for iter.HaySiguiente() {
		elemento := iter.VerActual()
		if elemento == 3 {
			iter.Borrar()

		}
		iter.Siguiente()
	}

	require.Equal(t, 1, l.VerPrimero(), "No se insertó en la posición correcta")
	require.Equal(t, 5, l.VerUltimo(), "No se insertó en la posición correcta")

}

func TestIteradorInternoNoModificaLaLista(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)

	visitar := func(dato int) bool {
		return dato == 4
	}

	lista.Iterar(visitar)

	require.Equal(t, 6, lista.VerUltimo(), "No se inserto correctamente")
	require.Equal(t, 2, lista.VerPrimero(), "No se inserto correctamente")
	require.Equal(t, 3, lista.Largo(), " insertar el numero correcto de elementos")

}
