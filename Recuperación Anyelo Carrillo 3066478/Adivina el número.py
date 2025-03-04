import random

class JuegoAdivinar:
    def __init__(self):
        self.numero_secreto = random.randint(1, 100)  # Número aleatorio entre 1 y 100
        self.intentos = 0

    def adivinar(self, numero):
        self.intentos += 1
        if numero < self.numero_secreto:
            return "El número es mayor. ¡Intenta de nuevo!"
        elif numero > self.numero_secreto:
            return "El número es menor. ¡Intenta de nuevo!"
        else:
            return f"¡Felicidades! Adivinaste el número en {self.intentos} intentos."

# Función principal para jugar
def jugar():
    juego = JuegoAdivinar()
    print("¡Bienvenido al juego de adivinar el número!")
    print("Estoy pensando en un número entre 1 y 100. ¿Puedes adivinarlo?")

    while True:
        try:
            guess = int(input("Ingresa tu número: "))  # El usuario ingresa un número
            resultado = juego.adivinar(guess)
            print(resultado)
            if "Felicidades" in resultado:
                break  # Termina el juego si el usuario adivina
        except ValueError:
            print("Por favor, ingresa un número válido.")

# Iniciar el juego
jugar()
