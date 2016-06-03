package tree

import "reflect"

type node struct {
	children  []node
	filedType string
	fieldName string
	isPrt     bool
	deep      int
}

func Print(ob interface{}) {
	t := reflect.TypeOf(ob).Kind().String()
}
