#include <iostream>

#include <iostream>

int main() {
    int a = 10, b = 4;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: *" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    int c = a * b;
    std::cout << "El resultado es: " << c << std::endl;
    std::cout << "Tipo de dato: " << typeid(c).name() << std::endl;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: /" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    double d = static_cast<double>(a) / b;
    std::cout << "El resultado es: " << d << std::endl;
    std::cout << "Tipo de dato: " << typeid(d).name() << std::endl;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: +" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    int e = a + b;
    std::cout << "El resultado es: " << e << std::endl;
    std::cout << "Tipo de dato: " << typeid(e).name() << std::endl;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: -" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    int f = a - b;
    std::cout << "El resultado es: " << f << std::endl;
    std::cout << "Tipo de dato: " << typeid(f).name() << std::endl;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: **" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    int g = pow(a, b);
    std::cout << "El resultado es: " << g << std::endl;
    std::cout << "Tipo de dato: " << typeid(g).name() << std::endl;

    std::cout << "La primera variable es: " << a << std::endl;
    std::cout << "El signo de la operación es: %" << std::endl;
    std::cout << "La segunda variable es: " << b << std::endl;
    int h = a % b;
    std::cout << "El resultado es: " << h << std::endl;
    std::cout << "Tipo de dato: " << typeid(h).name() << std::endl;

    return 0;
}