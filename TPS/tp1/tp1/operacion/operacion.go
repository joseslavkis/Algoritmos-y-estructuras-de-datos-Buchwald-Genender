package operacion

import (
	"errors"
	"math"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

type Operacion struct {
	Terminos int
	Operar   func(params []int) (int, error)
	Simbolo  string
}

var operaciones = []Operacion{
	{2, func(params []int) (int, error) {
		return params[0] + params[1], nil
	}, "+"},
	{2, func(params []int) (int, error) {
		return params[0] - params[1], nil
	}, "-"},
	{2, func(params []int) (int, error) {
		return params[0] * params[1], nil
	}, "*"},
	{2, func(params []int) (int, error) {
		if params[1] == 0 {
			return 0, errors.New("error: División por cero")
		}
		return params[0] / params[1], nil
	}, "/"},
	{1, func(params []int) (int, error) {
		if params[0] < 0 {
			return 0, errors.New("error: No existe solución real para la raíz cuadrada de un número negativo")
		}
		return int(math.Sqrt(float64(params[0]))), nil
	}, "sqrt"},
	{2, func(params []int) (int, error) {
		if params[1] < 0 {
			return 0, errors.New("error: El exponente no puede ser negativo")
		}
		return int(math.Pow(float64(params[0]), float64(params[1]))), nil
	}, "^"},
	{2, func(params []int) (int, error) {
		if params[1] < 2 {
			return 0, errors.New("error: La base del logaritmo debe ser mayor o igual a 2")
		}
		return int(math.Log(float64(params[0])) / math.Log(float64(params[1]))), nil
	}, "log"},
	{3, func(params []int) (int, error) {
		if params[0] != 0 {
			return params[1], nil
		}
		return params[2], nil
	}, "?"},
}

func RedirigirAccion(entrada string) (int, error) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for _, valor := range strings.Fields(entrada) {
		numero, err := strconv.Atoi(valor)
		if err == nil {
			pila.Apilar(numero)
		} else {
			var operacion Operacion
			encontrada := false
			for _, operador := range operaciones {
				if operador.Simbolo == valor {
					operacion = operador
					encontrada = true
					break
				}
			}
			if !encontrada {
				return 0, errors.New("error: valor no válido")
			}
			parametros := make([]int, operacion.Terminos)
			for i := len(parametros) - 1; i >= 0; i-- {
				if pila.EstaVacia() {
					return 0, errors.New("error: no hay suficientes operandos para la operación")
				}
				parametros[i] = pila.Desapilar()
			}
			resultado, err := operacion.Operar(parametros)
			if err != nil {
				return 0, err
			}
			pila.Apilar(resultado)
		}
	}
	if pila.EstaVacia() {
		return 0, errors.New("error: la pila está vacía después de evaluar todas las operaciones")
	}
	elemento := pila.Desapilar()
	if !pila.EstaVacia() {
		return 0, errors.New("error: al finalizar la evaluación queda más de un valor restante en la pila")
	}
	return elemento, nil
}
