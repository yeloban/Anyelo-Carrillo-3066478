package main

import (
	"fmt"
)

func main() {
	// Lista de nombres
	nombres := []string{"Esteban", "Hans", "Jhon", "Juan Pablo \n\n"}
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	// Diccionarios (en Go se usan mapas y estructuras)
	type Persona struct {
		Nombre string
		Edad   int
	}

	// Lista vacía de personas
	personas := []Persona{}

	// Crear instancias de Persona
	a := Persona{Nombre: "Esteban", Edad: 28}
	b := Persona{Nombre: "Hans", Edad: 27}
	c := Persona{Nombre: "Jhon", Edad: 41}
	d := Persona{Nombre: "Juan Pablo", Edad: 23}

	// Agregar personas a la lista
	personas = append(personas, a)
	personas = append(personas, b)
	personas = append(personas, c)
	personas = append(personas, d)

	// Recorrer la lista de personas
	for _, persona := range personas {
		fmt.Println(persona.Nombre, persona.Edad)
	}

	// Operar con valores de diccionarios (estructuras)
	sumaEdades := 0
	for _, persona := range personas {
		sumaEdades += persona.Edad
	}
	fmt.Printf("La suma de las edades es: %d\n", sumaEdades)
}
