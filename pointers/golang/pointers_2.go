package pointers

func swap(a, b *int) {
	*a = *b
	*b = *a
}
