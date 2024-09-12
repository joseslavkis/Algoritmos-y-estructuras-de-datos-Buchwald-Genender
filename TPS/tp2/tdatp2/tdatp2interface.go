package tdatp2

import (
	LISTA "tdas/lista"
)

type InformacionIPs interface {
	AgregarArchivo(nombreArchivo string) ([]string, error)

	VerVisitantes(ip1 string, ip2 string) (LISTA.Lista[string], error)

	VerMasVisitados(cantidad int) (LISTA.Lista[Recursos], error)
}
