#include <iostream>
using namespace std;

int main() {
    // Primera parte: Bucle while para contar hasta 30
    int Contador = 0;
    while (Contador < 30) {
        Contador++;
        cout << "Iteración " << Contador << endl;
    }

    // Segunda parte: Bucle while con if-else y manejo de errores
    while (true) {
        int a;
        cout << "Introduzca un valor mayor a 10: ";
        cin >> a;

        if (cin.fail()) { // Verifica si la entrada no es un número
            cin.clear(); // Limpia el estado de error de cin
            cin.ignore(numeric_limits<streamsize>::max(), '\n'); // Ignora la entrada incorrecta
            cout << "Por favor, ingrese un número válido." << endl;
        } else {
            if (a > 10) {
                cout << "Es correcto" << endl;
                break; // Sale del bucle cuando la condición se cumple
            } else {
                cout << "Es incorrecto, pruebe de nuevo" << endl;
            }
        }
    }

    return 0;
}