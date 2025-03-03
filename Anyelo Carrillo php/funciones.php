<?php
// Definir los arreglos
$a = [1, 2, 3, 4, 5];
$b = [6, 7, 8, 9, 10];
$c = [];

// Multiplicar elementos de los arreglos
for ($contador = 0; $contador < count($a); $contador++) {
    $c[] = $a[$contador] * $b[$contador];
}
print_r($c);

// Funciones

// Definición y llamada
function mostrar_texto()
{
    echo "hola\n";
}
mostrar_texto();

function multiplicar()
{
    $a = 5;
    $b = 8;
    echo $a * $b . "\n";
}
multiplicar();

// Ámbito de variables
function multiplicar_con_variables_globales()
{
    global $a, $b;
    echo $a * $b . "\n";
}
$a = 5;
$b = 8;
multiplicar_con_variables_globales();

// Función con retorno de valor
function multiplicar_con_retorno()
{
    $a = 5;
    $b = 8;
    return $a * $b;
}
$Resultado = multiplicar_con_retorno();
echo $Resultado + 5 . "\n";

// Función que devuelve un booleano
function validar_dato()
{
    global $a;
    if ($a == 5) {
        return true;
    } else {
        return false;
    }
}

$a = 5;
$dato = validar_dato();
echo ($dato ? "true" : "false") . "\n";
