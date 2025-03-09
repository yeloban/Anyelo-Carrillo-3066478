#include <iostream>
#include <cmath> // Para usar pow()

int main() {
    int A, C;
    std::cout << "Ingrese el primer valor: ";
    std::cin >> A;

    std::cout << "Ingrese el segundo valor: ";
    std::cin >> C;

    int valor = pow(A, C); // Calcula A elevado a C
    std::cout << "La potencia de " << A << " sobre " << C << " es: " << valor << std::endl;

    return 0;
}
