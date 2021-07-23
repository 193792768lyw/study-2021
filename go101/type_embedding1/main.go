package main

import (
	"fmt"
	"reflect"
)

//
//type x string
//
//func (x) M() {}
//
//type A struct {
//	x
//}
//
//func (A) y(int) bool {
//	return false
//}
//
//type y struct {
//	z byte
//}
//
//type B struct {
//	y
//}
//
//func (B) x(string) {}
//
//func mainc() {
//	var v struct {
//		A
//		B
//	}
//	//_ = v.x // error: 模棱两可的v.x
//	//_ = v.y // error: 模棱两可的v.y
//	_ = v.M // ok. <=> v.A.x.M
//	_ = v.z // ok. <=> v.B.y.z
//
//}

//
//type A struct {
//	x string
//}
//func (A) y(int) bool {
//	return false
//}
//
//type B struct {
//	y bool
//}
//func (B) x(string) {}
//
//type C struct {
//	B
//}
//
//var v1 struct {
//	A
//	B
//}
//
//func f1() {
//	_ = v1.x // error: 模棱两可的v1.x
//	_ = v1.y // error: 模棱两可的v1.y
//}
//
//var v2 struct {
//	A
//	C
//}
//
//func f2() {
//	v2.
//	fmt.Printf("%T \n", v2.x) // string
//	fmt.Printf("%T \n", v2.y) // func(int) bool
//}
//
//
//type F func(int) bool
//func (f F) Validate(n int) bool {
//	return f(n)
//}
//func (f *F) Modify(f2 F) {
//	*f = f2
//}
//
//type B bool
//func (b B) IsTrue() bool {
//	return bool(b)
//}
//func (pb *B) Invert() {
//	*pb = !*pb
//}
//
//type I interface {
//	Load()
//	Save()
//}
//
//func PrintTypeMethods(t reflect.Type) {
//	fmt.Println(t, "has", t.NumMethod(), "methods:")
//	for i := 0; i < t.NumMethod(); i++ {
//		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
//	}
//}
//
//func main() {
//	var s struct {
//		F
//		*B
//		I
//	}
//
//	PrintTypeMethods(reflect.TypeOf(s))
//	fmt.Println()
//	PrintTypeMethods(reflect.TypeOf(&s))
//}

type MyInt int

func (mi MyInt) IsOdd() bool {
	return mi%2 == 1
}

type Age MyInt

type X struct {
	MyInt
	Oo
}

func (x X) Double() MyInt {
	return x.MyInt + x.MyInt
}

type Y struct {
	Age
}

type Z X

type Oo struct {
	Name33
}

type Name33 struct {
}

func (Name33) Papp() {

}

func main() {

	var d Z
	t := reflect.TypeOf(d) // the Singer type
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&d) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}

}

type Ia interface {
	fa() int
}

type Ib = interface {
	fb()
}

type Ic interface {
	fa() bool
	fb()
}
