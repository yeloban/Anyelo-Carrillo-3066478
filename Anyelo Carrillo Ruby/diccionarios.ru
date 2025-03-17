# Creación de un hash (equivalente a diccionario en Python)
coche = {
  "marca" => "Chana",
  "color" => "blanco",
  "modelo" => "sedan",
  "placa" => "DAD190"
}

# Imprimir el valor de una clave
puts coche["color"]

# Cambiar el valor de una clave
coche["color"] = "Negro"
puts coche["color"]

puts coche["marca"]
coche["marca"] = "Renault"
puts coche["marca"]

# Imprimir todo el hash
puts coche
