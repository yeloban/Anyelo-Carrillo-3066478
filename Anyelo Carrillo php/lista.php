<?php
// Definir una lista con los días de la semana
$Lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'];

// Imprimir el quinto elemento de la lista (índice 4, ya que los arrays en PHP comienzan desde 0)
echo $Lista[4] . "\n"; // Salida: Viernes

// Volver a definir la lista y mostrar todos los elementos
$Lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'];
print_r($Lista); // Muestra la lista completa

// Imprimir los primeros 3 elementos (índices del 0 al 2)
$Lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado'];
print_r(array_slice($Lista, 0, 3)); // Salida: ['Lunes', 'Martes', 'Miércoles']

// Crear una lista con distintos tipos de datos, incluyendo una sublista
$Lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado', 1, 2, 3, ['Esteban', 0.2, 2.25, true]];

// Imprimir los primeros 4 elementos y los primeros 3 elementos de la sublista en la posición 9
print_r(array_slice($Lista, 0, 4));
print_r(array_slice($Lista[9], 0, 3));
