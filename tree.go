package tree

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	deepStr  = "    "
	commaStr = ","
)

type Node struct {
	Children      []Node
	FiledType     string
	FiledRealType string
	FieldName     string
	FieldValue    string
	Deep          int
	IsNil         bool
}

func Print(ob interface{}) {
	n := casToNode(reflect.ValueOf(ob), 1)
	print(n)
	fmt.Println("")
}

func casToNode(v reflect.Value, deep int) Node {
	t := v.Kind().String()
	n := Node{FiledType: t, Deep: deep, FiledRealType: t}
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
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64", "uintptr":
		n.FieldValue = fmt.Sprintf("%d", v)
	case "float32", "float64":
		n.FieldValue = fmt.Sprintf("%f", v)
	case "ptr":
		n.FiledType = "*" + v.Elem().Kind().String()
		n.FieldValue = fmt.Sprintf("%x", v.Pointer())
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
		} else {
			n.IsNil = true
		}
	case "struct":
		n.FiledRealType = v.Type().String()
		n.Children = make([]Node, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			kn := casToNode(f, deep+1)
			kn.FieldName = v.Type().Field(i).Name
			n.Children[i] = kn
		}
	case "array", "slice":
		length := v.Len()
		n.Children = make([]Node, length)
		for i := 0; i < length; i++ {
			kn := casToNode(v.Index(i), deep+1)
			n.Children[i] = kn
		}
	}
	return n
}

func print(node Node) {

	switch node.FiledType {
	case "string", "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128":
		fmt.Printf("%s(%s)", node.FiledType, node.FieldValue)
	case "array", "slice":
		if len(node.Children) == 0 {
			fmt.Print("[]\n")
		} else {
			fmt.Print("[\n")
			length := len(node.Children)
			for i, cnode := range node.Children {
				fmt.Printf("%s", strings.Repeat(deepStr, node.Deep))
				print(cnode)
				if i <= length-2 {
					fmt.Printf(commaStr)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("%s]", strings.Repeat(deepStr, node.Deep-1))
		}
	case "struct":
		fmt.Print("{\n")
		length := len(node.Children)
		for i, cnode := range node.Children {
			fmt.Printf("%s%s:", strings.Repeat(deepStr, node.Deep), cnode.FieldName)
			print(cnode)
			if i <= length-2 {
				fmt.Printf(commaStr)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("%s}(%s)", strings.Repeat(deepStr, node.Deep-1), node.FiledRealType)
	}
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
