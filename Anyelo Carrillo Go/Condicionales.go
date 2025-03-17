package main

import (
	"fmt"
)

func main() {
	// Solicitar al usuario que seleccione la figura
	fmt.Println("Seleccione la figura a calcular su área:")
	fmt.Println("1 para rombo")
	fmt.Println("2 para rectángulo")
	fmt.Println("3 para cuadrado")
	fmt.Println("4 para trapecio")
	var Figura string
	fmt.Scanln(&Figura)

	// Definir constantes
	const Pi = 3.1416
	const Const = 2

	// Usar un switch para manejar las opciones
	switch Figura {
	case "1":
		var Dmayor, Dmenor int
		fmt.Print("Ingresa el valor de la diagonal mayor: ")
		fmt.Scanln(&Dmayor)
		fmt.Print("Ingresa el valor de la diagonal menor: ")
		fmt.Scanln(&Dmenor)
		Area := (Dmayor * Dmenor) / Const
		fmt.Printf("El área del rombo es: %d\n", Area)

	case "2":
		var Largo, Ancho int
		fmt.Print("Ingresa el valor del largo del rectángulo: ")
		fmt.Scanln(&Largo)
		fmt.Print("Ingresa el valor del ancho del rectángulo: ")
		fmt.Scanln(&Ancho)
		Area := Largo * Ancho
		fmt.Printf("El área del rectángulo es: %d\n", Area)

	case "3":
		var Lado int
		fmt.Print("Ingresa el valor del lado del cuadrado: ")
		fmt.Scanln(&Lado)
		Area := Lado * Lado
		fmt.Printf("El área del cuadrado es: %d\n", Area)

	case "4":
		var Bmayor, Bmenor, H int
		fmt.Print("Ingresa el valor de la base mayor: ")
		fmt.Scanln(&Bmayor)
		fmt.Print("Ingresa el valor de la base menor: ")
		fmt.Scanln(&Bmenor)
		fmt.Print("Ingresa la altura del trapecio: ")
		fmt.Scanln(&H)
		Area := ((Bmayor + Bmenor) * H) / Const
		fmt.Printf("El área del trapecio es: %d\n", Area)

	default:
		fmt.Println("Opción no válida")
	}
}
