// Definir una variable con valor 2
let a = 2;

// Comprobar si 'a' es igual a 2
if (a === 2) {
    console.log("a vale 2"); // Se ejecuta esta línea porque 'a' es igual a 2
} else {
    console.log("a es diferente de 2");
}

console.log("\n"); // Salto de línea para separar secciones

// ------------------------------------------------

// Concatenar condiciones con '&&' y '!'

// Definir variables con información personal
let Nombre = "Esteban";
let Edad = 28;
let Pais = "Colombia";

// Primera condición: Verifica si todas las condiciones son verdaderas
if (Nombre === "Esteban" && Edad === 28 && Pais === "Colombia") {
    console.log(`Su nombre es ${Nombre}, tiene ${Edad} años y es de ${Pais}`);
} 
// Segunda condición: Verifica si el nombre es 'Esteban', pero la edad NO es 28
else if (Nombre === "Esteban" && !(Edad === 28)) {
    console.log("Su nombre es Esteban y no tiene 28 años");
} 
// Tercera condición: Verifica si el nombre NO es 'Esteban', pero la edad es 28
else if (Nombre !== "Esteban" && Edad === 28) {
    console.log("Su nombre no es Esteban y tiene 28 años");
} 
// Si ninguna condición anterior se cumple, entra en el bloque 'else'
else {
    console.log("No se llama Esteban ni tiene 28 años");
}
