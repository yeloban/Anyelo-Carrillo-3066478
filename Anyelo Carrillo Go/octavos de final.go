package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Lista de equipos en octavos de final
var equipos = []string{
	"Brasil", "Corea del Sur", "Francia", "Dinamarca",
	"Argentina", "México", "España", "Alemania",
	"Japón", "Inglaterra", "Perú", "USA",
	"Nigeria", "Chile", "Portugal", "Colombia",
}

// Función para jugar un partido
func jugarPartido(equipo1, equipo2 string) string {
	fmt.Printf("\n%s vs %s\n", equipo1, equipo2)

	var goles1, goles2 int
	fmt.Printf("Ingrese goles de %s: ", equipo1)
	fmt.Scanln(&goles1)
	fmt.Printf("Ingrese goles de %s: ", equipo2)
	fmt.Scanln(&goles2)

	if goles1 > goles2 {
		fmt.Printf(" %s gana y avanza a la siguiente ronda!\n", equipo1)
		return equipo1
	} else if goles2 > goles1 {
		fmt.Printf(" %s gana y avanza a la siguiente ronda!\n", equipo2)
		return equipo2
	} else {
		fmt.Println(" EMPATE! Se jugará tiempo extra...")
		return jugarTiempoExtra(equipo1, equipo2)
	}
}

// Función para jugar tiempo extra
func jugarTiempoExtra(equipo1, equipo2 string) string {
	var goles1, goles2 int
	fmt.Printf("Tiempo extra - Goles de %s: ", equipo1)
	fmt.Scanln(&goles1)
	fmt.Printf("Tiempo extra - Goles de %s: ", equipo2)
	fmt.Scanln(&goles2)

	if goles1 > goles2 {
		fmt.Printf(" %s gana en tiempo extra!\n", equipo1)
		return equipo1
	} else if goles2 > goles1 {
		fmt.Printf(" %s gana en tiempo extra!\n", equipo2)
		return equipo2
	} else {
		fmt.Println(" ¡Sigue el empate! Se define por penales...")
		return jugarPenales(equipo1, equipo2)
	}
}

// Función para jugar penales
func jugarPenales(equipo1, equipo2 string) string {
	rand.Seed(time.Now().UnixNano()) // Inicializar la semilla del generador de números aleatorios

	for {
		penales1 := rand.Intn(5) + 1
		penales2 := rand.Intn(5) + 1

		fmt.Printf(" Penales: %s %d - %d %s\n", equipo1, penales1, penales2, equipo2)

		if penales1 > penales2 {
			fmt.Printf(" %s gana por penales!\n", equipo1)
			return equipo1
		} else if penales2 > penales1 {
			fmt.Printf(" %s gana por penales!\n", equipo2)
			return equipo2
		} else {
			fmt.Println(" Sigue el empate en penales... ¡Se repiten los tiros!")
		}
	}
}

// Función para simular una ronda
func simularRonda(equipos []string, nombreRonda string) []string {
	fmt.Printf("\n===== %s =====\n", nombreRonda)
	ganadores := []string{}

	for i := 0; i < len(equipos); i += 2 {
		equipo1 := equipos[i]
		equipo2 := equipos[i+1]
		ganador := jugarPartido(equipo1, equipo2)
		ganadores = append(ganadores, ganador)
	}

	return ganadores
}

func main() {
	// Simulación del torneo
	rondaActual := equipos
	rondas := []string{"Octavos de Final", "Cuartos de Final", "Semifinal", "Final"}

	for _, nombreRonda := range rondas {
		rondaActual = simularRonda(rondaActual, nombreRonda)
	}

	// Anunciar al campeón
	fmt.Printf("\n  EL CAMPEÓN DEL MUNDIAL ES: %s \n", rondaActual[0])
}
