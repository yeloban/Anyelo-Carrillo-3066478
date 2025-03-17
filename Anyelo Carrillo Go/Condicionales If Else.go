package main

import "fmt"

func main() {
	// Asignar valor a la variable 'a'
	a := 2

	// Comprobar si 'a' es igual a 2
	if a == 2 {
		fmt.Println("a vale 2") // Se ejecuta esta línea porque 'a' es igual a 2
	} else {
		fmt.Println("a es diferente de 2")
	}

	// ------------------------------------------------

	// Concatenar condiciones con '&&' (and) y '!' (not)

	// Definir variables con información personal
	Nombre := "Esteban"
	Edad := 28
	Pais := "Colombia"

	// Primera condición: Verifica si todas las condiciones son verdaderas
	if Nombre == "Esteban" && Edad == 28 && Pais == "Colombia" {
		fmt.Printf("Su nombre es %s, tiene %d años y es de %s\n", Nombre, Edad, Pais)
	}

	// Segunda condición: Verifica si el nombre es 'Esteban', pero la edad NO es 28
	if Nombre == "Esteban" && Edad != 28 {
		fmt.Println("Su nombre es Esteban y no tiene 28 años")
	}

	// Tercera condición: Verifica si el nombre NO es 'Esteban', pero la edad es 28
	if Nombre != "Esteban" && Edad == 28 {
		fmt.Println("Su nombre no es Esteban y tiene 28 años")
	}

	// Si ninguna condición anterior se cumple, entra en el bloque 'else'
	if !(Nombre == "Esteban" && Edad == 28) {
		fmt.Println("No se llama Esteban ni tiene 28 años")
	}
}
