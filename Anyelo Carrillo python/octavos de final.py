import random

# Lista de equipos en octavos de final
equipos = [
    "Brasil", "Corea del Sur", "Francia", "Dinamarca",
    "Argentina", "México", "España", "Alemania",
    "Japón", "Inglaterra", "Perú", "USA",
    "Nigeria", "Chile", "Portugal", "Colombia"
]

def jugar_partido(equipo1, equipo2):
    """Solicita al usuario los goles de cada equipo y determina un ganador."""
    print(f"\n{equipo1} vs {equipo2}")
    
    goles1 = int(input(f"Ingrese goles de {equipo1}: "))
    goles2 = int(input(f"Ingrese goles de {equipo2}: "))
    
    if goles1 > goles2:
        print(f" {equipo1} gana y avanza a la siguiente ronda!")
        return equipo1
    elif goles2 > goles1:
        print(f" {equipo2} gana y avanza a la siguiente ronda!")
        return equipo2
    else:
        print(" EMPATE! Se jugará tiempo extra...")
        return jugar_tiempo_extra(equipo1, equipo2)

def jugar_tiempo_extra(equipo1, equipo2):
    """Juega tiempo extra y, si sigue el empate, define por penales."""
    goles1 = int(input(f"Tiempo extra - Goles de {equipo1}: "))
    goles2 = int(input(f"Tiempo extra - Goles de {equipo2}: "))

    if goles1 > goles2:
        print(f" {equipo1} gana en tiempo extra!")
        return equipo1
    elif goles2 > goles1:
        print(f" {equipo2} gana en tiempo extra!")
        return equipo2
    else:
        print(" ¡Sigue el empate! Se define por penales...")
        return jugar_penales(equipo1, equipo2)

def jugar_penales(equipo1, equipo2):
    """Simula penales hasta que haya un ganador."""
    while True:
        penales1 = random.randint(1, 5)
        penales2 = random.randint(1, 5)
        
        print(f" Penales: {equipo1} {penales1} - {penales2} {equipo2}")
        
        if penales1 > penales2:
            print(f" {equipo1} gana por penales!")
            return equipo1
        elif penales2 > penales1:
            print(f" {equipo2} gana por penales!")
            return equipo2
        else:
            print(" Sigue el empate en penales... ¡Se repiten los tiros!")

def simular_ronda(equipos, nombre_ronda):
    """Juega una ronda completa y devuelve los equipos clasificados."""
    print(f"\n===== {nombre_ronda} =====")
    ganadores = []
    
    for i in range(0, len(equipos), 2):
        equipo1 = equipos[i]
        equipo2 = equipos[i + 1]
        ganador = jugar_partido(equipo1, equipo2)
        ganadores.append(ganador)
    
    return ganadores

# Simulación del torneo
ronda_actual = equipos
rondas = ["Octavos de Final", "Cuartos de Final", "Semifinal", "Final"]

for nombre_ronda in rondas:
    ronda_actual = simular_ronda(ronda_actual, nombre_ronda)

# Anunciar al campeón
print(f"\n  EL CAMPEÓN DEL MUNDIAL ES: {ronda_actual[0]} ")