package main

import "fmt"

func main() {
	// Creación de un mapa (equivalente a un diccionario en Python)
	coche := map[string]string{
		"marca":  "Ford",
		"color":  "rojo",
		"modelo": "sedan",
		"placa":  "ROS227",
	}

	// Acceder e imprimir el valor de la clave "color"
	fmt.Println(coche["color"])

	// Cambiar el valor de la clave "color"
	coche["color"] = "verde"
	fmt.Println(coche["color"])

	// Acceder e imprimir el valor de la clave "marca"
	fmt.Println(coche["marca"])

	// Cambiar el valor de la clave "marca"
	coche["marca"] = "Renault"
	fmt.Println(coche["marca"])

	// Imprimir todo el mapa
	fmt.Println(coche)
}
