package main

import (
	"encoding/json"
	"fmt"
	"goutils"
)

type test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *test) Key() int {
	return t.Id
}

func main() {
	t1 := new(test)
	t1.Id = 1
	t1.Name = "t1"
	t2 := new(test)
	t2.Id = 1
	t2.Name = "t2"
	s := []*test{t1, t2}
	combine := goutils.Slice2MapWithConflict(s, func(old, new *test) *test {
		old.Name = old.Name + new.Name
		return old
	})
	str, _ := json.Marshal(combine)
	fmt.Println(string(str))
}
