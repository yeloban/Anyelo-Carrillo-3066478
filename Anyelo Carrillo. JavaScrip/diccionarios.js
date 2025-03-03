// Diccionario en JavaScript (se usa un objeto para almacenar datos del coche)
let coche = {
    marca: "Ford",
    color: "rojo",
    modelo: "sedan",
    placa: "ROS227"
};

// Imprimir el color actual del coche
console.log(coche.color);

// Cambiar el valor del color
coche.color = "verde";
console.log(coche.color); // Ahora el color es verde

// Imprimir la marca actual del coche
console.log(coche.marca);

// Cambiar el valor de la marca
coche.marca = "Renault";
console.log(coche.marca); // Ahora la marca es Renault

// Mostrar el diccionario completo con los cambios realizados
console.log(coche);
