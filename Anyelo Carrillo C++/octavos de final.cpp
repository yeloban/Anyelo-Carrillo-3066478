#include <iostream>
#include <vector>
#include <string>
#include <cstdlib> // Para rand()
#include <ctime>   // Para time()

std::string jugar_partido(const std::string& equipo1, const std::string& equipo2) {
    std::cout << "\n" << equipo1 << " vs " << equipo2 << std::endl;
    int goles1, goles2;
    std::cout << "Ingrese goles de " << equipo1 << ": ";
    std::cin >> goles1;
    std::cout << "Ingrese goles de " << equipo2 << ": ";
    std::cin >> goles2;

    if (goles1 > goles2) {
        std::cout << equipo1 << " gana y avanza a la siguiente ronda!" << std::endl;
        return equipo1;
    } else if (goles2 > goles1) {
        std::cout << equipo2 << " gana y avanza a la siguiente ronda!" << std::endl;
        return equipo2;
    } else {
        std::cout << "EMPATE! Se jugará tiempo extra..." << std::endl;
        return jugar_tiempo_extra(equipo1, equipo2);
    }
}

std::string jugar_tiempo_extra(const std::string& equipo1, const std::string& equipo2) {
    int goles1, goles2;
    std::cout << "Tiempo extra - Goles de " << equipo1 << ": ";
    std::cin >> goles1;
    std::cout << "Tiempo extra - Goles de " << equipo2 << ": ";
    std::cin >> goles2;

    if (goles1 > goles2) {
        std::cout << equipo1 << " gana en tiempo extra!" << std::endl;
        return equipo1;
    } else if (goles2 > goles1) {
        std::cout << equipo2 << " gana en tiempo extra!" << std::endl;
        return equipo2;
    } else {
        std::cout << "¡Sigue el empate! Se define por penales..." << std::endl;
        return jugar_penales(equipo1, equipo2);
    }
}

std::string jugar_penales(const std::string& equipo1, const std::string& equipo2) {
    std::srand(std::time(0));
    while (true) {
        int penales1 = std::rand() % 5 + 1;
        int penales2 = std::rand() % 5 + 1;
        std::cout << "Penales: " << equipo1 << " " << penales1 << " - " << penales2 << " " << equipo2 << std::endl;

        if (penales1 > penales2) {
            std::cout << equipo1 << " gana por penales!" << std::endl;
            return equipo1;
        } else if (penales2 > penales1) {
            std::cout << equipo2 << " gana por penales!" << std::endl;
            return equipo2;
        } else {
            std::cout << "Sigue el empate en penales... ¡Se repiten los tiros!" << std::endl;
        }
    }
}

std::vector<std::string> simular_ronda(const std::vector<std::string>& equipos, const std::string& nombre_ronda) {
    std::cout << "\n===== " << nombre_ronda << " =====" << std::endl;
    std::vector<std::string> ganadores;

    for (size_t i = 0; i < equipos.size(); i += 2) {
        std::string equipo1 = equipos[i];
        std::string equipo2 = equipos[i + 1];
        std::string ganador = jugar_partido(equipo1, equipo2);
        ganadores.push_back(ganador);
    }

    return ganadores;
}

int main() {
    std::vector<std::string> equipos = {
        "Brasil", "Corea del Sur", "Francia", "Dinamarca",
        "Argentina", "México", "España", "Alemania",
        "Japón", "Inglaterra", "Perú", "USA",
        "Nigeria", "Chile", "Portugal", "Colombia"
    };

    std::vector<std::string> rondas = {"Octavos de Final", "Cuartos de Final", "Semifinal", "Final"};
    std::vector<std::string> ronda_actual = equipos;

    for (const auto& nombre_ronda : rondas) {
        ronda_actual = simular_ronda(ronda_actual, nombre_ronda);
    }

    std::cout << "\nEL CAMPEÓN DEL MUNDIAL ES: " << ronda_actual[0] << std::endl;

    return 0;
}