#include <iostream>

int main() {
    bool a = true, b = false;
    std::cout << std::boolalpha << (a && b) << std::endl;

    int c = 4, d = 10;
    std::cout << std::boolalpha << (a == b && c < d) << std::endl;
    std::cout << std::boolalpha << (!(a == b) && c > d) << std::endl;

    return 0;
}