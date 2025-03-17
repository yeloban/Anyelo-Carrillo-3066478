# Asignar valor a la variable
a = 2

# Condicional simple
if a == 2
  puts "a vale 2"
else
  puts "a es diferente de 2"
end

# Concatenar condiciones

nombre = "Santiago"
edad = 23
pais = "Colombia"

if nombre == "Santiago" && edad == 23 && pais == "Colombia"
  puts "Su nombre es #{nombre}, tiene #{edad} y es de #{pais}"
elsif nombre == "Santiago" && edad != 23
  puts "Su nombre es Santiago y no tiene 23 años"
elsif nombre != "Santiago" && edad == 23
  puts "Su nombre no es Santiago y tiene 23 años"
else
  puts "No se llama Santiago ni tiene 23 años"
end
