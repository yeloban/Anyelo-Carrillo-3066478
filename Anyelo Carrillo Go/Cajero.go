package main

import (
	"fmt"
)

// Definición de la estructura Usuario
type Usuario struct {
	Nombre string
	Pin    int
	Saldo  float64
}

// Definición de la estructura Banco
type Banco struct {
	Usuarios []Usuario
}

// Método para autenticar un usuario
func (b *Banco) Autenticar(nombre string, pin int) bool {
	for _, usuario := range b.Usuarios {
		if usuario.Nombre == nombre {
			if usuario.Pin == pin {
				fmt.Println("Estás Logeado")
				return true
			} else {
				fmt.Println("Pin o nombre incorrecto")
				return false
			}
		}
	}
	fmt.Println("El usuario no existe")
	return false
}

// Método para sacar dinero
func (b *Banco) SacarDinero(nombre string, saldo float64) {
	for i, usuario := range b.Usuarios {
		if usuario.Nombre == nombre {
			if usuario.Saldo < saldo {
				fmt.Println("Saldo insuficiente")
				break
			} else if usuario.Saldo >= saldo {
				b.Usuarios[i].Saldo -= saldo
				fmt.Printf("El saldo disponible es %.2f\n", b.Usuarios[i].Saldo)
			}
		}
	}
}

func main() {
	// Crear usuarios
	ana := Usuario{Nombre: "Ana", Pin: 9872, Saldo: 450}
	pablo := Usuario{Nombre: "Pablo", Pin: 5555, Saldo: 200}
	rodrigo := Usuario{Nombre: "Rodrigo", Pin: 2222, Saldo: 900}

	// Crear banco con usuarios
	banco := Banco{Usuarios: []Usuario{ana, pablo, rodrigo}}

	// Bucle principal
	for {
		fmt.Println("Bienvenido al Banco, por favor, identifíquese.")
		fmt.Println("Introduzca el nombre.")
		var nombre string
		fmt.Scanln(&nombre)

		fmt.Println("Introduzca el pin")
		var pin int
		fmt.Scanln(&pin)

		if banco.Autenticar(nombre, pin) {
			for {
				fmt.Println("Por favor, elija una de estas opciones:\n 1. Sacar dinero\n 2. Terminar sesión.")
				var opcion string
				fmt.Scanln(&opcion)

				if opcion == "1" {
					fmt.Println("Introduce cantidad a sacar.")
					var saldo float64
					fmt.Scanln(&saldo)
					banco.SacarDinero(nombre, saldo)
				} else if opcion == "2" {
					fmt.Println("Saliendo del sistema.")
					break
				} else {
					fmt.Println("Opción incorrecta. Por favor, introduzca otra opción.")
				}
			}
		} else {
			fmt.Println("Usuario no autenticado. Por favor, introduzca nombre y pin correctos.")
		}
	}
}
