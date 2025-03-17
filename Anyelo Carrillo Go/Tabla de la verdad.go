package main

import "fmt"

func main() {
	// Operaciones con booleanos
	a := true
	b := false
	fmt.Println(a && b) // Imprime "false" (equivalente a "and" en Python)

	// Operaciones con enteros y comparaciones
	aInt := 2
	bInt := 3
	c := 4
	d := 10

	// Comparación 1: a == b and c < d
	fmt.Println(aInt == bInt && c < d) // Imprime "false"

	// Comparación 2: not a == b and c > d
	fmt.Println(!(aInt == bInt) && c > d) // Imprime "false"
}
