package main

import (
	"github.com/mhasmat/godesde0/teclado"
	// "runtime"
)

func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoATexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	/*
	os := runtime.GOOS

	if os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es: ", os)
	} else {
		fmt.Println("Esto es Windows: ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.ConversorNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)
	*/

	teclado.IngresoNumeros()
}