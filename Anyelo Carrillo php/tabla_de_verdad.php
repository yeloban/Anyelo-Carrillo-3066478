<?php
// Declaración de variables booleanas
$a = true;
$b = false;

// Operación lógica AND: Solo es true si ambas son true
echo ($a && $b) ? "true\n" : "false\n"; // false, porque $b es false

// Declaración de variables numéricas
$a = 2;
$b = 3;
$c = 4;
$d = 10;

// Evaluación de condiciones con operadores lógicos
echo ($a == $b && $c < $d) ? "true\n" : "false\n";
// false, porque $a == $b (2 == 3) es false, aunque $c < $d (4 < 10) es true, el AND necesita ambas condiciones true.

echo (!($a == $b) && $c > $d) ? "true\n" : "false\n";
// false, porque !( $a == $b ) (no 2 == 3) es true, pero $c > $d (4 > 10) es false. Con && ambas deben ser true para que sea true.
