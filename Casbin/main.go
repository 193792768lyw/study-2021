package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type I interface {
	m(int) bool
}

type T string

func (t T) m(n int) bool {
	return len(t) > n
}

func main() {
	var i I = T("gopher")
	fmt.Println(i.m(5))                           // true
	fmt.Println(I.m(i, 5))                        // true
	fmt.Println(interface{ m(int) bool }.m(i, 5)) // true

	//// 下面这几行被执行的时候都将会产生一个恐慌。
	//I(nil).m(5)
	//I.m(nil, 5)
	//interface {m(int) bool}.m(nil, 5)
}

func mainddd() {
	words := []string{
		"Go", "is", "a", "high",
		"efficient", "language.",
	}

	// fmt.Println函数的原型为：
	// func Println(a ...interface{}) (n int, err error)
	// 所以words...不能传递给此函数的调用。

	// fmt.Println(words...) // 编译不通过

	// 将[]string值转换为类型[]interface{}。
	iw := make([]interface{}, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}
	fmt.Println(iw...) // 编译没问题
}

func mainff() {
	var a, b, c interface{} = "abc", 123, "a" + "b" + "c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。
}

type Ia interface {
	fa()
}

type Ib = interface {
	fb()
}

type Ic interface {
	fa()
	fb()
}

type Id = interface {
	Ia // 内嵌Ia
	Ib // 内嵌Ib
}

type Ie interface {
	Ia // 内嵌Ia
	fb()
}

type Ix interface {
	Ia
	Ic
}

type Iy = interface {
	Ib
	Ic
}

type Iz interface {
	Ic
	fa()
}

func mainp() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// 这里，虽然变量v只被声明了一次，但是它在
		// 不同分支中可以表现为多个类型的变量值。
		switch v := x.(type) {
		case []int: // 一个类型字面表示
			// 在此分支中，v的类型为[]int。
			fmt.Println("int slice:", v)
		case string: // 一个类型名
			// 在此分支中，v的类型为string。
			fmt.Println("string:", v)
		case int, float64, int32: // 多个类型名
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("number:", v)
		case nil:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println(v)
		default:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("others:", v)
		}
		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
	}
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func mainbb() {
	var x interface{} = DummyWriter{}

	// y的动态类型为内置类型string。
	var y interface{} = "abc"
	var w Writer

	var ok bool

	// DummyWriter既实现了Writer，也实现了interface{}。
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// y的动态类型为string。string类型并没有实现Writer。
	w, ok = y.(Writer)
	fmt.Println(w, ok) // <nil> false
	w = y.(Writer)     // 将产生一个恐慌
}

// 类型*Book实现了接口类型Aboutable。
type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

//func (book Book) About() string {
//	return "Book: " + book.name
//}

type Aboutable interface {
	About() string
}

func main333() {
	//var dd interface{} = map[string]interface{}{}
	//
	//fmt.Println(dd)
	//a := dd.(map[string]interface{})
	//a["fds0"] = 99
	//fmt.Println(dd)

	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123

	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// 情形三：
}
func main22() {
	var gg Aboutable
	if gg == nil {
		fmt.Println("我是nil的接口")
	}
	var ff interface{}
	if ff == nil {
		fmt.Println("我是nil的接口")
	}
	ff = gg
	if ff == nil {
		fmt.Println("我是nil的接口")
	}

}

//// 一个*Book值被包裹在了一个Aboutable值中。
//var a Aboutable = &Book{"Go语言101"}
//fmt.Println(a) // &{Go语言101}
//
//// i是一个空接口值。类型*Book实现了任何空接口类型。
//var i interface{} = &Book{"Rust 101"}
//fmt.Println(i) // &{Rust 101}
//
//// Aboutable实现了空接口类型interface{}。
//i = a
//fmt.Println(i) // &{Go语言101}
type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func main33() {
	var d interface{ About() string }
	var h int = 99
	var ff *MyInt = (*MyInt)(&h)
	d = ff
	fmt.Println(d)
}
func main3() {
	addone := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }
	double := func(x int) int { return x + x }

	transforms := map[string][]func(int) int{
		"inc,inc,inc": {addone, addone, addone},
		"sqr,inc,dbl": {square, addone, double},
		"dbl,sqr,sqr": {double, double, square},
	}

	for _, n := range []int{2, 3, 5, 7} {
		fmt.Println(">>>", n)
		for name, transfers := range transforms {
			result := n
			for _, xfer := range transfers {
				result = xfer(result)
			}
			fmt.Printf(" %v: %v \n", name, result)
		}
	}

	var a []map[struct {
		a int
		b struct {
			x string
			y bool
		}
	}]interface {
		Build([]byte, struct {
			x string
			y bool
		}) error
		Update(dt float64)
		Destroy()
	}
	fmt.Println(a)
}
func main2() {
	e, err := casbin.NewEnforcer("D:\\goworkstation\\Study\\Casbin\\keymatch_model.conf", "D:\\goworkstation\\Study\\Casbin\\keymatch_policy.csv")

	fmt.Printf("RBAC TENANTS test start\n") // output for debug

	// superAdmin
	if falg, _ := e.Enforce("superAdmin", "gy", "project", "read"); falg {
		log.Println("superAdmin can read project in gy")
	} else {
		log.Fatal("ERROR: superAdmin can not read project in gy")
	}

	if falg, _ := e.Enforce("superAdmin", "gy", "project", "write"); falg {
		log.Println("superAdmin can write project in gy")
	} else {
		log.Fatal("ERROR: superAdmin can not write project in gy")
	}

	if falg, _ := e.Enforce("superAdmin", "jn", "project", "read"); falg {
		log.Println("superAdmin can read project in jn")
	} else {
		log.Fatal("ERROR: superAdmin can not read project in jn")
	}

	if falg, _ := e.Enforce("superAdmin", "jn", "project", "write"); falg {
		log.Println("superAdmin can write project in jn")
	} else {
		log.Fatal("ERROR: superAdmin can not write project in jn")
	}

	// admin
	if falg, _ := e.Enforce("quyuan", "gy", "project", "read"); falg {
		log.Println("quyuan can read project in gy")
	} else {
		log.Fatal("ERROR: quyuan can not read project in gy")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "project", "write"); falg {
		log.Println("quyuan can write project in gy")
	} else {
		log.Fatal("ERROR: quyuan can not write project in gy")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "project", "read"); falg {
		log.Fatal("ERROR: quyuan can read project in jn")
	} else {
		log.Println("quyuan can not read project in jn")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "project", "write"); falg {
		log.Fatal("ERROR: quyuan can write project in jn")
	} else {
		log.Println("quyuan can not write project in jn")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "asse", "read"); falg {
		log.Fatal("ERROR: quyuan can read asse in gy")
	} else {
		log.Println("quyuan can not read asse in gy")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "asse", "write"); falg {
		log.Fatal("ERROR: quyuan can write asse in gy")
	} else {
		log.Println("quyuan can not write asse in gy")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "asse", "read"); falg {
		log.Println("quyuan can read asse in jn")
	} else {
		log.Fatal("ERROR: quyuan can not read asse in jn")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "asse", "write"); falg {
		log.Println("quyuan can write asse in jn")
	} else {
		log.Fatal("ERROR: quyuan can not write asse in jn")
	}

	// wenyin
	if falg, _ := e.Enforce("wenyin", "gy", "asse", "write"); falg {
		log.Println("wenyin can write asse in gy")
	} else {
		log.Fatal("ERROR: wenyin can not write asse in gy")
	}

	if falg, _ := e.Enforce("wenyin", "jn", "asse", "write"); falg {
		log.Fatal("ERROR: wenyin can write asse in jn")
	} else {
		log.Println("wenyin can not write asse in jn")
	}

	// shangshang
	if falg, _ := e.Enforce("shangshang", "jn", "project", "write"); falg {
		log.Println("shangshang can write project in jn")
	} else {
		log.Fatal("ERROR: shangshang can not write project in jn")
	}

	if falg, _ := e.Enforce("shangshang", "gy", "project", "write"); falg {
		log.Fatal("ERROR: shangshang can write project in gy")
	} else {
		log.Println("shangshang can not write project in gy")
	}
	fmt.Println(err)
}

func main1() {

	e, err := casbin.NewEnforcer("D:\\goworkstation\\Study\\Casbin\\keymatch_model.conf", "D:\\goworkstation\\Study\\Casbin\\keymatch_policy.csv")

	fmt.Printf("RBAC test start\n") // output for debug

	// superAdmin
	if falg, _ := e.Enforce("superAdmin", "project", "read"); falg {
		log.Println("superAdmin can read project")
	} else {
		log.Fatal("ERROR: superAdmin can not read project")
	}

	if falg, _ := e.Enforce("superAdmin", "project", "write"); falg {
		log.Println("superAdmin can write project")
	} else {
		log.Fatal("ERROR: superAdmin can not write project")
	}

	// admin
	if falg, _ := e.Enforce("quyuan", "project", "read"); falg {
		log.Println("quyuan can read project")
	} else {
		log.Fatal("ERROR: quyuan can not read project")
	}

	if falg, _ := e.Enforce("quyuan", "project", "write"); falg {
		log.Println("quyuan can write project")
	} else {
		log.Fatal("ERROR: quyuan can not write project")
	}

	if falg, _ := e.Enforce("quyuan", "asse", "read"); falg {
		log.Println("quyuan can read asse")
	} else {
		log.Fatal("ERROR: quyuan can not read asse")
	}

	if falg, _ := e.Enforce("quyuan", "asse", "write"); falg {
		log.Println("quyuan can write asse")
	} else {
		log.Fatal("ERROR: quyuan can not write asse")
	}

	// zhuangjia
	if falg, _ := e.Enforce("wenyin", "project", "read"); falg {
		log.Fatal("ERROR: wenyin can read project")
	} else {
		log.Println("wenyin can not read project")
	}

	if falg, _ := e.Enforce("wenyin", "project", "write"); falg {
		log.Println("wenyin can write project")
	} else {
		log.Fatal("ERROR: wenyin can not write project")
	}

	if falg, _ := e.Enforce("wenyin", "asse", "read"); falg {
		log.Fatal("ERROR: wenyin can read asse")
	} else {
		log.Println("wenyin can not read asse")
	}

	if falg, _ := e.Enforce("wenyin", "asse", "write"); falg {
		log.Println("wenyin can write asse")
	} else {
		log.Fatal("ERROR: wenyin can not write asse")
	}

	// shangshang
	if falg, _ := e.Enforce("shangshang", "project", "read"); falg {
		log.Println("shangshang can read project")
	} else {
		log.Fatal("ERROR: shangshang can not read project")
	}

	if falg, _ := e.Enforce("shangshang", "project", "write"); falg {
		log.Fatal("ERROR: shangshang can write project")
	} else {
		log.Println("shangshang can not write project")
	}

	if falg, _ := e.Enforce("shangshang", "asse", "read"); falg {
		log.Println("shangshang can read asse")
	} else {
		log.Fatal("ERROR: shangshang can not read asse")
	}

	if falg, _ := e.Enforce("shangshang", "asse", "write"); falg {
		log.Fatal("ERROR: shangshang can write asse")
	} else {
		log.Println("shangshang can not write asse")
	}

	fmt.Println(err)
	//	// 使用MySQL数据库初始化一个Xorm适配器
	//	engine, err := xorm.NewEngine("mysql", "root:193792@tcp(127.0.0.1:3306)/test")
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//	a, err := xormadapter.NewAdapterByEngine(engine)
	//	if err != nil {
	//		log.Fatalf("error: adapter: %s", err)
	//	}
	//
	//	m, err := model.NewModelFromString(`
	//[request_definition]
	//r = sub, obj, act
	//
	//[policy_definition]
	//p = sub, obj, act
	//
	//[policy_effect]
	//e = some(where (p.eft == allow))
	//
	//[matchers]
	//m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	//`)
	//	if err != nil {
	//		log.Fatalf("error: model: %s", err)
	//	}
	//
	//	e, err := casbin.NewEnforcer(m, a)
	//	if err != nil {
	//		log.Fatalf("error: enforcer: %s", err)
	//	}
	//	fmt.Println(e)
}
