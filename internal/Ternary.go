package internal

func Ternary[T any](condition bool, If, Else T) T {
	if condition {
		return If
	}

	return Else
}
