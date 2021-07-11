package main

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/tunny"
	"golang.org/x/xerrors"
	"log"
	"math"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 控制协程(goroutine)的并发数量

func TestOne1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

/*
panic: too many concurrent operations on a single file or socket (max 1048575)
对单个 file/socket 的并发操作个数超过了系统上限，这个报错是 fmt.Printf 函数引起的，fmt.Printf 将格式化后的字符串打印到屏幕，即标准输出。
在 linux 系统中，标准输出也可以视为文件，内核(kernel)利用文件描述符(file descriptor)来访问文件，标准输出的文件描述符为 1，
错误输出文件描述符为 2，标准输入的文件描述符为 0。

简而言之，系统的资源被耗尽了。

那如果我们将 fmt.Printf 这行代码去掉呢？那程序很可能会因为内存不足而崩溃。这一点更好理解，每个协程至少需要消耗 2KB 的空间，
那么假设计算机的内存是 2GB，那么至多允许 2GB/2KB = 1M 个协程同时存在。
那如果协程中还存在着其他需要分配内存的操作，那么允许并发执行的协程将会数量级地减少。
*/

// 利用 channel 的缓存区限制并发的协程数量
func TestOne2(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

// 利用第三方库 目前有很多第三方库实现了协程池，可以很方便地用来控制协程的并发数量
// https://github.com/Jeffail/tunny
// https://github.com/panjf2000/ants
func TestOne3(t *testing.T) {
	// 第一个参数是协程池的大小(poolSize)，第二个参数是协程运行的函数(worker)。
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	// 关闭协程池。
	defer pool.Close()
	// 将参数 i 传递给协程池定义好的 worker 处理。
	for i := 0; i < 10; i++ {
		go pool.Process(i)
	}
	time.Sleep(time.Second * 4)
}

// Go sync.Pool
/*
1 sync.Pool 的使用场景
一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力。
*/
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func TestOne4(t *testing.T) {

	stu := &Student{}
	json.Unmarshal(buf, stu)
}

func TestOne5(t *testing.T) {
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	stu := studentPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
	fmt.Println(studentPool.Get())
}

// Go sync.Once
/*
考虑一个简单的场景，函数 ReadConfig 需要读取环境变量，并转换为对应的配置。环境变量在程序执行前已经确定，执行过程中不会发生改变。
ReadConfig 可能会被多个协程并发调用，为了提升性能（减少执行时间和内存占用），使用 sync.Once 是一个比较好的方式。

在这个例子中，声明了 2 个全局变量，once 和 config；
config 是需要在 ReadConfig 函数中初始化的(将环境变量转换为 Config 结构体)，ReadConfig 可能会被并发调用。
如果 ReadConfig 每次都构造出一个新的 Config 结构体，既浪费内存，又浪费初始化时间。如果 ReadConfig 中不加锁，
初始化全局变量 config 就可能出现并发冲突。这种情况下，使用 sync.Once 既能够保证全局变量初始化时是线程安全的，又能节省内存和初始化时间。
*/
type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Println("init config")
	})
	return config
}
func TestOne6(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}

// Go sync.Cond

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		fmt.Println("------------")
		c.Wait()
	}
	log.Println(name, "starts reading", time.Now())
	time.Sleep(time.Second)
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func TestOne7(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

func TestOne8(t *testing.T) {
	err := xerrors.New("sdfv")
	err = nil
	_, ok := err.(error)
	fmt.Println(ok)
}

type Result struct {
	Num, Ans int
}

type Calc int

// Square calculates the square of num
func (calc *Calc) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}
func TestOne9(t *testing.T) {
	rpc.Register(new(Calc))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}

type Demo struct {
	name string
}

func TestOne10(t *testing.T) {
	var a1 *Demo = &Demo{name: "hjvakf"}
	inter(a1)
	fmt.Println(a1)
}
func inter(a interface{}) {
	fmt.Println(a)
	a1 := a.(*Demo)
	a1.name = "000"
	fmt.Println(a1)
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func TestOne19(t *testing.T) {

	fmt.Println(CurWeekStart(time.Now().AddDate(0, 0, -30)).Format("2006-01-02"))
	fmt.Println(CurWeekEnd(time.Now().AddDate(0, 0, -1)).Format("2006-01-02"))
	fmt.Println(CurMonthStart(time.Now().AddDate(0, 0, -30)).Format("2006-01-02"))
	fmt.Println(CurMonthEnd(time.Now().AddDate(0, 0, -1)).Format("2006-01-02"))
}

// CurMonthEnd 当月末
func CurMonthEnd(t time.Time) time.Time {
	return CurMonthStart(t).AddDate(0, 1, -1)
}

// LastMonthStart 上月初
func LastMonthStart(t time.Time) time.Time {
	return CurMonthStart(t).AddDate(0, -1, 0)
}

// LastMonthEnd 上月末
func LastMonthEnd(t time.Time) time.Time {
	return CurMonthStart(t).AddDate(0, 0, -1)
}

// CurMonthStart 当月初
func CurMonthStart(t time.Time) time.Time {
	curYear, curMonth, _ := t.Date()
	curLocation := t.Location()
	return time.Date(curYear, curMonth, 1, 0, 0, 0, 0, curLocation)
}

// CurWeekEnd 这周末
func CurWeekEnd(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return t.AddDate(0, 0, int(7-weekday))
}

// CurWeekStart 这周初
func CurWeekStart(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return t.AddDate(0, 0, int(-weekday+1))
}

// LastWeekStart 上周初
func LastWeekStart(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return t.AddDate(0, 0, int(-weekday-6))
}

// LastWeekEnd 上周末
func LastWeekEnd(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	return t.AddDate(0, 0, int(-weekday))
}
