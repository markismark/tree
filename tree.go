package tree

import (
	"fmt"
	"reflect"
)

type Node struct {
	Children   []*Node
	FiledType  string
	FieldName  *Node
	FieldValue string
	Deep       int
}

func Print(ob interface{}) {
	n := casToNode(reflect.ValueOf(ob), 1)
	fmt.Printf("%#v\n", n.Children[0].FieldName)
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
		n.FieldName = casToNode(reflect.ValueOf(v.Pointer()), 1)
	case "struct":
		//n.FiledType = reflect.TypeOf(ob).String()
	case "complex64", "complex128":
		n.FieldValue = fmt.Sprintf("%v", v)
	case "map":
		if !v.IsNil() {
			keys := v.MapKeys()
			n.Children = make([]*Node, len(keys))
			for i, key := range keys {
				kv := v.MapIndex(key)
				kn := casToNode(kv, deep+1)
				kn.FieldName = casToNode(key, 1)
				n.Children[i] = kn
				i++
			}
		}
	}
	return n
}
