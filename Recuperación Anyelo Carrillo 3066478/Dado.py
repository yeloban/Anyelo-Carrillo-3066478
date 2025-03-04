import random


class Dado:
    def lanzar(self):
        return random.randint(1, 6)

# Uso
dado = Dado()
print(f"El dado cayó en: {dado.lanzar()}")