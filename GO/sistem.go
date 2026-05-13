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
n :=
matrix := make([][]Fraction, n)
for i := 0; i < n; i++ {
matrix[i] = make([]Fraction, n+1)
}
func Gaus(matrix [][]Fraction) Fraction {
	for i := 0; i < len(matrix); i++ {
		Pivot := Fraction{
			n: matrix[i][i].n,
			d: matrix[i][i].d,
		}

		if Pivot.n == 0 {//если опорный 0 то в цикле ищем как его опустить или получаем ошибку
			for j := i+1; j < len(matrix); j++ {
				if matrix[j][i].n != 0 {
					matrix[i], matrix[j] = matrix[j], matrix[i]
				} else if matrix[i][j].n == 0 && matrix[i][len(matrix[i])-1].n == 0 {
					fmt.Print('No solution')
				} else {
					fmt.Print('infinity')
				}
			}
		}//теперь можем начинать вычисления по идее
		for q:=i+1; q < len(matrix); q++ {
			var K Fraction = Pivot.Div(matrix[q][i])
			K = Normalize(K)
			for r:=i; r < len(matrix[q]); r++ {
				matrix[q][r] = (K.MP(matrix[q][r])).Sub(matrix[i][r])
				matrix[q][r] = Normalize(matrix[q][r])

			}
		}



	}
}