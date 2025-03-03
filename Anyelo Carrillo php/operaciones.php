<?php
// Solicitar valores al usuario desde la consola o formulario
echo "Ingrese el primer valor: ";
$A = intval(trim(fgets(STDIN))); // Convertir entrada en entero
echo "Ingrese el segundo valor: ";
$B = intval(trim(fgets(STDIN))); // Convertir entrada en entero

// Realizar operaciones matemáticas básicas
$suma = $A + $B;
$resta = $A - $B;
$multi = $A * $B;

// Manejo de división para evitar error si B es 0
$div = ($B != 0) ? ($A / $B) : "No se puede dividir por 0";

// Comparaciones entre los números
$igual = ($A == $B);
$mayor = ($A > $B);
$menor = ($A < $B);

// Mostrar los resultados en pantalla
echo "\n--- Resultados ---\n";
echo "La suma de los números es: $suma\n";
echo "La resta de los números es: $resta\n";
echo "La multiplicación de los números es: $multi\n";
echo "La división de los números es: $div\n";

if ($igual) {
    echo "Ambos números son iguales.\n";
} else {
    echo "El número menor es: " . ($menor ? $A : $B) . "\n";
    echo "El número mayor es: " . ($mayor ? $A : $B) . "\n";
}
