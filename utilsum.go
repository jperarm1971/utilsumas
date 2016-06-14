//utilcadenas es un paquete que contiene funciones utiles para trabajar
//con cadenas de texto
package utilsum

import "fmt"
import "strconv"

func ValidaOperando(op string) int {
	n1, err := strconv.Atoi(op)
	if err != nil {
		fmt.Println("El primer parámetro no es un nº valido")
		return 0
	}
	return n1
}

func Operacionsuma(op1 int, op2 int) int {
	return op1 + op2
}

func ValidaNumParams(nparams int) int {
	if nparams < 3 {
		fmt.Println("Se necesitan dos nºs como parámetros", nparams)
		return 0
	}
	return 1
}
