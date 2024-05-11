package goutils

// IF 函数 if flag return a else b
func IF[T any](flag bool, a T, b T) T {
	if flag {
		return a
	}
	return b
}
