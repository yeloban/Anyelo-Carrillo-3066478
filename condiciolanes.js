const readline = require("readline");

// Crear la interfaz de entrada y salida
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Definir constantes
const PI = 3.1416;
const CONST_DIV = 2;

// Mostrar el menú de selección
console.log("Seleccione la figura para calcular su área:\n");
console.log("1 para Rombo");
console.log("2 para Rectángulo");
console.log("3 para Cuadrado");
console.log("4 para Trapecio");

// Leer la opción del usuario
rl.question("Ingrese el número de la figura: ", function (Figura) {
    switch (Figura) {
        case "1": // Rombo
            rl.question("Ingrese el valor de la diagonal mayor: ", function (Dmayor) {
                rl.question("Ingrese el valor de la diagonal menor: ", function (Dmenor) {
                    let Area = (parseInt(Dmayor) * parseInt(Dmenor)) / CONST_DIV;
                    console.log(`El área del rombo es: ${Area}`);
                    rl.close();
                });
            });
            break;

        case "2": // Rectángulo
            rl.question("Ingrese el valor del largo del rectángulo: ", function (Largo) {
                rl.question("Ingrese el valor del ancho del rectángulo: ", function (Ancho) {
                    let Area = parseInt(Largo) * parseInt(Ancho);
                    console.log(`El área del rectángulo es: ${Area}`);
                    rl.close();
                });
            });
            break;

        case "3": // Cuadrado
            rl.question("Ingrese el valor del lado del cuadrado: ", function (Lado) {
                let Area = parseInt(Lado) * parseInt(Lado);
                console.log(`El área del cuadrado es: ${Area}`);
                rl.close();
            });
            break;

        case "4": // Trapecio
            rl.question("Ingrese el valor de la base mayor: ", function (Bmayor) {
                rl.question("Ingrese el valor de la base menor: ", function (Bmenor) {
                    rl.question("Ingrese la altura del trapecio: ", function (H) {
                        let Area = ((parseInt(Bmayor) + parseInt(Bmenor)) * parseInt(H)) / CONST_DIV;
                        console.log(`El área del trapecio es: ${Area}`);
                        rl.close();
                    });
                });
            });
            break;

        default: // Opción inválida
            console.log("Opción no válida. Por favor, seleccione un número entre 1 y 4.");
            rl.close();
    }
});
