package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) > 0 {
		max := vector[0]
		posicionMaximo := 0
		for i := 0; i < len(vector); i++ {
			if vector[i] > max {
				max = vector[i]
				posicionMaximo = i
			}
		}
		return posicionMaximo
	}

	return -1
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	largoMenor := len(vector2)
	if len(vector1) < largoMenor {
		largoMenor = len(vector1)
	}
	for i := 0; i < largoMenor; i++ {
		if vector2[i] > vector1[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if len(vector1) > len(vector2) {
		return 1
	} else if len(vector1) < len(vector2) {
		return -1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector)-1; i++ {
		max := Maximo(vector[0 : len(vector)-i])
		Swap(&vector[max], &vector[len(vector)-1-i])
	}
}

func sumaRecursiva(vector []int, indice int) int {
	if indice == len(vector) {
		return 0
	}
	return vector[indice] + sumaRecursiva(vector, indice+1)
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}

	return sumaRecursiva(vector, 0)
}

func capicuaRecursiva(cadena string, inicio int, fin int) bool {
	if inicio >= fin {
		return true
	}
	if cadena[inicio] != cadena[fin] {
		return false
	}
	return capicuaRecursiva(cadena, inicio+1, fin-1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return capicuaRecursiva(cadena, 0, len(cadena)-1)
}
