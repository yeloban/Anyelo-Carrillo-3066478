// Declaración de variables booleanas
let a = true;
let b = false;

// Operación lógica AND: Solo es true si ambas son true
console.log(a && b); // false, porque b es false

// Declaración de variables numéricas
a = 2;
b = 3;
let c = 4;
let d = 10;

// Evaluación de condiciones con operadores lógicos
console.log(a === b && c < d);  
// false, porque a === b (2 === 3) es false, aunque c < d (4 < 10) es true, el AND necesita ambas condiciones true.

console.log(!(a === b) && c > d);
// false, porque !(a === b) (no 2 === 3) es true, pero c > d (4 > 10) es false. Con && ambas deben ser true para que sea true.
