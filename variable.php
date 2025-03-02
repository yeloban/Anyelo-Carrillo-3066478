<?php
// Definimos dos variables numéricas
$a = 10;
$b = 4;

// Función para realizar operaciones y mostrar resultados
function realizarOperacion($operador, $resultado)
{
    echo "La primera variable es: $GLOBALS[a]\n";
    echo "El signo de la operación es: $operador\n";
    echo "La segunda variable es: $GLOBALS[b]\n";
    echo "El resultado es: $resultado\n";
    echo "Tipo de dato: " . gettype($resultado) . "\n\n";
}

// Multiplicación
realizarOperacion('*', $a * $b);

// División
realizarOperacion('/', $a / $b);

// Suma
realizarOperacion('+', $a + $b);

// Resta
realizarOperacion('-', $a - $b);

// Potencia
realizarOperacion('**', pow($a, $b)); // Usamos pow() en PHP

// Módulo (residuo de la división)
realizarOperacion('%', $a % $b);
