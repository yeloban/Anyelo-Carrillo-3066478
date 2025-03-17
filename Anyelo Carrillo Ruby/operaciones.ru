# Solicitar los valores al usuario y convertirlos a enteros
print "Ingrese valor: "
a = gets.chomp.to_i  

print "Ingrese valor: "
b = gets.chomp.to_i  

# Operaciones básicas
suma = a + b
puts "La suma de los números es: #{suma}"

resta = a - b
puts "La resta de los números es: #{resta}"

multiplicacion = a * b
puts "La multiplicación de los números es: #{multiplicacion}"

division = a.to_f / b  # Convertimos a flotante para evitar divisiones enteras
puts "La división de los números es: #{division}"

# Comparaciones
igual = a == b
puts "¿Los números son iguales? #{igual}"

# Determinar el mayor y menor
if a > b
  puts "El número menor es: #{b}"
  puts "El número mayor es: #{a}"
else
  puts "El número menor es: #{a}"
  puts "El número mayor es: #{b}"
end
