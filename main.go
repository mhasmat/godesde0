package main

import (
	"fmt"

	"github.com/mhasmat/godesde0/variables"
)

func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}