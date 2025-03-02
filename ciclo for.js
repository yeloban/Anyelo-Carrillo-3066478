const readline = require("readline");

// Configurar readline para entrada y salida en la consola
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Solicitar al usuario que ingrese el primer valor
rl.question("Ingrese el primer valor: ", (A) => {
    A = parseInt(A); // Convertir la entrada a número entero

    // Solicitar al usuario que ingrese el segundo valor
    rl.question("Ingrese el segundo valor: ", (C) => {
        C = parseInt(C); // Convertir la entrada a número entero

        // Calcular la potencia de A elevado a C
        let valor = Math.pow(A, C); // Función Math.pow() para calcular potencias

        // Mostrar el resultado
        console.log(`La potencia de ${A} sobre ${C} es: ${valor}`);

        rl.close(); // Cerrar la interfaz readline
    });
});
