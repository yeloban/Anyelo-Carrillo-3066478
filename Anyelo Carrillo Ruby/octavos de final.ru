# Requerir la biblioteca de números aleatorios
require 'securerandom'

# Lista de partidos (array de tuplas en Python → array de arrays en Ruby)
partidos = [
  ["Países Bajos", "Corea del Sur"], ["Senegal", "Portugal"], ["Inglaterra", "Suiza"], 
  ["USA", "Brasil"], ["Argentina", "Croacia"], ["Polonia", "Marruecos"], 
  ["Francia", "España"], ["Australia", "Japón"]
]

# Diccionario en Python → Hash en Ruby
resultados = {}

# Lista para almacenar los equipos que avanzan a la siguiente ronda
equipos_pasan = []

# Iterar sobre cada partido
partidos.each_with_index do |(equipo1, equipo2), index|
  resultado = rand(0..2)  # Generar un resultado aleatorio (0, 1 o 2)
  
  # Almacenar el resultado en el hash
  resultados[index] = resultado

  # Imprimir el resultado del partido y decidir qué equipo avanza
  if resultado == 2
    puts "#{equipo1} gana contra #{equipo2}"
    equipos_pasan << equipo1  # `<<` en Ruby es como `append()` en Python
  elsif resultado == 1
    puts "Empate entre #{equipo1} y #{equipo2}"
    # No se agrega ningún equipo en caso de empate
  else
    puts "#{equipo1} pierde contra #{equipo2}"
    equipos_pasan << equipo2
  end
end

# Imprimir todos los resultados
puts "\nResultados finales: #{resultados.inspect}"

# Imprimir los equipos que pasan a la siguiente ronda
puts "\nEquipos que pasan a la siguiente ronda:"
equipos_pasan.each { |equipo| puts equipo }
