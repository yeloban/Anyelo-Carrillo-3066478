a = 10
b = 4

puts "La primera variable es: #{a}"
puts "El signo de la operación es: *"
puts "La segunda variable es: #{b}"
c = a * b
puts "El resultado es: #{c}"
puts c.class

puts "La primera variable es: #{a}"
puts "El signo de la operación es: /"
puts "La segunda variable es: #{b}"
c = a.to_f / b  # Convierte 'a' a float para obtener una división decimal
puts "El resultado es: #{c}"
puts c.class

puts "La primera variable es: #{a}"
puts "El signo de la operación es: +"
puts "La segunda variable es: #{b}"
c = a + b
puts "El resultado es: #{c}"
puts c.class

puts "La primera variable es: #{a}"
puts "El signo de la operación es: -"
puts "La segunda variable es: #{b}"
c = a - b
puts "El resultado es: #{c}"
puts c.class

puts "La primera variable es: #{a}"
puts "El signo de la operación es: **"
puts "La segunda variable es: #{b}"
c = a**b
puts "El resultado es: #{c}"
puts c.class

puts "La primera variable es: #{a}"
puts "El signo de la operación es: %"
puts "La segunda variable es: #{b}"
c = a % b
puts "El resultado es: #{c}"
puts c.class
