package main

import "fmt"

func main() {
	// Definición de slices (equivalentes a las tuplas y listas en Python)
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	var c []int

	// Multiplicar elementos de los slices y almacenar el resultado en c
	for contador := 0; contador < len(a); contador++ {
		c = append(c, a[contador]*b[contador])
	}
	fmt.Println(c)

	// Funciones

	// Definición y llamada de una función sin parámetros ni retorno
	mostrarTexto := func() {
		fmt.Println("hola")
	}
	mostrarTexto()

	// Definición y llamada de una función que multiplica dos valores internos
	multiplicar := func() {
		a := 5
		b := 8
		fmt.Println(a * b)
	}
	multiplicar()

	// Ámbito de las variables
	// Las variables declaradas fuera de la función son accesibles dentro de la función
	multiplicarConVariablesExternas := func() {
		fmt.Println(aExterna * bExterna)
	}
	aExterna := 5
	bExterna := 8
	multiplicarConVariablesExternas()

	// Funciones con retorno
	multiplicarConRetorno := func() int {
		a := 5
		b := 8
		return a * b
	}
	Resultado := multiplicarConRetorno()
	fmt.Println(Resultado + 5)

	// Funciones que devuelven cualquier tipo de dato
	validarDato := func() bool {
		if aExterna == 5 {
			return true
		}
		return false
	}
	dato := validarDato()
	fmt.Println(dato)
}
