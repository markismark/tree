package tree

import (
	"fmt"
	"reflect"
)

type Node struct {
	Children   []Node
	FiledType  string
	FieldName  string
	FieldValue string
	Deep       int
}

func Print(ob interface{}) {
	n := casToNode(reflect.ValueOf(ob), 1)
	fmt.Printf("%#v\n", n)
}

func casToNode(v reflect.Value, deep int) *Node {
	t := v.Kind().String()
	fmt.Println(t)
	n := &Node{FiledType: t, Deep: deep}
	switch t {
	case "bool":
		b := v.Bool()
		if b {
			n.FieldValue = "true"
		} else {
			n.FieldValue = "false"
		}
	case "string":
		n.FieldValue = v.String()
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64":
		n.FieldValue = fmt.Sprintf("%d", v)
	case "float32", "float64":
		n.FieldValue = fmt.Sprintf("%f", v)
	case "ptr":
		n.FiledType = "*" + v.Elem().Kind().String()
		n.FieldName = fmt.Sprintf("Ox%x", v.Pointer())
	case "struct":
		//n.FiledType = reflect.TypeOf(ob).String()
	case "complex64", "complex128":
		n.FieldValue = fmt.Sprintf("%v", v)
	case "map":
		// if !reflect.ValueOf(ob).IsNil() {
		// 	keys := reflect.ValueOf(ob).MapKeys()\
		//     for key := range keys{

		//     }
		// }
	}
	return n
}
