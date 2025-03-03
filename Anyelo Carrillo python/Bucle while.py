Contador = 0
while Contador < 30:
    Contador += 1
    print('Iteración', Contador)

# Combinado con if-else y manejo de errores
while True:
    try:
        a = int(input('Introduzca un valor mayor a 10: '))
        if a > 10:
            print('Es correcto')
            break  # Sale del bucle cuando la condición se cumple
        else:
            print('Es incorrecto, pruebe de nuevo')
    except ValueError:
        print('Por favor, ingrese un número válido')
