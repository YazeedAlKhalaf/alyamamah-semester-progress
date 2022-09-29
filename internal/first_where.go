package internal

func FirstWhere[T any](slice []T, filter func(*T) bool) (element *T) {
	for _, e := range slice {
		if filter(&e) {
			return &e
		}
	}

	return nil
}
