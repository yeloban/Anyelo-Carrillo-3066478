#include <iostream>

int main() {
    int Voltaje, Resistencia;
    std::cout << "Ingrese el valor del voltaje del circuito: ";
    std::cin >> Voltaje;
    std::cout << "Ingrese el valor de la resistencia del circuito: ";
    std::cin >> Resistencia;

    double Intensidad = static_cast<double>(Voltaje) / Resistencia;
    std::cout << "Al conectar un resistor de R " << Resistencia << " ohm a una fuente de V " << Voltaje << " voltios, circulará una corriente de " << Intensidad << " amperios" << std::endl;

    return 0;
}