// Importar el módulo 'readline' para leer la entrada del usuario
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

// Función principal asíncrona para solicitar y procesar los datos del usuario
async function obtenerDatos() {
    // Solicitar el nombre completo
    let a = await preguntar("Escriba sus nombres completos: ");

    // Solicitar el apellido completo
    let b = await preguntar("Escriba sus Apellidos completos: ");

    // Solicitar la profesión
    let c = await preguntar("Escriba su profesión: ");

    // Solicitar el año de nacimiento y convertirlo a número
    let d = parseInt(await preguntar("Escriba su año de nacimiento: "), 10);

    // Calcular la edad restando el año de nacimiento al año actual (2025)
    let e = 2025 - d;

    // Mostrar el resultado
    console.log(`El (La) ${c} ${a} ${b} tiene ${e} años.`);

    // Cerrar la interfaz de lectura
    readline.close();
}

// Ejecutar la función principal
obtenerDatos();
