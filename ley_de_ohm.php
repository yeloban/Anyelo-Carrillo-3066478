<?php
// Solicitar al usuario el voltaje del circuito
echo "Ingrese el valor del voltaje del circuito: ";
$Voltaje = (int) trim(fgets(STDIN)); // Leer entrada y convertirla a número entero

// Solicitar al usuario la resistencia del circuito
echo "Ingrese el valor de la resistencia del circuito: ";
$Resistencia = (int) trim(fgets(STDIN));

// Calcular la intensidad de corriente (Amperaje) usando la Ley de Ohm: I = V / R
$Intensidad = $Voltaje / $Resistencia;

// Mostrar el resultado con el formato adecuado
echo "Al conectar un resistor de R {$Resistencia} ohm a una fuente de V {$Voltaje} volts, circulará una corriente de {$Intensidad} amperios.\n";
