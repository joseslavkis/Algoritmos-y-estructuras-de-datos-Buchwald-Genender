package tdados

import (
	"strconv"
	"strings"
	DICCIONARIO "tdas/diccionario"
	LISTA "tdas/lista"
	"time"
)

const _TIEMPO_DOS = 2
const _RANGO_IP = 256
const _CANT_CAMPOS_IP = 4
const _CANTIDAD_ATACANTES_MINIMA = 5
const LAYOUT = "2006-01-02T15:04:05-07:00"

type estructuraDoS struct {
	diccionarioIPs       DICCIONARIO.Diccionario[string, LISTA.Lista[time.Time]]
	diccionarioAtacantes DICCIONARIO.Diccionario[string, string]
}

func CrearEstructuraDoS() *estructuraDoS {
	_IPs := DICCIONARIO.CrearHash[string, LISTA.Lista[time.Time]]()
	_atacantes := DICCIONARIO.CrearHash[string, string]()
	return &estructuraDoS{diccionarioIPs: _IPs, diccionarioAtacantes: _atacantes}
}

func (estructura *estructuraDoS) AñadirVisita(ip string, horario time.Time) {
	if !estructura.diccionarioIPs.Pertenece(ip) {
		listaAux := LISTA.CrearListaEnlazada[time.Time]()
		listaAux.InsertarUltimo(horario)
		estructura.diccionarioIPs.Guardar(ip, listaAux)
	} else {
		tiemposDoS := estructura.diccionarioIPs.Obtener(ip)
		if tiemposDoS.Largo() < _CANTIDAD_ATACANTES_MINIMA-1 {
			tiemposDoS.InsertarUltimo(horario)
		} else {
			diferencia := horario.Sub(tiemposDoS.VerPrimero()).Seconds()
			tiemposDoS.BorrarPrimero()
			tiemposDoS.InsertarUltimo(horario)
			if diferencia < _TIEMPO_DOS {
				estructura.diccionarioAtacantes.Guardar(ip, ip)
			}
		}
	}
}

func (estructura *estructuraDoS) CrearArrAtacantes() []string {
	atacantes := make([][]string, estructura.diccionarioAtacantes.Cantidad())
	for i, iter := 0, estructura.diccionarioAtacantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, _ := iter.VerActual()
		atacantes[i] = strings.Split(ip, ".")
		i++
	}
	for i := _CANT_CAMPOS_IP - 1; i >= 0; i-- {
		countingSort(atacantes, _RANGO_IP, i)
	}
	result := make([]string, len(atacantes))
	for i, elem := range atacantes {
		result[i] = strings.Join(elem, ".")
	}
	return result
}

func countingSort(atacantes [][]string, rango, digito int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([][]string, len(atacantes))
	for _, elem := range atacantes {
		valor, _ := strconv.Atoi(elem[digito])
		frecuencias[valor]++
	}
	for i := 1; i < rango; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	for _, elem := range atacantes {
		valor, _ := strconv.Atoi(elem[digito])
		resultado[sumasAcumuladas[valor]] = elem
		sumasAcumuladas[valor]++
	}
	copy(atacantes, resultado)
}
