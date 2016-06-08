package tree

import "testing"

func Test_print(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "jim"
	m["sex"] = "man"
	Print(m)
}
