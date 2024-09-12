package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func abrirArchivo(nombreArchivo string) []int {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		panic(fmt.Sprintf("No se pudo abrir %s debido a %s", nombreArchivo, err))
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	slice := []int{}
	for scanner.Scan() {
		linea := scanner.Text()
		numero, err := strconv.Atoi(linea)
		if err != nil {
			panic(fmt.Sprintf("Error al convertir en entero un número (archivo1): %s\n", err))
		}
		slice = append(slice, numero)
	}
	return slice
}

func imprimirMayorArreglo(slice1 []int, slice2 []int) {
	resultadoComparacion := ejercicios.Comparar(slice1, slice2)
	var mayorSlice []int
	if resultadoComparacion == -1 {
		mayorSlice = slice2
	} else {
		mayorSlice = slice1
	}

	ejercicios.Seleccion(mayorSlice)
	for _, num := range mayorSlice {
		fmt.Println(num)
	}
}

func main() {
	slice1 := abrirArchivo("archivo1.in")
	slice2 := abrirArchivo("archivo2.in")
	imprimirMayorArreglo(slice1, slice2)
}
