# Solicitar datos al usuario
print "Ingrese el valor del voltaje del circuito: "
voltaje = gets.chomp.to_f  # Convierte la entrada a número decimal (float)

print "Ingrese el valor de la resistencia del circuito: "
resistencia = gets.chomp.to_f  # Convierte la entrada a número decimal (float)

# Calcular la intensidad (amperaje)
intensidad = voltaje / resistencia  

# Mostrar el resultado con interpolación de cadenas
puts "Al conectar un resistor de R#{resistencia} ohm a una fuente de V#{voltaje} voltios, circulará una corriente de #{intensidad} amperios."
# Output