package tree

import "testing"

type People struct {
	Name string
	Age  int
	sex  string

	Father   *People
	Children []*People
	Pro      map[string]string
}

func Test_print(t *testing.T) {
	// m := make(map[string]string)
	// m["name"] = "jim"
	// m["sex"] = "man"
	// Print(m)
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
	son.Children = nil
	Print(son.Children)
	// stp := People{}
	// stp = nil
	// Print(stp)
	//fmt.Printf("%#v", &st)
	// ar := [5]int{1, 3, 7, 9, 11}
	// Print(ar)
	// sp := []string{"啊哈", "哈哈"}
	// Print(sp)

	// arr := make([][]int, 5, 5)
	// for i := 0; i < 5; i++ {
	// 	m2 := make([]int, 5, 5) //可用循环对m2赋值，默认建立初值为0
	// 	arr[i] = m2             //建立第二维
	// }
	// Print(arr)

	// star := make([]Mystruct, 2)
	// star[0] = Mystruct{Name: "jim", Age: 12, sex: "boy"}
	// star[1] = Mystruct{Name: "Lucy", Age: 13, sex: "girl"}
	// Print(star)
}
