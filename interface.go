package goutils

type Mappable[K comparable] interface {
	Key() K
}
