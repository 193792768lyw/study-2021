package main

import (
	"Study/go101/foo"
	"fmt"
)

type B struct {
	n bool
}

func (b B) M() {
	fmt.Println("B", b.n)
}

type C struct {
	foo.A
	//B
}

func main() {
	var c C
	//c.m()      // B false
	foo.Bar(c) // A 0

	//t := reflect.TypeOf(C{}) // the Singer type
	//fmt.Println(t, "has", t.NumField(), "fields:")
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	//}
	//fmt.Println(t, "has", t.NumMethod(), "methods:")
	//for i := 0; i < t.NumMethod(); i++ {
	//	fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	//}
	//
	//pt := reflect.TypeOf(&C{}) // the *Singer type
	//fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	//for i := 0; i < pt.NumMethod(); i++ {
	//	fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	//}
}
