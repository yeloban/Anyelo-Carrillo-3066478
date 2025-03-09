

// Clase Usuario
class Usuario {
public:
    std::string nombre;
    int pin;
    int saldo;

    // Constructor
    Usuario(std::string nombre, int pin, int saldo) 
        : nombre(nombre), pin(pin), saldo(saldo) {}
};

// Clase Banco
class Banco {
public:
    std::vector<Usuario> usuarios;

    // Constructor
    Banco(std::vector<Usuario> usuarios = {}) 
        : usuarios(usuarios) {}

    // Método para autenticar usuarios
    bool autenticar(std::string nombre, int pin) {
        for (auto& usuario : usuarios) {
            if (usuario.nombre == nombre) {
                if (usuario.pin == pin) {
                    std::cout << "Estas Logeado" << std::endl;
                    return true;
                } else {
                    std::cout << "Pin o nombre incorrecto" << std::endl;
                    return false;
                }
            }
        }
        std::cout << "El usuario no existe" << std::endl;
        return false;
    }

    // Método para sacar dinero
    void sacar_dinero(std::string nombre, int saldo) {
        for (auto& usuario : usuarios) {
            if (usuario.nombre == nombre) {
                if (usuario.saldo < saldo) {
                    std::cout << "Saldo insuficiente" << std::endl;
                    break;
                } else if (usuario.saldo >= saldo) {
                    usuario.saldo -= saldo;
                    std::cout << "El saldo disponible es " << usuario.saldo << std::endl;
                }
            }
        }
    }
}

int main() {
    // Crear usuarios
    Usuario ana("Ana", 9872, 450);
    Usuario pablo("Pablo", 5555, 200);
    Usuario rodrigo("Rodrigo", 2222, 900);

    // Crear banco y agregar usuarios
    Banco banco({ana, pablo, rodrigo});

    // Menú de autenticación y operaciones
    while (true) {
        std::cout << "Bienvenido al Banco, por favor, identifiquese." << std::endl;
        std::cout << "Introduzca el nombre: ";
        std::string nombre;
        std::cin >> nombre;

        std::cout << "Introduzca el pin: ";
        int pin;
        std::cin >> pin;

        if (banco.autenticar(nombre, pin)) {
            while (true) {
                std::cout << "Por favor, elija una de estas opciones:\n 1. Sacar dinero\n 2. Terminar sesion." << std::endl;
                std::string opcion;
                std::cin >> opcion;

                if (opcion == "1") {
                    std::cout << "Introduce cantidad a sacar: ";
                    int saldo;
                    std::cin >> saldo;
                    banco.sacar_dinero(nombre, saldo);
                } else if (opcion == "2") {
                    std::cout << "Saliendo del sistema." << std::endl;
                    break;
                } else {
                    std::cout << "Opcion incorrecta. Por favor, introduzca otra opcion." << std::endl;
                }
            }
        } else {
            std::cout << "Usuario no autenticado. Por favor, introduzca nombre y pin correctos." << std::endl;
        }
    }

    return 0;
}
