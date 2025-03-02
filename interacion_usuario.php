<?php
// Solicitar al usuario su nombre completo
echo "Escriba sus nombres completos: ";
$a = trim(fgets(STDIN)); // Leer entrada y eliminar espacios en blanco extra

// Solicitar al usuario su apellido completo
echo "Escriba sus Apellidos completos: ";
$b = trim(fgets(STDIN));

// Solicitar al usuario su profesión
echo "Escriba su profesión: ";
$c = trim(fgets(STDIN));

// Solicitar al usuario su año de nacimiento
echo "Escriba su año de nacimiento: ";
$d = (int) trim(fgets(STDIN)); // Convertir la entrada a un número entero

// Calcular la edad restando el año de nacimiento al año actual (2025)
$e = 2025 - $d;

// Imprimir el resultado con el formato adecuado
echo "El (La) $c $a $b tiene $e años.\n";
