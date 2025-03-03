<?php
// Definir el arreglo de nombres
$nombres = ["Esteban", "Hans", "Jhon", "Juan Pablo \n\n"];

foreach ($nombres as $nombre) {
    echo $nombre . "\n";
}

// Diccionarios (en PHP se usan arrays asociativos)
$Personas = []; // Lista vacía

$a = ["nombre" => "Esteban", "Edad" => 28];
$b = ["nombre" => "Hans", "Edad" => 27];
$c = ["nombre" => "Jhon", "Edad" => 41];
$d = ["nombre" => "Juan Pablo", "Edad" => 23];

// Agregar los arrays asociativos a la lista
array_push($Personas, $a, $b, $c, $d);

// Recorrer y mostrar los valores
foreach ($Personas as $persona) {
    echo $persona["nombre"] . " " . $persona["Edad"] . "\n";
}

// Operar con valores de diccionarios (arrays asociativos en PHP)
