class Verificador:
    def es_par(self, numero):
        if numero % 2 == 0:
            return "Es par"
        return "Es impar"

# Uso
verificador = Verificador()
print(verificador.es_par(4))  # Es par
print(verificador.es_par(7))  # Es impar