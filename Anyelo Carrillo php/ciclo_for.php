<?php
// Solicitar al usuario que ingrese el primer valor
echo "Ingrese el primer valor: ";
$A = trim(fgets(STDIN)); // Leer entrada

// Solicitar al usuario que ingrese el segundo valor
echo "Ingrese el segundo valor: ";
$C = trim(fgets(STDIN)); // Leer entrada

// Validar que sean números
if (!is_numeric($A) || !is_numeric($C)) {
    echo "Error: Ingrese solo números.\n";
    exit;
}

// Convertir a enteros
$A = (int) $A;
$C = (int) $C;

// Calcular la potencia de A elevado a C
$valor = pow($A, $C);

// Mostrar el resultado
echo "La potencia de $A sobre $C es: $valor\n";
