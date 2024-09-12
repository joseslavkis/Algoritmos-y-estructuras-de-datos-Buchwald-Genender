package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDAAbb "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararStrings(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")

	dic := TDAAbb.CrearABB[string, string](compararStrings)

	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un árbol binario de búsqueda vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")

	dic := TDAAbb.CrearABB[string, string](compararStrings)
	require.NotNil(t, dic)

	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDAAbb.CrearABB[int, string](funcionComparacion)
	require.NotNil(t, dicNum)

	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que árbol binario de búsqueda con un elemento tiene esa Clave, únicamente")
	dic := TDAAbb.CrearABB[string, int](compararStrings)
	require.NotNil(t, dic)

	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Venezuela"
	clave2 := "Argentina"
	clave3 := "España"
	valor1 := "Italia"
	valor2 := "Grecia"
	valor3 := "Mexico"

	// Crear un nuevo árbol
	arbol := TDAAbb.CrearABB[string, string](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	require.False(t, arbol.Pertenece(clave1))
	require.False(t, arbol.Pertenece(clave2))
	require.False(t, arbol.Pertenece(clave3))
	arbol.Guardar(clave1, valor1)
	arbol.Guardar(clave2, valor2)
	arbol.Guardar(clave3, valor3)
	require.True(t, arbol.Pertenece(clave1))
	require.True(t, arbol.Pertenece(clave2))
	require.True(t, arbol.Pertenece(clave3))
	require.Equal(t, valor1, arbol.Obtener(clave1))
	require.Equal(t, valor2, arbol.Obtener(clave2))
	require.Equal(t, valor3, arbol.Obtener(clave3))
	require.Equal(t, 3, arbol.Cantidad())
	arbol.Guardar(clave1, "nuevo_valor")
	require.Equal(t, 3, arbol.Cantidad())
	require.Equal(t, "nuevo_valor", arbol.Obtener(clave1))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Libro"
	clave2 := "Cartuchera"
	dic := TDAAbb.CrearABB[string, string](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	dic.Guardar(clave, "lapiz")
	dic.Guardar(clave2, "mochila")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "lapiz", dic.Obtener(clave))
	require.EqualValues(t, "mochila", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "borrador")
	dic.Guardar(clave2, "cartera")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "borrador", dic.Obtener(clave))
	require.EqualValues(t, "cartera", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	cmpFunc := func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	dic := TDAAbb.CrearABB[int, int](cmpFunc)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestReutilizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un ABB, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDAAbb.CrearABB[string, string](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Mickey"
	clave2 := "Donald"
	clave3 := "Pluto"
	valor1 := "Minnie"
	valor2 := "Daisy"
	valor3 := "Goofy"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDAAbb.CrearABB[string, string](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDAAbb.CrearABB[int, string](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	clave := 10
	valor := "Buenas"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

//fail

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y debería funcionar sin problemas")
	dic := TDAAbb.CrearABB[string, string](func(a, b string) int {
		return strings.Compare(a, b)
	})
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDAAbb.CrearABB[string, *int](func(a, b string) int {
		return strings.Compare(a, b)
	})
	clave := "Niño"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error común es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	abb := TDAAbb.CrearABB[int, int](funcionComparacion)

	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		if !abb.Pertenece(i) {
			t.Errorf("Error: El valor %d no se insertó correctamente", i)
		}
		abb.Borrar(i)
		if abb.Pertenece(i) {
			t.Errorf("Error: El valor %d no se eliminó correctamente", i)
		}
	}
	for i := 0; i < 1000; i++ {
		numeroAleatorio := rand.Intn(10000) + 1
		abb.Guardar(numeroAleatorio, numeroAleatorio)
		if !abb.Pertenece(numeroAleatorio) {
			t.Errorf("Error: El valor %d no se insertó correctamente", numeroAleatorio)
		}
		abb.Borrar(numeroAleatorio)
		if abb.Pertenece(numeroAleatorio) {
			t.Errorf("Error: El valor %d no se eliminó correctamente", numeroAleatorio)
		}
	}
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")

	clave1 := "Cafe"
	clave2 := "Mate"
	clave3 := "Te"
	claves := []string{clave1, clave2, clave3}

	arbol := TDAAbb.CrearABB[string, *int](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	arbol.Guardar(claves[0], nil)
	arbol.Guardar(claves[1], nil)
	arbol.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0

	arbol.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		cantidad++
		return true
	})

	require.Equal(t, 3, cantidad, "La cantidad de claves recorridas no es 3")

	for _, clave := range claves {
		contador := 0
		for _, c := range cs {
			if c == clave {
				contador++
			}
		}
		require.Equal(t, 1, contador, fmt.Sprintf("La clave %s no se recorrió exactamente una vez", clave))
	}
}

//fail

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDAAbb.CrearABB[string, int](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

// fail

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDAAbb.CrearABB[int, int](funcionComparacion)
	dic1 := TDAAbb.CrearABB[int, int](funcionComparacion)
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
	for i := 0; i < 10000; i++ {
		numeroAleatorio := rand.Intn(10000) + 1
		dic1.Guardar(numeroAleatorio, numeroAleatorio)
	}

	seguirEjecutando1 := true
	siguioEjecutandoCuandoNoDebia1 := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando1 {
			siguioEjecutandoCuandoNoDebia1 = true
			return false
		}
		if c%2 == 0 {
			seguirEjecutando1 = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando1, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia1,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")

}

func TestSumaElementos(t *testing.T) {
	fmt.Println("Prueba de suma de elementos en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for i := 1; i <= 5; i++ {
		dic.Guardar(i, i)
	}

	var suma int
	dic.Iterar(func(c int, v int) bool {
		suma += v
		return true
	})

	require.Equal(t, 15, suma, "La suma de elementos no es la esperada")
}

func TestIteracionSinCondicionCorte(t *testing.T) {
	t.Log("Prueba de iteración sin condición de corte en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for i := 1; i <= 5; i++ {
		dic.Guardar(i, i)
	}

	t.Log("Iteración sin condición de corte:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestIteracionConCondicionCorte(t *testing.T) {
	t.Log("Prueba de iteración con condición de corte en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for i := 1; i <= 10; i++ {
		dic.Guardar(i, i)
	}

	t.Log("Iteración con condición de corte:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return c != 5
	})
}

func TestIteracionConRangos(t *testing.T) {
	t.Log("Prueba de iteración con rangos en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for i := 1; i <= 10; i++ {
		dic.Guardar(i, i)
	}

	t.Log("Iteración con rangos [3, 7]:")
	desde := 3
	hasta := 7
	dic.IterarRango(&desde, &hasta, func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestIteracionSinRangos(t *testing.T) {
	t.Log("Prueba de iteración sin rangos en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for i := 1; i <= 5; i++ {
		dic.Guardar(i, i)
	}

	t.Log("Iteración sin rangos:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestIteracionSinCondicionCorteDesordenado(t *testing.T) {
	t.Log("Prueba de iteración sin condición de corte en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)

	t.Log("Iteración sin condición de corte:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestIteracionConCondicionCorteDesordenado(t *testing.T) {
	t.Log("Prueba de iteración con condición de corte en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)

	t.Log("Iteración con condición de corte:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return c != 3
	})
}

func TestIteracionConRangosDesordenado(t *testing.T) {
	t.Log("Prueba de iteración con rangos en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)

	t.Log("Iteración con rangos [3, 7]:")
	desde := 3
	hasta := 7
	dic.IterarRango(&desde, &hasta, func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestIteracionSinRangosDesordenado(t *testing.T) {
	t.Log("Prueba de iteración sin rangos en el ABB")
	dic := TDAAbb.CrearABB[int, int](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)

	t.Log("Iteración sin rangos:")
	dic.Iterar(func(c int, v int) bool {
		require.True(t, dic.Pertenece(c), "La clave debería pertenecer al diccionario")
		require.Equal(t, v, dic.Obtener(c), "El valor obtenido no coincide con el esperado")
		return true
	})
}

func TestTrabajanBienConFuncionDeComparacion(t *testing.T) {

	funcionComparacion := func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	arbol := TDAAbb.CrearABB[string, int](funcionComparacion)

	arbol.Guardar("clave1", 1)
	arbol.Guardar("clave2", 2)
	arbol.Guardar("clave3", 3)

	if !arbol.Pertenece("clave1") {
		t.Error("La clave 'clave1' no pertenece al árbol")
	}
	if !arbol.Pertenece("clave2") {
		t.Error("La clave 'clave2' no pertenece al árbol")
	}
	if !arbol.Pertenece("clave3") {
		t.Error("La clave 'clave3' no pertenece al árbol")
	}

	if arbol.Obtener("clave1") != 1 {
		t.Error("El valor asociado a la clave 'clave1' es incorrecto")
	}
	if arbol.Obtener("clave2") != 2 {
		t.Error("El valor asociado a la clave 'clave2' es incorrecto")
	}
	if arbol.Obtener("clave3") != 3 {
		t.Error("El valor asociado a la clave 'clave3' es incorrecto")
	}
}

func TestDiccionario(t *testing.T) {
	funcionComparacion := func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	arbol := TDAAbb.CrearABB[string, int](funcionComparacion)

	arbol.Guardar("123456789", 123)
	arbol.Guardar("987654321", 456)
	arbol.Guardar("555123456", 789)

	require.True(t, arbol.Pertenece("123456789"), "La clave '123456789' no pertenece al árbol")
	require.True(t, arbol.Pertenece("987654321"), "La clave '987654321' no pertenece al árbol")
	require.True(t, arbol.Pertenece("555123456"), "La clave '555123456' no pertenece al árbol")

	arbol.Borrar("987654321")

	require.True(t, arbol.Pertenece("123456789"), "La clave '123456789' no pertenece al árbol después de borrar '987654321'")
	require.False(t, arbol.Pertenece("987654321"), "La clave '987654321' sigue perteneciendo al árbol después de borrarla")
	require.True(t, arbol.Pertenece("555123456"), "La clave '555123456' no pertenece al árbol después de borrar '987654321'")

	require.Equal(t, arbol.Obtener("123456789"), 123, "El valor asociado a la clave '123456789' es incorrecto después de borrar '987654321'")
	require.Equal(t, arbol.Obtener("555123456"), 789, "El valor asociado a la clave '555123456' es incorrecto después de borrar '987654321'")
}

type basico struct {
	a string
	b int
}

type avanzado struct {
	w int
	x basico
	y basico
	z string
}

func compararAvanzado(a, b avanzado) int {

	if a.z < b.z {
		return -1
	} else if a.z > b.z {
		return 1
	}

	if a.w < b.w {
		return -1
	} else if a.w > b.w {
		return 1
	}

	if cmp := compararBasico(a.x, b.x); cmp != 0 {
		return cmp
	}
	return compararBasico(a.y, b.y)
}

func compararBasico(a, b basico) int {

	if a.a < b.a {
		return -1
	} else if a.a > b.a {
		return 1
	}

	if a.b < b.b {
		return -1
	} else if a.b > b.b {
		return 1
	}

	return 0
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que también funcione con estructuras más complejas")

	arbol := TDAAbb.CrearABB[avanzado, int](compararAvanzado)

	a1 := avanzado{w: 10, z: "Super Mario Bros", x: basico{a: "Super Mario", b: 1985}, y: basico{a: "Bros", b: 1}}
	a2 := avanzado{w: 10, z: "Super Mario Odyssey", x: basico{a: "Super Mario", b: 2017}, y: basico{a: "Odyssey", b: 2}}
	a3 := avanzado{w: 10, z: "Super Mario Kart", x: basico{a: "Super Mario", b: 1992}, y: basico{a: "Kart", b: 3}}

	arbol.Guardar(a1, 0)
	arbol.Guardar(a2, 1)
	arbol.Guardar(a3, 2)

	require.True(t, arbol.Pertenece(a1))
	require.True(t, arbol.Pertenece(a2))
	require.True(t, arbol.Pertenece(a3))
	require.EqualValues(t, 0, arbol.Obtener(a1))
	require.EqualValues(t, 1, arbol.Obtener(a2))
	require.EqualValues(t, 2, arbol.Obtener(a3))

	arbol.Guardar(a1, 5)
	require.EqualValues(t, 5, arbol.Obtener(a1))
	require.EqualValues(t, 2, arbol.Obtener(a3))

	require.EqualValues(t, 5, arbol.Borrar(a1))
	require.False(t, arbol.Pertenece(a1))
	require.EqualValues(t, 2, arbol.Obtener(a3))
}

func funcionComparacion(a, b int) int {
	return a - b
}

func TestABB(t *testing.T) {
	// Crear un árbol binario de búsqueda utilizando la función de comparación personalizada
	arbol := TDAAbb.CrearABB[int, int](funcionComparacion)

	// Insertar 6 enteros en el árbol en desorden
	arbol.Guardar(10, 10)
	arbol.Guardar(5, 5)
	arbol.Guardar(15, 15)
	arbol.Guardar(2, 2)
	arbol.Guardar(7, 7)
	arbol.Guardar(12, 12)

	// Verificar que los elementos estén en el árbol
	tests := []struct {
		clave    int
		esperado bool
	}{
		{2, true},
		{5, true},
		{7, true},
		{10, true},
		{12, true},
		{15, true},
	}

	for _, tt := range tests {
		if arbol.Pertenece(tt.clave) != tt.esperado {
			t.Errorf("Pertenece(%d) = %v; want %v", tt.clave, !tt.esperado, tt.esperado)
		}
	}

	// Borrar un elemento del medio (10)
	arbol.Borrar(10)

	// Verificar que los elementos restantes sigan estando en el árbol
	testsAfterDelete := []struct {
		clave    int
		esperado bool
	}{
		{2, true},
		{5, true},
		{7, true},
		{10, false},
		{12, true},
		{15, true},
	}

	for _, tt := range testsAfterDelete {
		if arbol.Pertenece(tt.clave) != tt.esperado {
			t.Errorf("Pertenece(%d) después de borrar 10 = %v; want %v", tt.clave, !tt.esperado, tt.esperado)
		}
	}

	// Verificar que los valores asociados a las claves sean correctos
	valuesTests := []struct {
		clave    int
		esperado int
	}{
		{5, 5},
		{12, 12},
		{2, 2},
		{15, 15},
	}

	for _, tt := range valuesTests {
		if valor := arbol.Obtener(tt.clave); valor != tt.esperado {
			t.Errorf("Obtener(%d) = %d; want %d", tt.clave, valor, tt.esperado)
		}
	}
}
