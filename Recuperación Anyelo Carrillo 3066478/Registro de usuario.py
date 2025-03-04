class Usuario:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

class SistemaRegistro:
    def __init__(self):
        self.usuarios = []

    def registrar_usuario(self, nombre, edad):
        usuario = Usuario(nombre, edad)
        self.usuarios.append(usuario)

    def mostrar_usuarios(self):
        for usuario in self.usuarios:
            print(f"Nombre: {usuario.nombre}, Edad: {usuario.edad}")

# Uso
sistema = SistemaRegistro()
sistema.registrar_usuario("Ana", 25)
sistema.registrar_usuario("Luis", 30)
sistema.mostrar_usuarios()