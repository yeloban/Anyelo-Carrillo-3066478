package main

import "fmt"

func main() {
	// Crear una lista (slice) de strings
	Lista := []string{"Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado"}

	// Acceder al elemento en la posición 4 (índice 4)
	fmt.Println(Lista[4]) // Imprime "viernes"

	// Imprimir la lista completa
	fmt.Println(Lista)

	// Obtener un subconjunto de la lista (desde el índice 0 hasta el 3, excluyendo el 3)
	fmt.Println(Lista[0:3]) // Imprime ["Lunes", "Martes", "Miercoles"]

	// Crear una lista heterogénea (en Go, esto no es directamente posible, se usa una estructura o interface{})
	// Para simular una lista heterogénea, se puede usar []interface{}
	ListaHeterogenea := []interface{}{
		"Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado", 1, 2, 3,
		[]interface{}{"Esteban", 0.2, 2.25, true},
	}

	// Acceder a los primeros 4 elementos y a un subconjunto de la lista anidada
	fmt.Println(ListaHeterogenea[0:4], ListaHeterogenea[9].([]interface{})[0:3])
}
