package main

import (
	"bufio"
	"fmt"
	"os"

	"dc/operacion"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		resultado, err := operacion.RedirigirAccion(linea)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(resultado)
		}
	}
}
