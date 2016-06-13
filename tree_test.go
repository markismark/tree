package tree

import "testing"

type Mystruct struct {
	Name string
	Age  int
	sex  string
}

func Test_print(t *testing.T) {
	// m := make(map[string]string)
	// m["name"] = "jim"
	// m["sex"] = "man"
	// Print(m)
	// st := Mystruct{Name: "jim", Age: 12, sex: "man"}
	// Print(st)
	ar := [5]int{1, 3, 7, 9, 11}
	Print(ar)
	sp := []string{"啊哈", "哈哈"}
	Print(sp)
}
