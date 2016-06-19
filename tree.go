package tree

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	deepStr  = "  "
	commaStr = ","
)

type Node struct {
	Children      []*Node
	FiledType     string
	FiledRealType string
	FieldName     string
	FieldValue    string
	Deep          int
	IsNil         bool
}

type pp struct {
	dataNode *Node
	ptrs     []uintptr
}

func Print(ob interface{}) {
	p := &pp{}
	p.ptrs = make([]uintptr, 0)
	n := p.casToNode(reflect.ValueOf(ob), 1)
	p.dataNode = n
	p.print()
	fmt.Println("")
}

func (this *pp) casToNode(v reflect.Value, deep int) *Node {
	t := v.Kind().String()
	n := &Node{FiledType: t, Deep: deep, FiledRealType: t}
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
		vptr := v.Pointer()
		if vptr == 0 {
			n.FiledRealType = v.Type().String()
			n.FieldValue = "nil"
		} else if ptrInArray(vptr, this.ptrs) {
			n.FieldValue = fmt.Sprintf("%x,reference to each other", vptr)
			n.FiledRealType = v.Type().String()
		} else {
			this.ptrs = append(this.ptrs, vptr)
			n.FiledRealType = v.Type().String()
			pn := this.casToNode(v.Elem(), deep)
			n.Children = pn.Children
			n.FieldValue = pn.FieldValue
			n.FiledType = pn.FiledType
		}

	case "complex64", "complex128":
		n.FieldValue = fmt.Sprintf("%v", v)
	case "map":
		n.FiledRealType = v.Type().String()
		if !v.IsNil() {
			keys := v.MapKeys()
			n.Children = make([]*Node, len(keys))
			for i, key := range keys {
				kv := v.MapIndex(key)
				kn := this.casToNode(kv, deep+1)
				kn.FieldName = keyToString(key)
				n.Children[i] = kn
				i++
			}
		} else {
			n.IsNil = true
		}
	case "struct":
		n.FiledRealType = v.Type().String()
		n.Children = make([]*Node, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			kn := this.casToNode(f, deep+1)
			kn.FieldName = v.Type().Field(i).Name
			n.Children[i] = kn
		}
	case "array", "slice":
		length := v.Len()
		n.Children = make([]*Node, length)
		for i := 0; i < length; i++ {
			kn := this.casToNode(v.Index(i), deep+1)
			n.Children[i] = kn
		}
	}
	return n
}
func (this *pp) print() {
	this.printNode(this.dataNode)
}

func (this *pp) printNode(node *Node) {

	switch node.FiledType {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128", "bool":
		fmt.Printf("%s(%s)", node.FiledType, node.FieldValue)
	case "string":
		fmt.Printf("%s(\"%s\")", node.FiledType, node.FieldValue)
	case "ptr":
		fmt.Printf("%s(%s)", node.FiledRealType, node.FieldValue)
	case "array", "slice":
		if len(node.Children) == 0 {
			fmt.Print("[]")
		} else {
			fmt.Print("[\n")
			length := len(node.Children)
			for i, cnode := range node.Children {
				fmt.Printf("%s", strings.Repeat(deepStr, node.Deep))
				this.printNode(cnode)
				if i <= length-2 {
					fmt.Printf(commaStr)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("%s]", strings.Repeat(deepStr, node.Deep-1))
		}
	case "struct", "map":

		length := len(node.Children)
		if length == 0 {
			fmt.Printf("{}(%s)", node.FiledRealType)
		} else {
			fmt.Print("{\n")
			for i, cnode := range node.Children {
				fmt.Printf("%s%s:", strings.Repeat(deepStr, node.Deep), cnode.FieldName)
				this.printNode(cnode)
				if i <= length-2 {
					fmt.Printf(commaStr)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("%s}(%s)", strings.Repeat(deepStr, node.Deep-1), node.FiledRealType)
		}
	default:
		fmt.Printf("unknown type")
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

func ptrInArray(ptr uintptr, arr []uintptr) bool {
	for _, nptr := range arr {
		if ptr == nptr {
			return true
		}
	}
	return false
}
