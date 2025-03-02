<?php
// Definir una variable con valor 2
$a = 2;

// Comprobar si 'a' es igual a 2
if ($a == 2) {
    echo "a vale 2\n"; // Se ejecuta esta línea porque 'a' es igual a 2
} else {
    echo "a es diferente de 2\n";
}

echo "\n"; // Salto de línea para separar secciones

// ------------------------------------------------

// Concatenar condiciones con 'and' y 'not'

// Definir variables con información personal
$Nombre = "Esteban";
$Edad = 28;
$Pais = "Colombia";

// Primera condición: Verifica si todas las condiciones son verdaderas
if ($Nombre == "Esteban" && $Edad == 28 && $Pais == "Colombia") {
    echo "Su nombre es $Nombre, tiene $Edad y es de $Pais\n";
}
// Segunda condición: Verifica si el nombre es 'Esteban', pero la edad NO es 28
elseif ($Nombre == "Esteban" && !($Edad == 28)) {
    echo "Su nombre es Esteban y no tiene 28 años\n";
}
// Tercera condición: Verifica si el nombre NO es 'Esteban', pero la edad es 28
elseif ($Nombre != "Esteban" && $Edad == 28) {
    echo "Su nombre no es Esteban y tiene 28 años\n";
}
// Si ninguna condición anterior se cumple, entra en el bloque 'else'
else {
    echo "No se llama Esteban ni tiene 28 años\n";
}
