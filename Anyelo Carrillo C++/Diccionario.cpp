#include <iostream>
#include <map>

int main() {
    std::map<std::string, std::string> coche = {
        {"marca", "Ford"},
        {"color", "rojo"},
        {"modelo", "sedan"},
        {"placa", "ROS227"}
    };

    std::cout << coche["color"] << std::endl;
    coche["color"] = "verde";
    std::cout << coche["color"] << std::endl;

    std::cout << coche["marca"] << std::endl;
    coche["marca"] = "Renault";
    std::cout << coche["marca"] << std::endl;

    for (const auto& par : coche) {
        std::cout << par.first << ": " << par.second << std::endl;
    }

    return 0;
}