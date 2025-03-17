package main

import (
	"fmt"
	"math"
)

func main() {
	// Solicitar al usuario que ingrese el primer valor
	fmt.Println("Ingrese el primer valor:")
	var A int
	fmt.Scanln(&A) // Leer el valor ingresado y convertirlo a entero

	// Inicializar la variable B (aunque no se usa en el código)
	var B int = 0

	// Solicitar al usuario que ingrese el segundo valor
	fmt.Println("Ingrese el segundo valor:")
	var C int
	fmt.Scanln(&C) // Leer el valor ingresado y convertirlo a entero

	// Calcular la potencia de A elevado a C
	valor := math.Pow(float64(A), float64(C)) // Usar math.Pow para calcular la potencia

	// Mostrar el resultado
	fmt.Printf("La potencia de %d sobre %d es: %.0f\n", A, C, valor)
}
