package utils

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func IsIn(main string, other ...string) bool {
	for _, str := range other {
		if str == main {
			return true
		}
	}
	return false
}
