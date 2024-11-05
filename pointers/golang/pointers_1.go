package pointers

func sum(a, b *int) int {
	sum := *a + *b

	return sum
}
