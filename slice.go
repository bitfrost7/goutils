package goutils

type joinable[K comparable] interface {
	key() K
}

func SliceJoin[K comparable, A joinable[K], B joinable[K]](a []A, b []B, join func(A, B) A) []A {
	m := make(map[K]A)
	for _, v := range a {
		m[v.key()] = v
	}
	for _, v := range b {
		if _, ok := m[v.key()]; ok {
			join(m[v.key()], v)
		}
	}
	return a
}
