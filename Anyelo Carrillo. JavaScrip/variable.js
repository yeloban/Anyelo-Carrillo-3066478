// Definimos dos variables numéricas
let a = 10;
let b = 4;

// Función para realizar operaciones y mostrar resultados
function realizarOperacion(operador, resultado) {
    console.log(`La primera variable es: ${a}`);
    console.log(`El signo de la operación es: ${operador}`);
    console.log(`La segunda variable es: ${b}`);
    console.log(`El resultado es: ${resultado}`);
    console.log(`Tipo de dato: ${typeof resultado}`);
}

// Multiplicación
realizarOperacion('*', a * b);

// División
realizarOperacion('/', a / b);

// Suma
realizarOperacion('+', a + b);

// Resta
realizarOperacion('-', a - b);

// Potencia
realizarOperacion('**', a ** b);

// Módulo (residuo de la división)
realizarOperacion('%', a % b);
