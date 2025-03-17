package main

import (
	"fmt"
)

func main() {
	// Solicitar al usuario que ingrese sus nombres completos
	fmt.Print("Escriba sus nombres completos: ")
	var a string
	fmt.Scanln(&a)

	// Solicitar al usuario que ingrese sus apellidos completos
	fmt.Print("Escriba sus apellidos completos: ")
	var b string
	fmt.Scanln(&b)

	// Solicitar al usuario que ingrese su profesión
	fmt.Print("Escriba su profesión: ")
	var c string
	fmt.Scanln(&c)

	// Solicitar al usuario que ingrese su año de nacimiento
	fmt.Print("Escriba su año de nacimiento: ")
	var d int
	fmt.Scanln(&d)

	// Calcular la edad en el año 2025
	e := 2025 - d

	// Mostrar el resultado
	fmt.Printf("El (La) %s %s %s tiene %d años\n", c, a, b, e)
}
