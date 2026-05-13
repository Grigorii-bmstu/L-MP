package main

import "fmt"

type Fraction struct {
	n, d int
}

func (f1 Fraction) MP(f2 Fraction) Fraction {
	f1 = Normalize(f1)
	f2 = Normalize(f2)

	return Normalize(Fraction{
		n: f1.n * f2.n,
		d: f1.d * f2.d,
	})
} //Умножение

func MPMany(fractions ...Fraction) Fraction {
	res := Fraction{}
	res.n, res.d = 1, 1

	for _, fraction := range fractions {
		res.n = res.n * fraction.n
		res.d = res.d * fraction.d
		res = Normalize(res)
	}
	return res
} //Умножение многих

// region Вспомогательные функции (abs / gcd)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// endregion

func Normalize(fraction Fraction) Fraction {
	if fraction.d == 0 {
		return fraction
	}
	if fraction.d < 0 {
		fraction.d = -fraction.d
		fraction.n = -fraction.n
	}
	c := gcd(abs(fraction.d), abs(fraction.n))
	fraction.n /= c
	fraction.d /= c
	return fraction
} //Нормализация

func (f1 Fraction) Div(f2 Fraction) Fraction {
	res := Fraction{
		n: f1.n * f2.d,
		d: f1.d * f2.n,
	}
	return Normalize(res)
} //Деление

func (f1 Fraction) Sub(f2 Fraction) Fraction {
	res := Fraction{
		n: f1.n*f2.d - f1.d*f2.n,
		d: f1.d * f2.d,
	}
	return Normalize(res)
} //Вычитание

func (f1 Fraction) Sum(f2 Fraction) Fraction {
	res := Fraction{
		n: f1.n*f2.d + f1.d*f2.n,
		d: f1.d * f2.d,
	}
	return Normalize(res)
} //Сложение

func Gaus(matrix [][]Fraction) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		Pivot := matrix[i][i]

		if Pivot.n == 0 { //если опорный 0 то в цикле ищем как его опустить или получаем ошибку
			found := false
			for j := i + 1; j < n; j++ {
				if matrix[j][i].n != 0 {
					matrix[i], matrix[j] = matrix[j], matrix[i]
					Pivot = matrix[i][i]
					found = true
					break
				}
			}
			if !found {
				fmt.Print("No solution")
				return
			}
		} //теперь можем начинать вычисления по идее

		for q := i + 1; q < n; q++ {
			var K Fraction = matrix[q][i].Div(Pivot)
			for r := i; r <= n; r++ {
				matrix[q][r] = matrix[q][r].Sub(K.MP(matrix[i][r]))
				//привели матрицу в треугольный вид
			}
		}
	}

	x := make([]Fraction, n)
	for i := n - 1; i >= 0; i-- {
		if matrix[i][i].n == 0 {
			fmt.Print("No solution")
			return
		}
		res := matrix[i][n]
		//будем считать сумму известных элементов
		for j := i + 1; j < n; j++ {
			res = res.Sub(matrix[i][j].MP(x[j]))
		}
		x[i] = res.Div(matrix[i][i])
	}

	for _, val := range x {
		fmt.Printf("%d/%d\n", val.n, val.d)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	matrix := make([][]Fraction, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]Fraction, n+1)
		for j := 0; j <= n; j++ {
			var val int
			fmt.Scan(&val)
			matrix[i][j] = Fraction{val, 1}
		}
	}

	Gaus(matrix)
}
