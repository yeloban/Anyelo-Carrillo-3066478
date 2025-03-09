#include <iostream>
#include <string>

int main() {
    int a = 2;

    if (a == 2) {
        std::cout << "a vale 2" << std::endl;
    } else {
        std::cout << "a es diferente de 2" << std::endl;
    }

    std::string Nombre = "Esteban";
    int Edad = 28;
    std::string Pais = "Colombia";

    if (Nombre == "Esteban" && Edad == 28 && Pais == "Colombia") {
        std::cout << "Su nombre es " << Nombre << ", tiene " << Edad << " años y es de " << Pais << std::endl;
    } else if (Nombre == "Esteban" && !(Edad == 28)) {
        std::cout << "Su nombre es Esteban y no tiene 28 años" << std::endl;
    } else if (Nombre != "Esteban" && Edad == 28) {
        std::cout << "Su nombre no es Esteban y tiene 28 años" << std::endl;
    } else {
        std::cout << "No se llama Esteban ni tiene 28 años" << std::endl;
    }

    return 0;
}
         