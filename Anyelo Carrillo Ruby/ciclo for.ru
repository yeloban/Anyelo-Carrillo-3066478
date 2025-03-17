# Solicitar el primer número al usuario
print "Ingrese el primer valor: "
a = gets.chomp.to_i  # Captura la entrada y la convierte a un número entero

# Solicitar el segundo número (exponente) al usuario
print "Ingrese el segundo valor: "
c = gets.chomp.to_i  # Captura la entrada y la convierte a un número entero

# Calcular la potencia
valor = a**c  # `**` es el operador de potencia en Ruby

# Mostrar el resultado
puts "La potencia de #{a} sobre #{c} es: #{valor}"
