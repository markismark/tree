#Print object when debugging

## How to use


```shell
go get github.com/Maxgis/tree
```
```
import (github.com/Maxgis/tree)
...
type People struct {
    Name string
	Age  int
	sex  string

	Father   *People
	Children []*People
	Pro      map[string]string
}

func Test_print(t *testing.T) {
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
	fmt.Printf("%#v\n", fa)
}


```

##Print Example

```shell
{
  Name:string("jim"),
  Age:int(44),
  sex:string("man"),
  Father:*tree.People(nil),
  Children:[
    {
      Name:string("jimson"),
      Age:int(12),
      sex:string("boy"),
      Father:*tree.People(c8200980f0,reference to each other),
      Children:[],
      Pro:{}(map[string]string)
    }(*tree.People),
    {
      Name:string("jimdaughter"),
      Age:int(14),
      sex:string("girl"),
      Father:*tree.People(nil),
      Children:[],
      Pro:{}(map[string]string)
    }(*tree.People)
  ],
  Pro:{
    city:string("Beijing")
  }(map[string]string)
}(*tree.People)
```

## Original Print

```shell
&tree.People{Name:"jim", Age:44, sex:"man", Father:(*tree.People)(nil), Children:[]*tree.People{(*tree.People)(0xc820014280), (*tree.People)(0xc8200142d0)}, Pro:map[string]string{"city":"Beijing"}}
```