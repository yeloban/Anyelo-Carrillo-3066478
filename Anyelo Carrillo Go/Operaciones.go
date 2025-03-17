package main

import "fmt"

func main() {
	// Solicitar al usuario que ingrese los valores
	var A, B int
	fmt.Print("Ingrese valor: ")
	fmt.Scanln(&A)
	fmt.Print("Ingrese valor: ")
	fmt.Scanln(&B)

	// Realizar operaciones aritméticas
	suma := A + B
	fmt.Printf("La suma de los números es: %d\n", suma)

	res := A - B
	fmt.Printf("La resta de los números es: %d\n", res)

	multi := A * B
	fmt.Printf("La multiplicación de los números es: %d\n", multi)

	div := float64(A) / float64(B)
	fmt.Printf("La división de los números es: %.2f\n", div)

	// Comparaciones
	igual := A == B
	fmt.Printf("¿El número %d es igual a %d? %t\n", A, B, igual)

	mayor := A > B
	if mayor {
		fmt.Printf("El número mayor es: %d\n", A)
		fmt.Printf("El número menor es: %d\n", B)
	} else {
		fmt.Printf("El número mayor es: %d\n", B)
		fmt.Printf("El número menor es: %d\n", A)
	}
}
