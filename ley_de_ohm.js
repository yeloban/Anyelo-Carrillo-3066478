// Importar el módulo 'readline' para leer datos del usuario en Node.js
const readline = require("readline").createInterface({
    input: process.stdin,
    output: process.stdout
});

// Función para solicitar datos al usuario
function preguntar(pregunta) {
    return new Promise((resolve) => {
        readline.question(pregunta, (respuesta) => {
            resolve(respuesta);
        });
    });
}

// Función principal asíncrona para calcular la corriente
async function calcularCorriente() {
    // Solicitar el voltaje y convertirlo a número
    let Voltaje = parseFloat(await preguntar("Ingrese el valor del voltaje del circuito: "));

    // Solicitar la resistencia y convertirla a número
    let Resistencia = parseFloat(await preguntar("Ingrese el valor de la resistencia del circuito: "));

    // Calcular la intensidad de corriente usando la Ley de Ohm: I = V / R
    let Intensidad = Voltaje / Resistencia;

    // Mostrar el resultado
    console.log(`Al conectar un resistor de R ${Resistencia} ohm a una fuente de V ${Voltaje} volts, circulará una corriente de ${Intensidad} amperios.`);

    // Cerrar la interfaz de lectura
    readline.close();
}

// Ejecutar la función principal
calcularCorriente();
