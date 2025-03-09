#include <iostream>
#include <vector>
#include <map>
#include <string>

int main() {
    // Primera parte: Imprimir nombres
    std::vector<std::string> nombres = {"Esteban", "Hans", "Jhon", "Juan Pablo \n\n"};
    for (const auto& nombre : nombres) {
        std::cout << nombre << std::endl;
    }

    // Segunda parte: Diccionarios (mapas en C++)
    std::vector<std::map<std::string, std::string>> Personas; // Lista de mapas

    // Crear diccionarios (mapas)
    std::map<std::string, std::string> a = {{"nombre", "Esteban"}, {"Edad", "28"}};
    std::map<std::string, std::string> b = {{"nombre", "Hans"}, {"Edad", "27"}};
    std::map<std::string, std::string> c = {{"nombre", "Jhon"}, {"Edad", "41"}};
    std::map<std::string, std::string> d = {{"nombre", "Juan Pablo"}, {"Edad", "23"}};

    // Agregar mapas a la lista
    Personas.push_back(a);
    Personas.push_back(b);
    Personas.push_back(c);
    Personas.push_back(d);

    // Imprimir nombres y edades
    for (const auto& persona : Personas) {
        std::cout << persona.at("nombre") << " " << persona.at("Edad") << std::endl;
    }

    // Tercera parte: Operar con valores de diccionarios
    int sumaEdades = 0;
    for (const auto& persona : Personas) {
        sumaEdades += std::stoi(persona.at("Edad")); // Convertir edad a entero y sumar
    }
    std::cout << "La suma de las edades es: " << sumaEdades << std::endl;

    return 0;
}
