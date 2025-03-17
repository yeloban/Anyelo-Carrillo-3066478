package main

import (
	"fmt"
	"math"
)

func main() {
	// Definir variables
	a := 10
	b := 4

	// Multiplicación
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: *")
	fmt.Println("La segunda variable es:", b)
	c := a * b
	fmt.Println("El resultado es:", c)
	fmt.Printf("El tipo de c es: %T\n", c)

	// División
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: /")
	fmt.Println("La segunda variable es:", b)
	c = a / b
	fmt.Println("El resultado es:", c)
	fmt.Printf("El tipo de c es: %T\n", c)

	// Suma
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: +")
	fmt.Println("La segunda variable es:", b)
	c = a + b
	fmt.Println("El resultado es:", c)
	fmt.Printf("El tipo de c es: %T\n", c)

	// Resta
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: -")
	fmt.Println("La segunda variable es:", b)
	c = a - b
	fmt.Println("El resultado es:", c)
	fmt.Printf("El tipo de c es: %T\n", c)

	// Potencia (Go no tiene un operador de potencia, se usa math.Pow)
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: **")
	fmt.Println("La segunda variable es:", b)
	cFloat := math.Pow(float64(a), float64(b))
	fmt.Println("El resultado es:", cFloat)
	fmt.Printf("El tipo de c es: %T\n", cFloat)

	// Módulo
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: %")
	fmt.Println("La segunda variable es:", b)
	c = a % b
	fmt.Println("El resultado es:", c)
	fmt.Printf("El tipo de c es: %T\n", c)
}
