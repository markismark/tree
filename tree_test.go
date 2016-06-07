package tree

import (
	"fmt"
	"testing"
)

func Test_print(t *testing.T) {
	b := true
	Print(b)
	var d64 int64 = 12
	Print(d64)
	var f64 float64 = 12.5
	Print(f64)
	Print(&f64)
	fmt.Println(f64)
	x := 101 + 45i
	Print(x)
}
