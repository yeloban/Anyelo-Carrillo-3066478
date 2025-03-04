class Calculadora:
    def sumar(self, a, b):
        return a + b

    def restar(self, a, b):
        return a - b

    def multiplicar(self, a, b):
        return a * b

    def dividir(self, a, b):
        if b == 0:
            return "Error: División por cero"
        return a / b

# Uso
calc = Calculadora()
print(calc.sumar(5, 9))  # 8
print(calc.restar(9, 4)) # 5
print(calc.multiplicar(5, 4)) #20
print(calc.dividir(10, 0))  # Error: División por cero