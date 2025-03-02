// Parte 1: Bucle while para contar hasta 30
let contador = 0;
while (contador < 30) {
    contador++;
    console.log('Iteración', contador);
}

// Parte 2: Bucle while con if-else y manejo de errores
while (true) {
    try {
        // Solicitar entrada al usuario (en un entorno de navegador)
        let a = parseInt(prompt('Introduzca un valor mayor a 10: '));

        if (a > 10) {
            console.log('Es correcto');
            break; // Sale del bucle cuando la condición se cumple
        } else {
            console.log('Es incorrecto, pruebe de nuevo');
        }
    } catch (error) {
        console.log('Por favor, ingrese un número válido');
    }
}