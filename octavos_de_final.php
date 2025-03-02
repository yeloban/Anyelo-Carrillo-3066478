<?php
// Función para solicitar los goles de los equipos y determinar el ganador
function jugar_partido($equipo1, $equipo2)
{
    echo "\n$equipo1 vs $equipo2\n";

    // Solicitar al usuario los goles de cada equipo
    $goles1 = (int) readline("Ingrese goles de $equipo1: ");
    $goles2 = (int) readline("Ingrese goles de $equipo2: ");

    if ($goles1 > $goles2) {
        echo "$equipo1 gana y avanza a la siguiente ronda!\n";
        return $equipo1;
    } elseif ($goles2 > $goles1) {
        echo "$equipo2 gana y avanza a la siguiente ronda!\n";
        return $equipo2;
    } else {
        echo "EMPATE! Se jugará tiempo extra...\n";
        return jugar_tiempo_extra($equipo1, $equipo2);
    }
}

// Función para manejar el tiempo extra
function jugar_tiempo_extra($equipo1, $equipo2)
{
    $goles1 = (int) readline("Tiempo extra - Goles de $equipo1: ");
    $goles2 = (int) readline("Tiempo extra - Goles de $equipo2: ");

    if ($goles1 > $goles2) {
        echo "$equipo1 gana en tiempo extra!\n";
        return $equipo1;
    } elseif ($goles2 > $goles1) {
        echo "$equipo2 gana en tiempo extra!\n";
        return $equipo2;
    } else {
        echo "¡Sigue el empate! Se define por penales...\n";
        return jugar_penales($equipo1, $equipo2);
    }
}

// Función para simular penales
function jugar_penales($equipo1, $equipo2)
{
    while (true) {
        $penales1 = rand(1, 5);
        $penales2 = rand(1, 5);

        echo "Penales: $equipo1 $penales1 - $penales2 $equipo2\n";

        if ($penales1 > $penales2) {
            echo "$equipo1 gana por penales!\n";
            return $equipo1;
        } elseif ($penales2 > $penales1) {
            echo "$equipo2 gana por penales!\n";
            return $equipo2;
        } else {
            echo "Sigue el empate en penales... ¡Se repiten los tiros!\n";
        }
    }
}

// Equipos y rondas
$equipos = ["Brasil", "Corea del Sur", "Francia", "Dinamarca", "Argentina", "México", "España", "Alemania", "Japón", "Inglaterra", "Perú", "USA", "Nigeria", "Chile", "Portugal", "Colombia"];
$rondas = ["Octavos de Final", "Cuartos de Final", "Semifinal", "Final"];

foreach ($rondas as $nombre_ronda) {
    echo "\n===== $nombre_ronda =====\n";
    $nueva_ronda = [];

    for ($i = 0; $i < count($equipos); $i += 2) {
        $ganador = jugar_partido($equipos[$i], $equipos[$i + 1]);
        $nueva_ronda[] = $ganador;
    }

    $equipos = $nueva_ronda;
}

// Anunciar al campeón
echo "\n  EL CAMPEÓN DEL MUNDIAL ES: " . $equipos[0] . "\n";
