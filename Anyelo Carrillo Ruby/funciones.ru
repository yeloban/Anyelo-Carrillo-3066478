# Multiplicación de elementos de dos arrays

a = [1, 2, 3, 4, 5]
b = [6, 7, 8, 9, 10]
c = []

# Iteramos sobre los índices y multiplicamos los valores correspondientes
a.each_with_index do |valor, index|
  c << valor * b[index]
end

puts c

# Funciones en Ruby

# Definición y llamada de una función
def mostrar_texto
  puts "hola"
end

mostrar_texto

# Función sin parámetros que multiplica valores dentro de ella
def multiplicar
  a = 5
  b = 8
  puts a * b
end

multiplicar

# Variables globales vs. locales en funciones
def multiplicar_externa
  puts $a * $b  # Se usa $ para variables globales
end

$a = 5
$b = 8
multiplicar_externa

# Función con return
def multiplicar_retorno
  a = 5
  b = 8
  return a * b
end

resultado = multiplicar_retorno
puts resultado + 5  # Sumar 5 al resultado de la función

# Devolviendo booleanos desde una función
def validar_dato
  return $a == 5
end

$a = 5
dato = validar_dato
puts dato
