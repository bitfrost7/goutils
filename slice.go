package goutils

type keyable[K comparable] interface {
	key() K
}

func SliceJoin[K comparable, A keyable[K], B keyable[K]](a []A, b []B, join func(A, B) A) []A {
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

func Slice2Map[K comparable, T keyable[K]](s []T) map[K]T {
	m := make(map[K]T)
	for _, v := range s {
		m[v.key()] = v
	}
	return m
}

func SliceReduce[U any, V any](initial V, s1 []U, reduce func(V, U) V) []V {
	result := make([]V, len(s1)+1)
	result[0] = initial
	for i, v := range s1 {
		result[i+1] = reduce(result[i], v)
	}
	return result
}
