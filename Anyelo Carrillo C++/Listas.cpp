#include <iostream>
#include <vector>

int main() {
    std::vector<std::string> Lista = {"Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado"};
    std::cout << Lista[4] << std::endl; // Imprime "viernes"

    for (const auto& dia : Lista) {
        std::cout << dia << " ";
    }
    std::cout << std::endl;

    for (size_t i = 0; i < 3; i++) {
        std::cout << Lista[i] << " ";
    }
    std::cout << std::endl;

    std::vector<std::variant<std::string, int, double, bool>> ListaCompleja = {
        "Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado", 1, 2, 3, std::vector<std::variant<std::string, double, bool>>{"Esteban", 0.2, 2.25, true}
    };

    // Nota: C++ no tiene una forma directa de manejar listas heterogéneas como en Python.
    // Se puede usar std::variant o una estructura personalizada.

    return 0;
}