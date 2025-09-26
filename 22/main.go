package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений.

Для очень больших чисел можно использовать math/big.
*/

func main() {
	a := big.NewInt(123456789)
	b := big.NewInt(987654321)
	fmt.Println("a + b =", big.NewInt(0).Add(a, b))
	fmt.Println("b - a =", big.NewInt(0).Sub(b, a))
	fmt.Println("a * b =", big.NewInt(0).Mul(a, b))
	fmt.Println("b / a =", big.NewInt(0).Div(b, a))
}
