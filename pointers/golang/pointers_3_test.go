package pointers

import "testing"

func TestDoubleSize(t *testing.T) {

	rectangle := &Rectangle{}

	rectangle.Width = 20
	rectangle.Height = 21

	doubleSize(rectangle)

	if rectangle.Width != 40 || rectangle.Height != 42 {
		t.Errorf("Error test")
	}

	t.Log(rectangle)
}
