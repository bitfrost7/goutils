package goutils

import (
	"fmt"
	"testing"
)

type Student struct {
	id   int
	name string
}

func (t *Student) Key() int {
	return t.id
}

func (t *Student) Conflict(other *Student) *Student {
	t.name = t.name + other.name
	return t
}

func TestSlice2Map(t *testing.T) {
	s := make([]*Student, 0)
	s = append(s, &Student{
		id:   1,
		name: "a",
	})
	s = append(s, &Student{
		id:   1,
		name: "b",
	})
	s = append(s, &Student{
		id:   2,
		name: "ff",
	})
	s = append(s, &Student{
		id:   3,
		name: "dd",
	})
	m := Slice2MapWithCombine(s)
	fmt.Println(m)
}
