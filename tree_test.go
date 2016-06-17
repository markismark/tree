package tree

import "testing"

type People struct {
	Name string
	Age  int
	sex  string

	Father   *People
	Children []*People
}

func Test_print(t *testing.T) {
	// m := make(map[string]string)
	// m["name"] = "jim"
	// m["sex"] = "man"
	// Print(m)
	st := &People{Name: "jim", Age: 44, sex: "man"}
	son := &People{Name: "jimson", Age: 12, sex: "boy"}
	daughter := &People{Name: "jimdaughter", Age: 14, sex: "girl"}
	st.Children = make([]*People, 2)
	st.Children[0] = son
	st.Children[1] = daughter
	son.Father = st
	Print(st)
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
