// Definir el arreglo de nombres
let nombres = ["Esteban", "Hans", "Jhon", "Juan Pablo \n\n"];

for (let nombre of nombres) {
    console.log(nombre);
}

// Diccionarios (Objetos en JavaScript)
let Personas = []; // Lista vacía

let a = { nombre: "Esteban", Edad: 28 };
let b = { nombre: "Hans", Edad: 27 };
let c = { nombre: "Jhon", Edad: 41 };
let d = { nombre: "Juan Pablo", Edad: 23 };

// Agregar objetos a la lista
Personas.push(a);
Personas.push(b);
Personas.push(c);
Personas.push(d);

// Recorrer y mostrar los valores de los objetos
for (let persona of Personas) {
    console.log(persona.nombre, persona.Edad);
}

// Operar con valores de diccionarios (objetos)
