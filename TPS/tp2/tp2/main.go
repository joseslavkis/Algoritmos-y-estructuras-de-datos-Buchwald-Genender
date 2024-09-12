package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TP2 "tdatp2"
)

const (
	_POSICION_COMANDO        = 0
	_NOMBRE_ARCHIVO_POSICION = 1
	_IP1                     = 1
	_IP2                     = 2
	_POSICION_MAS_VISITADOS  = 1
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	infoIPs := TP2.CrearInformacionIPs()
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Fields(linea)
		if len(campos) > 0 {
			comando := campos[_POSICION_COMANDO]
			switch comando {
			case "agregar_archivo":
				if len(campos) != 2 {
					fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
					return
				}
				nombre_archivo := campos[_NOMBRE_ARCHIVO_POSICION]
				atacantes, err := infoIPs.AgregarArchivo(nombre_archivo)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
					return
				}
				for i := 0; i < len(atacantes); i++ {
					fmt.Printf("DoS: %s\n", atacantes[i])
				}
				fmt.Println("OK")
			case "ver_visitantes":
				if len(campos) != 3 {
					fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
					return
				}
				ip1 := campos[_IP1]
				ip2 := campos[_IP2]
				visitantes, err := infoIPs.VerVisitantes(ip1, ip2)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error en comando %s: %v\n", comando, err)
					return
				}
				fmt.Println("Visitantes:")
				for !visitantes.EstaVacia() {
					fmt.Printf("\t%s\n", visitantes.BorrarPrimero())
				}
				fmt.Println("OK")
			case "ver_mas_visitados":
				if len(campos) != 2 {
					fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
					return
				}
				cant_mas_visitados := campos[_POSICION_MAS_VISITADOS]
				num, err := strconv.Atoi(cant_mas_visitados)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error en comando %s: %v\n", comando, err)
					return
				}
				if num <= 0 {
					fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
					return
				}
				masVisitados, _err := infoIPs.VerMasVisitados(num)
				if _err != nil {
					fmt.Fprintf(os.Stderr, "Error en comando %s: %v\n", comando, _err)
					return
				}
				fmt.Println("Sitios más visitados:")
				for !masVisitados.EstaVacia() {
					elem := masVisitados.BorrarPrimero()
					fmt.Printf("\t%s - %d\n", elem.Nombre, elem.Visitas)
				}
				fmt.Println("OK")
			default:
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
				return
			}
		}
	}
}
