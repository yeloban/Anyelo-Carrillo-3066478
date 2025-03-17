# Solicitar la figura al usuario
print "Seleccione la figura para calcular su área: \n\n 1 para rombo\n 2 para rectángulo \n 3 para cuadrado \n 4 para trapecio \n\n: "
figura = gets.chomp  # Captura la opción seleccionada

pi = 3.1416
const = 2

# Evaluar la opción seleccionada usando case-when
case figura
when '1'
  print "Ingresa el valor de la diagonal mayor: "
  d_mayor = gets.chomp.to_f  # Convierte a flotante
  print "Ingresa el valor de la diagonal menor: "
  d_menor = gets.chomp.to_f

  area = (d_mayor * d_menor) / const  # Fórmula correcta del área del rombo
  puts "El área del rombo es: #{area}"

when '2'
  print "Ingresa el valor del largo del rectángulo: "
  largo = gets.chomp.to_f
  print "Ingresa el valor del ancho del rectángulo: "
  ancho = gets.chomp.to_f

  area = largo * ancho
  puts "El área del rectángulo es: #{area}"

when '3'
  print "Ingresa el valor del lado del cuadrado: "
  lado = gets.chomp.to_f

  area = lado**2
  puts "El área del cuadrado es: #{area}"

when '4'
  print "Ingresa el valor de la base mayor: "
  b_mayor = gets.chomp.to_f
  print "Ingresa el valor de la base menor: "
  b_menor = gets.chomp.to_f
  print "Ingrese la altura del trapecio: "
  h = gets.chomp.to_f

  area = ((b_mayor + b_menor) * h) / 2
  puts "El área del trapecio es: #{area}"

else
  puts "Opción inválida, por favor seleccione una opción válida."
end
