Edad = int(input("ingrese el valor de su edad: "))
print (f"su edad es  {Edad}")

class ValidadorEdad:


 def es_mayor_de_edad(self, edad):
        if Edad >= 18:
            print("Es mayor de edad") 

        else:
         print("Es menor de edad")

# Uso
validador = ValidadorEdad()
print(validador.es_mayor_de_edad(20))  # Es mayor de edad