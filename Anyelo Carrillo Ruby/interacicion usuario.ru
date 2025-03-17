# Solicitar información al usuario
print "Escriba sus nombres completos: "
a = gets.chomp  # Captura el nombre

print "Escriba sus apellidos completos: "
b = gets.chomp  # Captura los apellidos

print "Escriba su profesión: "
c = gets.chomp  # Captura la profesión

print "Escriba su año de nacimiento: "
d = gets.chomp.to_i  # Convierte la entrada en un número entero

# Calcular la edad
e = 2025 - d

# Mostrar la información formateada
puts "El (La) #{c} #{a} #{b} tiene #{e} años"
