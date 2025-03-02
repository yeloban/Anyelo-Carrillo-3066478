<?php
while (true) {
    try {
        echo "Introduzca un valor mayor a 10: ";
        $a = trim(fgets(STDIN)); // Leer entrada del usuario

        if (!is_numeric($a)) {
            throw new Exception("No es un número válido.");
        }

        $a = intval($a); // Convertir a entero

        if ($a > 10) {
            echo "Es correcto\n";
            break; // Sale del bucle cuando la condición se cumple
        } else {
            echo "Es incorrecto, pruebe de nuevo\n";
        }
    } catch (Exception $e) {
        echo "Por favor, ingrese un número válido\n";
    }
}
?>