from Guerrero import Guerrero
from Hechicero import Hechicero
from Espada import Espada
from Conjuro import Conjuro

# Crear personajes
guerrero = Guerrero("Thorg", 50)
hechicero = Hechicero("Merlin", 50)

# Crear objetos
Espada = Espada("Espada del Valor", 50)
Conjuro = Conjuro("Bola de Fuego", 30)

# Equipar objetos
guerrero.equipar_arma(Espada)
hechicero.aprender_conjuro(Conjuro)

# Acciones del juego
guerrero.atacar(Hechicero)
hechicero.lanzar_conjuro(Guerrero)

