# Lista de nombres
nombres = ['Santiago', 'Hans', 'Jhon', 'Juan Pablo']

# Imprimir nombres
nombres.each do |nombre|  # Iteramos sobre cada nombre en la lista
  puts nombre  # Imprimimos el nombre
end

# Lista vacía para almacenar los diccionarios (hashes en Ruby)
personas = []

# Creación de los hashes con los datos de cada persona
a = {nombre: 'Santiago', edad: 23}
b = {nombre: 'Hans', edad: 27}
c = {nombre: 'Jhon', edad: 41}
d = {nombre: 'Juan Pablo', edad: 23}

# Agregar los hashes a la lista
personas.push(a, b, c, d)  # push permite agregar múltiples elementos a la lista

# Imprimir nombre y edad de cada persona
personas.each do |persona|  # Iteramos sobre cada hash en la lista personas
  puts "#{persona[:nombre]} #{persona[:edad]}"  # Accedemos a los valores con `:clave`
end

# Operaciones con valores de diccionarios

# Sumar todas las edades
suma_edades = personas.sum { |persona| persona[:edad] }  # sumamos todas las edades

# Calcular el promedio de edades
promedio_edades = suma_edades.to_f / personas.size  # Convertimos a float para obtener decimales

# Imprimir resultados
puts "Suma de edades: #{suma_edades}"
puts "Promedio de edades: #{'%.2f' % promedio_edades}"  # Formateamos a 2 decimales
