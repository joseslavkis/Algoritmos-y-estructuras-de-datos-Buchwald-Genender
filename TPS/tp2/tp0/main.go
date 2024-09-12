package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func main() {
	archivo1, error1 := os.Open("archivo1.in")
	archivo2, error2 := os.Open("archivo2.in")
	if error1 != nil {
		fmt.Printf("No se pudo abrir el archivo1: %s", error1)
		return
	}
	if error2 != nil {
		fmt.Printf("No se pudo abrir el archivo2: %s", error2)
		return // Agrega este retorno para salir de la función en caso de error al abrir archivo2
	}
	defer archivo1.Close()
	defer archivo2.Close()

	scanner1 := bufio.NewScanner(archivo1)
	scanner2 := bufio.NewScanner(archivo2)

	slice1 := []int{}
	slice2 := []int{}

	for scanner1.Scan() {
		linea := scanner1.Text()
		numero, error := strconv.Atoi(linea)
		if error != nil {
			fmt.Printf("Error al convertir en entero un numero(archivo1): %s", error)
		}
		slice1 = append(slice1, numero)
	}
	for scanner2.Scan() {
		linea := scanner2.Text()
		numero, error := strconv.Atoi(linea)
		if error != nil {
			fmt.Printf("Error al convertir en entero un numero(archivo2): %s", error)
		}
		slice2 = append(slice2, numero)
	}
	resultadoComparacion := ejercicios.Comparar(slice1, slice2)
	if resultadoComparacion == -1 {
		ejercicios.Seleccion(slice2)
		for i := 0; i < len(slice2); i++ {
			fmt.Println(slice2[i])
		}
	} else if resultadoComparacion == 1 {
		ejercicios.Seleccion(slice1)
		for i := 0; i < len(slice1); i++ {
			fmt.Println(slice1[i])
		}
	} else {
		ejercicios.Seleccion(slice2)
		for i := 0; i < len(slice2); i++ {
			fmt.Println(slice2[i])
		}
	}
}
