# Primer bucle while
contador = 0
while contador < 30
  contador += 1
  puts "Iteración #{contador}"
end

# Combinado con if-else y manejo de errores
loop do
  begin
    print 'Introduzca un valor mayor a 10: '
    a = Integer(gets.chomp)
    if a > 10
      puts 'Es correcto'
      break  # Sale del bucle cuando la condición se cumple
    else
      puts 'Es incorrecto, pruebe de nuevo'
    end
  rescue ArgumentError
    puts 'Por favor, ingrese un número válido'
  end
end