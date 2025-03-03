<?php
// Diccionario en PHP (se usa un array asociativo para almacenar datos del coche)
$coche = array(
    "marca" => "Ford",
    "color" => "rojo",
    "modelo" => "sedan",
    "placa" => "ROS227"
);

// Imprimir el color actual del coche
echo $coche["color"] . "\n";

// Cambiar el valor del color
$coche["color"] = "verde";
echo $coche["color"] . "\n"; // Ahora el color es verde

// Imprimir la marca actual del coche
echo $coche["marca"] . "\n";

// Cambiar el valor de la marca
$coche["marca"] = "Renault";
echo $coche["marca"] . "\n"; // Ahora la marca es Renault

// Mostrar el diccionario completo con los cambios realizados
print_r($coche);
