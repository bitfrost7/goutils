package goutils

// SliceFilter 过滤切片中的一些特定值并返回结果。
//
//	s: 	指定切片。
//	filter: 	过滤函数,true则过滤。
func SliceFilter[T any](s []T, filter func(v T) bool) []T {
	res := make([]T, 0, len(s))
	for _, e := range s {
		if !filter(e) {
			res = append(res, e)
		}
	}
	return res
}

// Slice2Map 将实现Mappable接口的slice转化为map
//
// 参数:
//
//	a: 	指定切片。
//	key: 	返回切片类型中的索引。
//
// 返回值:
//
//	返回切片转化成map结构的结果。
func Slice2Map[K comparable, T Mappable[K]](s []T) map[K]T {
	m := make(map[K]T)
	for _, v := range s {
		m[v.Key()] = v
	}
	return m
}

// Slice2MapWithConflict 将实现Mappable接口的slice转化为map 并指定冲突解决方式
//
// 参数:
//
//	  a: 	指定切片。
//	  key: 	返回切片类型中的索引。
//		 conflict:	当map中出现冲突时进行处理。
//
// 返回值:
//
//	返回切片转化成map结构的结果。
func Slice2MapWithConflict[K comparable, T Mappable[K]](s []T, conflict func(old, new T) T) map[K]T {
	m := make(map[K]T)
	for _, v := range s {
		e, ok := m[v.Key()]
		m[v.Key()] = IF(ok, conflict(e, v), e)
	}
	return m
}

func SliceReduce[U any, V any](initial U, s1 []V, reduce func(U, V) U) U {
	result := initial
	for _, v := range s1 {
		result = reduce(result, v)
	}
	return result
}
