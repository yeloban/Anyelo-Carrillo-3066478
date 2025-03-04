class GestorTareas:
    def __init__(self):
        self.tareas = []

    def agregar_tarea(self, tarea):
        self.tareas.append(tarea)

    def mostrar_tareas(self):
        if not self.tareas:
            print("No hay tareas pendientes.")
        else:
            for i, tarea in enumerate(self.tareas, 1):
                print(f"{i}. {tarea}")

# Uso
gestor = GestorTareas()
gestor.agregar_tarea("Hacer la compra")
gestor.agregar_tarea("Estudiar Python")
gestor.agregar_tarea("Estudiar Mongo")
gestor.agregar_tarea("Estudiar SQ")
gestor.mostrar_tareas()