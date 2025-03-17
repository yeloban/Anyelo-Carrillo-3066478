package main

import (
	"fmt"
)

func main() {
	// Solicitar al usuario que ingrese el valor del voltaje
	fmt.Print("Ingrese el valor del voltaje del circuito: ")
	var Voltaje int
	fmt.Scanln(&Voltaje)

	// Solicitar al usuario que ingrese el valor de la resistencia
	fmt.Print("Ingrese el valor de la resistencia del circuito: ")
	var Resistencia int
	fmt.Scanln(&Resistencia)

	// Calcular la intensidad (amperaje)
	Intensidad := float64(Voltaje) / float64(Resistencia)

	// Mostrar el resultado
	fmt.Printf("Al conectar un resistor de R %d ohm a una fuente de V %d voltios, circulará una corriente de %.2f amperios\n", Resistencia, Voltaje, Intensidad)
}
