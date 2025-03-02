// Definir los arreglos
let a = [1, 2, 3, 4, 5];
let b = [6, 7, 8, 9, 10];
let c = [];

// Multiplicar elementos de los arreglos
for (let i = 0; i < a.length; i++) {
    c.push(a[i] * b[i]);
}
console.log(c);

// Funciones

// Definición y llamada
function mostrarTexto() {
    console.log("hola");
}
mostrarTexto();

function multiplicar() {
    let a = 5;
    let b = 8;
    console.log(a * b);
}
multiplicar();

// Ámbito de variables
function multiplicarConVariablesGlobales() {
    console.log(a * b);
}
let aGlobal = 5;
let bGlobal = 8;
multiplicarConVariablesGlobales();

// Función con retorno de valor
function multiplicarConRetorno() {
    let a = 5;
    let b = 8;
    return a * b;
}
let resultado = multiplicarConRetorno();
console.log(resultado + 5);

// Función que devuelve un booleano
function validarDato() {
    return a === 5;
}

let aBoolean = 5;
let dato = validarDato();
console.log(dato);
