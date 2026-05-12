#include <iostream>
#include <vector>

template <typename T>
class LazyMatrix {
private:
    std::vector<std::vector<T>> data;

    void grow_to(size_t rows, size_t cols) {
        if (rows > m()) {
            data.resize(rows, std::vector<T>(n()));
        }
        if (cols > n()) {
            for (auto& row : data) {
                row.resize(cols);
            }
        }
    }

public:
    LazyMatrix() = default;

   
    size_t m() const { return data.size(); }
    size_t n() const { return data.empty() ? 0 : data[0].size(); }

   
    T& operator()(size_t i, size_t j) {
        grow_to(i + 1, j + 1);
        return data[i][j];
    }

   
    T operator()(size_t i, size_t j) const {
        if (i >= m() || j >= n()) return T();
        return data[i][j];
    }

    
    LazyMatrix<T> operator!() const {
        LazyMatrix<T> res;
        size_t current_m = m();
        size_t current_n = n();

        if (current_m > 0 && current_n > 0) {
            res.grow_to(current_n, current_m);
            for (size_t i = 0; i < current_m; ++i) {
                for (size_t j = 0; j < current_n; ++j) {
                    res.data[j][i] = data[i][j];
                }
            }
        }
        return res;
    }

    bool operator==(const LazyMatrix<T>& other) const {
        if (m() != other.m() || n() != other.n()) return false;
        return data == other.data;
    }

    bool operator!=(const LazyMatrix<T>& other) const {
        return !(*this == other);
    }
};

int main() {
    LazyMatrix<int> matrix;
    matrix(1, 1) = 5;
    matrix(0, 2) = 10;

    std::cout << "Original: " << matrix.m() << "x" << matrix.n() << std::endl;

    LazyMatrix<int> transposed = !matrix;
    
    std::cout << "Transposed: " << transposed.m() << "x" << transposed.n() << std::endl;
    std::cout << "Value at (2,0): " << transposed(2, 0) << " (must be 10)" << std::endl;

    return 0;
}

