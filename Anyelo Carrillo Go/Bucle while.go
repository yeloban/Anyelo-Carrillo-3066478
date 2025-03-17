package main

import (
	"fmt"
)

func main() {
	// Primer bucle for (equivalente al while en Python)
	contador := 0
	for contador < 30 {
		contador++
		fmt.Printf("Iteración %d\n", contador)
	}

	// Combinado con if-else y manejo de errores
	for {
		var a int
		fmt.Print("Introduzca un valor mayor a 10: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Por favor, ingrese un número válido")
			continue // Vuelve al inicio del bucle si hay un error
		}

		if a > 10 {
			fmt.Println("Es correcto")
			break // Sale del bucle cuando la condición se cumple
		} else {
			fmt.Println("Es incorrecto, pruebe de nuevo")
		}
	}
}
