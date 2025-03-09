#include <iostream>
#include <string>

int main() {
    std::string a, b, c;
    int d;

    std::cout << "Escriba sus nombres completos: ";
    std::getline(std::cin, a);

    std::cout << "Escriba sus apellidos completos: ";
    std::getline(std::cin, b);

    std::cout << "Escriba su profesión: ";
    std::getline(std::cin, c);

    std::cout << "Escriba su año de nacimiento: ";
    std::cin >> d;

    int e = 2025 - d;
    std::cout << "El (La) " << c << " " << a << " " << b << " tiene " << e << " años" << std::endl;

    return 0;
}