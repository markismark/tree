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

func casToNode(v reflect.Value, deep int) Node {
	t := v.Kind().String()
	n := Node{FiledType: t, Deep: deep}
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
		n.FieldValue = fmt.Sprintf("%x", v.Pointer())
	case "struct":
		//n.FiledType = reflect.TypeOf(ob).String()
	case "complex64", "complex128":
		n.FieldValue = fmt.Sprintf("%v", v)
	case "map":
		if !v.IsNil() {
			keys := v.MapKeys()
			n.Children = make([]Node, len(keys))
			for i, key := range keys {
				kv := v.MapIndex(key)
				kn := casToNode(kv, deep+1)
				kn.FieldName = keyToString(key)
				n.Children[i] = kn
				i++
			}
		}
	}
	return n
}

func keyToString(key reflect.Value) string {
	t := key.Kind().String()
	switch t {
	case "bool":
		b := key.Bool()
		if b {
			return "true"
		} else {
			return "false"
		}
	case "string":
		return key.String()
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64":
		return fmt.Sprintf("%d", key)
	case "float32", "float64":
		return fmt.Sprintf("%f", key)
	case "ptr":
		return fmt.Sprintf("%x", key.Pointer())
	default:
		return fmt.Sprintf("%x", key.Pointer())
	}
}
