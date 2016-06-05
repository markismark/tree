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
	n := casToNode(ob, 1)
	fmt.Printf("%#v\n", n)
}

func casToNode(ob interface{}, deep int) *Node {
	t := reflect.TypeOf(ob).Kind().String()
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(ob))
	n := &Node{FiledType: t, Deep: deep}
	switch t {
	case "bool":
		b := ob.(bool)
		if b {
			n.FieldValue = "true"
		} else {
			n.FieldValue = "false"
		}
	case "string":
		n.FieldValue = ob.(string)
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64":
		n.FieldValue = fmt.Sprintf("%d", ob)
	case "float32", "float64":
		n.FieldValue = fmt.Sprintf("%f", ob)
	case "ptr":
		n.FiledType = reflect.TypeOf(ob).String()
		n.FieldValue = fmt.Sprintf("%p", ob)
	case "struct":
		n.FiledType = reflect.TypeOf(ob).String()
	}
	return n
}
