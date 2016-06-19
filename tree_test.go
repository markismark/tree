package tree

import (
	"fmt"
	"testing"
)

type People struct {
	Name string
	Age  int
	sex  string

	Father   *People
	Children []*People
	Pro      map[string]string
}

func Test_print(t *testing.T) {
	fa := &People{Name: "jim", Age: 44, sex: "man"}
	fa.Pro = make(map[string]string)
	fa.Pro["city"] = "Beijing"
	son := &People{Name: "jimson", Age: 12, sex: "boy"}
	daughter := &People{Name: "jimdaughter", Age: 14, sex: "girl"}
	fa.Children = make([]*People, 2)
	fa.Children[0] = son
	fa.Children[1] = daughter
	son.Father = fa
	Print(fa)
	fmt.Printf("%#v\n", fa)
}
