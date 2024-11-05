package pointers

import "testing"

func TestSwap(t *testing.T) {
	a, b := 4, 5

	swap(&a, &b)

	if a != b && b != a {
		t.Errorf("Ожидается A равна %d, а B равна %d", b, a)
	}
}
