int main() {
    std::vector<int> a = {1, 2, 3, 4, 5};
    std::vector<int> b = {6, 7, 8, 9, 10};
    std::vector<int> c;

    for (size_t i = 0; i < a.size(); i++) {
        c.push_back(a[i] * b[i]);
    }

    for (const auto& valor : c) {
        std::cout << valor << " ";
    }
    std::cout << std::endl;

    return 0;
}




