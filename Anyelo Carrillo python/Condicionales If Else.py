# Definir una variable con valor 2
a = 2

# Comprobar si 'a' es igual a 2
if a == 2:
    print('a vale 2')  # Se ejecuta esta línea porque 'a' es igual a 2
else:
    print('a es diferente de 2')

# ------------------------------------------------

# Concatenar condiciones con 'and' y 'not'

# Definir variables con información personal
Nombre = 'Esteban'
Edad = 28
Pais = 'Colombia'

# Primera condición: Verifica si todas las condiciones son verdaderas
if Nombre == 'Esteban' and Edad == 28 and Pais == 'Colombia':
    print('Su nombre es', Nombre, 'tiene', Edad, 'y es de', Pais)

# Segunda condición: Verifica si el nombre es 'Esteban', pero la edad NO es 28
elif Nombre == 'Esteban' and not Edad == 28:
    print('Su nombre es Esteban y no tiene 28 años')

# Tercera condición: Verifica si el nombre NO es 'Esteban', pero la edad es 28
elif Nombre != 'Esteban' and Edad == 28:
    print('Su nombre no es Esteban y tiene 28 años')

# Si ninguna condición anterior se cumple, entra en el bloque 'else'
else:
    print('No se llama Esteban ni tiene 28 años')
         