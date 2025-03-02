const readline = require('readline-sync'); // Importamos readline-sync para capturar la entrada del usuario
const equipos = [
    "Brasil", "Corea del Sur", "Francia", "Dinamarca",
    "Argentina", "México", "España", "Alemania",
    "Japón", "Inglaterra", "Perú", "USA",
    "Nigeria", "Chile", "Portugal", "Colombia"
];

/**
 * Simula un partido y determina el ganador según los goles ingresados por el usuario.
 */
function jugarPartido(equipo1, equipo2) {
    console.log(`\n${equipo1} vs ${equipo2}`);

    // Solicitar goles al usuario
    let goles1 = parseInt(readline.question(`Ingrese goles de ${equipo1}: `));
    let goles2 = parseInt(readline.question(`Ingrese goles de ${equipo2}: `));

    // Determinar el ganador o si hay empate
    if (goles1 > goles2) {
        console.log(`${equipo1} gana y avanza a la siguiente ronda!`);
        return equipo1;
    } else if (goles2 > goles1) {
        console.log(`${equipo2} gana y avanza a la siguiente ronda!`);
        return equipo2;
    } else {
        console.log("EMPATE! Se jugará tiempo extra...");
        return jugarTiempoExtra(equipo1, equipo2);
    }
}

/**
 * Maneja la fase de tiempo extra en caso de empate.
 */
function jugarTiempoExtra(equipo1, equipo2) {
    let goles1 = parseInt(readline.question(`Tiempo extra - Goles de ${equipo1}: `));
    let goles2 = parseInt(readline.question(`Tiempo extra - Goles de ${equipo2}: `));

    if (goles1 > goles2) {
        console.log(`${equipo1} gana en tiempo extra!`);
        return equipo1;
    } else if (goles2 > goles1) {
        console.log(`${equipo2} gana en tiempo extra!`);
        return equipo2;
    } else {
        console.log("¡Sigue el empate! Se define por penales...");
        return jugarPenales(equipo1, equipo2);
    }
}

/**
 * Simula una tanda de penales hasta que haya un ganador.
 */
function jugarPenales(equipo1, equipo2) {
    while (true) {
        let penales1 = Math.floor(Math.random() * 5) + 1; // Número aleatorio entre 1 y 5
        let penales2 = Math.floor(Math.random() * 5) + 1;

        console.log(`Penales: ${equipo1} ${penales1} - ${penales2} ${equipo2}`);

        if (penales1 > penales2) {
            console.log(`${equipo1} gana por penales!`);
            return equipo1;
        } else if (penales2 > penales1) {
            console.log(`${equipo2} gana por penales!`);
            return equipo2;
        } else {
            console.log("Sigue el empate en penales... ¡Se repiten los tiros!");
        }
    }
}

/**
 * Simula una ronda del torneo y devuelve los equipos clasificados.
 */
function simularRonda(equipos, nombreRonda) {
    console.log(`\n===== ${nombreRonda} =====`);
    let ganadores = [];

    for (let i = 0; i < equipos.length; i += 2) {
        let equipo1 = equipos[i];
        let equipo2 = equipos[i + 1];
        let ganador = jugarPartido(equipo1, equipo2);
        ganadores.push(ganador);
    }

    return ganadores;
}

// Simulación del torneo
let rondaActual = equipos;
let rondas = ["Octavos de Final", "Cuartos de Final", "Semifinal", "Final"];

for (let nombreRonda of rondas) {
    rondaActual = simularRonda(rondaActual, nombreRonda);
}

// Anunciar al campeón
console.log(`\n🏆 EL CAMPEÓN DEL MUNDIAL ES: ${rondaActual[0]} 🏆`);
