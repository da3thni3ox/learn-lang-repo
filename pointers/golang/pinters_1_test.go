package pointers

import "testing"

func TestSum(t *testing.T) {

	a, b := 10, 15

	sum := sum(&a, &b)

	if sum != 25 {
		t.Errorf("Сумма чисел %d + %d не верна, сумма равна %d ожидается 25", a, b, sum)
	}
}
