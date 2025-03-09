#include <iostream>

int main() {
    int A, B;
    std::cout << "Ingrese valor: ";
    std::cin >> A;
    std::cout << "Ingrese valor: ";
    std::cin >> B;

    int suma = A + B;
    std::cout << "La suma de los números es: " << suma << std::endl;

    int res = A - B;
    std::cout << "La resta de los números es: " << res << std::endl;

    int multi = A * B;
    std::cout << "La multiplicación de los números es: " << multi << std::endl;

    double div = static_cast<double>(A) / B;
    std::cout << "La división de los números es: " << div << std::endl;

    bool igual = (A == B);
    std::cout << "El número es igual: " << std::boolalpha << igual << std::endl;

    bool mayor = (A > B);
    std::cout << "El número mayor es: " << (mayor ? A : B) << std::endl;

    return 0;
}

