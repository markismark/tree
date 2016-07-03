package tree

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	deepStr      = "  "
	commaStr     = ","
	blue         = "\033[32;1m"
	red          = "\033[31;1m"
	yellow       = "\033[33;1m"
	green        = "\033[34;1m"
	defaultColor = "\033[0m"
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
	buffer   string
}

func Print(ob interface{}) {
	str := Sprint(ob)
	fmt.Println(str)
}

func Sprint(ob interface{}) string {
	p := &pp{}
	p.ptrs = make([]uintptr, 0)
	n := p.casToNode(reflect.ValueOf(ob), 1)
	p.dataNode = n
	p.sprint()
	return p.buffer
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
			n.FieldValue = fmt.Sprintf("0x%x,reference to each other", vptr)
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
func (this *pp) sprint() {
	this.printNode(this.dataNode)
}

func (this *pp) write(str string) {
	this.buffer = this.buffer + str
}

func (this *pp) printNode(node *Node) {

	switch node.FiledType {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint16", "uint8", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128", "bool":
		this.write(fmt.Sprintf("%s("+blue+"%s"+defaultColor+")", node.FiledType, node.FieldValue))
	case "string":
		this.write(fmt.Sprintf("%s(\""+yellow+"%s"+defaultColor+"\")", node.FiledType, node.FieldValue))
	case "ptr":
		this.write(fmt.Sprintf("%s(%s)", node.FiledRealType, node.FieldValue))
	case "array", "slice":
		if len(node.Children) == 0 {
			this.write(fmt.Sprintf("[]"))
		} else {
			this.write(fmt.Sprintf("[\n"))
			length := len(node.Children)
			for i, cnode := range node.Children {
				this.write(fmt.Sprintf("%s", strings.Repeat(deepStr, node.Deep)))
				this.printNode(cnode)
				if i <= length-2 {
					this.write(fmt.Sprintf(commaStr))
				}
				this.write(fmt.Sprintf("\n"))
			}
			this.write(fmt.Sprintf("%s]", strings.Repeat(deepStr, node.Deep-1)))
		}
	case "struct", "map":

		length := len(node.Children)
		if length == 0 {
			this.write(fmt.Sprintf("{}(%s)", node.FiledRealType))
		} else {
			this.write(fmt.Sprintf("{\n"))
			for i, cnode := range node.Children {
				this.write(fmt.Sprintf("%s"+green+"%s"+defaultColor+":", strings.Repeat(deepStr, node.Deep), cnode.FieldName))
				this.printNode(cnode)
				if i <= length-2 {
					this.write(fmt.Sprintf(commaStr))
				}
				this.write(fmt.Sprintf("\n"))
			}
			this.write(fmt.Sprintf("%s}(%s)", strings.Repeat(deepStr, node.Deep-1), node.FiledRealType))
		}
	default:
		this.write(fmt.Sprintf("unknown type"))
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
		return fmt.Sprintf("Ox%x", key.Pointer())
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
