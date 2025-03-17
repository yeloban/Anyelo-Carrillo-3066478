# Definir una lista (Array en Ruby)
lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado']

# Acceder a un elemento específico por índice
puts lista[4]  # Traerá el dato según la posición (índice empieza en 0)

# Imprimir la lista completa
puts lista.inspect

# Slicing (Extraer una parte de la lista)
puts lista[0..2].inspect  # En Ruby, el rango 0..2 incluye los tres primeros elementos

# Lista con elementos mixtos, incluyendo una lista dentro de otra
lista = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado', 1, 2, 3, ['Santiago', 0.2, 2.25, true]]

# Mostrar un rango de elementos y acceder a un sub-array
puts "#{lista[0..3].inspect}, #{lista[9][0..2].inspect}"
