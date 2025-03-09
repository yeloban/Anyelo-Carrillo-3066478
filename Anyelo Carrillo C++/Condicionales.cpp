#include <iostream>

int main() {
    std::string Figura;
    std::cout << "Seleccione la figura a calcular su área:\n\n1 para rombo\n2 para rectángulo\n3 para cuadrado\n4 para trapecio\n: ";
    std::cin >> Figura;

    const double Pi = 3.1416;
    const int const = 2;

    if (Figura == "1") {
        int Dmayor, Dmenor;
        std::cout << "Ingresa el valor de la diagonal mayor: ";
        std::cin >> Dmayor;
        std::cout << "Ingresa el valor de la diagonal menor: ";
        std::cin >> Dmenor;
        double Area = (Dmayor + Dmenor) / const;
        std::cout << "El área del rombo es: " << Area << std::endl;
    } else if (Figura == "2") {
        int Largo, Ancho;
        std::cout << "Ingresa el valor del largo del rectángulo: ";
        std::cin >> Largo;
        std::cout << "Ingresa el valor del ancho del rectángulo: ";
        std::cin >> Ancho;
        int Area = Largo * Ancho;
        std::cout << "El área del rectángulo es: " << Area << std::endl;
    } else if (Figura == "3") {
        int Lado;
        std::cout << "Ingresa el valor del lado del cuadrado: ";
        std::cin >> Lado;
        int Area = Lado * Lado;
        std::cout << "El área del cuadrado es: " << Area << std::endl;
    } else if (Figura == "4") {
        int Bmayor, Bmenor, H;
        std::cout << "Ingresa el valor de la base mayor: ";
        std::cin >> Bmayor;
        std::cout << "Ingresa el valor de la base menor: ";
        std::cin >> Bmenor;
        std::cout << "Ingresa la altura del trapecio: ";
        std::cin >> H;
        double Area = ((Bmayor + Bmenor) * H) / 2;
        std::cout << "El área del trapecio es: " << Area << std::endl;
    } else {
        std::cout << "Opción no válida." << std::endl;
    }

    return 0;
}