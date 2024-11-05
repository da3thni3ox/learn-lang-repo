package pointers

type Rectangle struct {
	Width, Height int
}

func doubleSize(r *Rectangle) {
	r.Width *= 2
	r.Height *= 2
}
