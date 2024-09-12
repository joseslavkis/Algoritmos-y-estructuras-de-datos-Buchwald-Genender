package tdatp2

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	DoS "tdados"
	HEAP "tdas/cola_prioridad"
	DICCIONARIO "tdas/diccionario"
	LISTA "tdas/lista"
	"time"
)

type Recursos struct {
	Nombre  string
	Visitas int
}

const LAYOUT = "2006-01-02T15:04:05-07:00"
const _POSICION_IP = 0
const _POSICION_TIEMPO = 1
const _POSICION_URL = 3

type informacionIPs struct {
	visitantes   DICCIONARIO.DiccionarioOrdenado[string, string]
	cantVisitas  DICCIONARIO.Diccionario[string, int]
	masVisitados HEAP.ColaPrioridad[Recursos]
}

func CompararIPs(ip1, ip2 string) int {
	ip1Parts := strings.Split(ip1, ".")
	ip2Parts := strings.Split(ip2, ".")
	for i := 0; i < 4; i++ {
		part1, _ := strconv.Atoi(ip1Parts[i])
		part2, _ := strconv.Atoi(ip2Parts[i])
		if part1 < part2 {
			return -1
		} else if part1 > part2 {
			return 1
		}
	}
	return 0
}

func (informacionIPs *informacionIPs) cargarInformacion(scanner *bufio.Scanner, estructuraDoS DoS.EstructuraDoS) error {
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Fields(linea)
		if !informacionIPs.cantVisitas.Pertenece(campos[_POSICION_URL]) {
			informacionIPs.cantVisitas.Guardar(campos[_POSICION_URL], 1)
		} else {
			dato := informacionIPs.cantVisitas.Obtener(campos[_POSICION_URL])
			informacionIPs.cantVisitas.Guardar(campos[_POSICION_URL], dato+1)
		}
		informacionIPs.visitantes.Guardar(campos[_POSICION_IP], campos[_POSICION_IP])
		tiempo, err := time.Parse(LAYOUT, campos[_POSICION_TIEMPO])
		if err != nil {
			return err
		}
		estructuraDoS.AñadirVisita(campos[_POSICION_IP], tiempo)
	}
	iter := informacionIPs.cantVisitas.Iterador()
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		informacionIPs.masVisitados.Encolar(Recursos{Nombre: clave, Visitas: dato})
		iter.Siguiente()
	}
	return nil
}

func CrearInformacionIPs() InformacionIPs {
	return &informacionIPs{
		visitantes:   DICCIONARIO.CrearABB[string, string](CompararIPs),
		cantVisitas:  DICCIONARIO.CrearHash[string, int](),
		masVisitados: HEAP.CrearHeap[Recursos](func(a, b Recursos) int { return a.Visitas - b.Visitas }),
	}
}

func (estructura *informacionIPs) AgregarArchivo(nombreArchivo string) ([]string, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	estructuraDoS := DoS.CrearEstructuraDoS()
	err = estructura.cargarInformacion(scanner, estructuraDoS)
	if err != nil {
		return nil, err
	}
	_, err = archivo.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return estructuraDoS.CrearArrAtacantes(), nil
}

func (estructura *informacionIPs) VerVisitantes(ip1 string, ip2 string) (LISTA.Lista[string], error) {
	lista := LISTA.CrearListaEnlazada[string]()
	iter := estructura.visitantes.IteradorRango(&ip1, &ip2)
	for iter.HaySiguiente() {
		ip, _ := iter.VerActual()
		lista.InsertarUltimo(ip)
		iter.Siguiente()
	}
	if lista.EstaVacia() {
		return nil, errors.New("error")
	}

	return lista, nil
}

func (estructura *informacionIPs) VerMasVisitados(cantidad int) (LISTA.Lista[Recursos], error) {
	lista := LISTA.CrearListaEnlazada[Recursos]()
	iterLista := lista.Iterador()
	heap := estructura.masVisitados
	for i := 0; i < cantidad && heap.Cantidad() > 0; i++ {
		lista.InsertarUltimo(heap.Desencolar())
	}
	for iterLista.HaySiguiente() {
		heap.Encolar(iterLista.VerActual())
		iterLista.Siguiente()
	}
	if lista.EstaVacia() {
		return nil, errors.New("error")
	}
	return lista, nil
}
