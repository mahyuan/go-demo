## go

### 包

如果一个main导入其他包，包将被顺序导入
如果导入的包中依赖其他包(如B)，会首先导入包B，然后初始化B包中的常量和变量，最后如果B包中有init函数，则会自动执行init
所有包导入完成后才会对main中的常量和变量进行初始化，然后执行manin中的init函数（如果存在），最后执行main函数
如果一个包被导入多次则该包只会被导入一次

import 别名，给包取一个别名，在导入后以别名来调用包中的函数
点（.）操作的含义，点标识的包导入后，调用该包中函数时可以省略前缀包名
下划线操作符的含义，导入该包，但不导入整个包，而是执行包中的init函数，因此 无法通过包名来调用包中的函数，使用下划线操作往往是为了注册包里的引擎，
让外部可以方便的使用。

init 函数执行实机，单一函数内，对全局变量进行求职，然后执行init, 不同包，按导入顺序先后执行

```go
package main
import {
	. 'hello'
	formater 'fmt'
	_ 'math'
}

formater.Println("hello")
```

### 数据类型

- 数值类型： 整型，浮点型，
有符号整型和无符号整型  
复数
  uint8 无符号8位整型（0至256）  
  uint16 无符号16位整型（0至66536）
  uint32 无符号32位整型（0至?）
  uint64
  int8 有符号8位整型(-128至127)
  int16
  int21
  int64

float32
float64
complex64 32位实数和虚数
complex128 64位实数和虚数

byte 类似 uint8
rune 类似 int32
uint 32或64位
int 和uint大小一样
uintptr 无符号整型，用于存储一个指针

最小单位字节，一个字节8位

```go
package  main

import {
	'fmt'
	'unsafe'
}
func main() {
    var i int = 1
    fmt.Print(unsafe.Sizeof(i))
}
```

- 字符串 
字符串类型string，编码统一为 "UTF-8"
  
- 布尔值
布尔型的值可以是常量true或false
  eg: var a bool = true


#### 派生类型

指针类型 (pointer)
数组类型
结构化类型 (struct)
Channel类型 （chan）
函数类型 （func）
切片类型 （slice）
接口类型 （interface）
Map类型 （map）

类型零值，零值不是控制，是声明后的默认值，一般，数值类型默认值为0，
布尔型默认值为false，string默认值为空字符串

类型别名：
使用type 关键字定义别名

```go
package main

import (
  "fmt"
  "reflect"
)

type hello int32

func main() {
  var i hello
  var j int32
  fmt.Println(i)
  fmt.Println(j)
  // 判断类型 reflect.TypeOf(i) => main.hello
  fmt.Println(reflect.TypeOf(i))
}
```

### 变量和常量

单个变量声明和赋值方式：
1) var <变量名称> [变量类型]
2) <变量名称> = <值，表达式，函数等>
3) var <变量名称> [变量类型] = <值，表达式，函数等>
4) 分组式声明格式
```go
package  main
 var (
   i int
   j float32
   name string
   )
```

同一行声明多个变量和赋值
var a, b, c int = 1, 2, 3 或者 a, b := 1, 2 （a=1, b=2, c=3）
全局变量声明必须使用var，局部变量可以省略
特殊变量下划线 "_" （在接下来的程序运行中不再需要该变量）

类型转换
Go中不存在隐式转换，类型转换必须是显式的
类型转换只能发生在两种兼容类型之间
类型转换格式： <变量名称> [:] = <目标类型>（<需要转换的变量>）

可见性规则
大写字母开头的变量可以导出
小写字母开头的变量不可导出，是私有变量


#### make 方法，可以创建slice，map, chan

make返回的是引用类型
```go
package main
import "fmt"

func main()  {
	makeSlice()
    makeMap()
    makeChan()
}

//	make 创建切片
func makeSlice()  {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"
	fmt.Println("mSlice", mSlice)
}

// 创建map
func makeMap()  {
	mMap := make(map[int]string, 3)
	mMap[10] = "dog"
	mMap[101] = "tiger"
	mMap[102] = "cat"
	fmt.Println("mMap", mMap)
}

// 创建channel类型, 返回的是引用类型
func makeChan()  {
	// type, size, interType ; size 是缓存的尺寸, 没有缓存的chan，不用传第二个参数size
	mChan := make(chan int)
	close(mChan)
	fmt.Println("mChan", mChan)
}
```


### new
new 返回的变量是指针类型

```go
package  main
import (
	"fmt"
	"reflect"
)
func  main()  {
	mNewMap := new (map[int]string)
	mMakeMap := make(map[int]string)
	fmt.Println("mNewMap", reflect.TypeOf(mNewMap))   // *map[int]string
	fmt.Println("mMakeMap", reflect.TypeOf(mMakeMap)) // map[int]string
}
```

slice -> append & copy (切片可以append和copy操作)
map -> delete (map里面的key和valye都可以进行delete操作)

copy
```go
package main
import "fmt"

func main() {
	slice1 := make([]string, 2)
	slice2 := make([]string, 3)
	slice1[0] = "slice1-0"
	slice1[1] = "slice1-1"
	//slice1[2] = "slice1-2"

	slice2[0] = "slice2-0"
	slice2[1] = "slice2-1"
	slice2[2] = "slice2-2"

	copy(slice1, slice2)
	fmt.Println(slice1)
}
```
copy 操作，两个参数dst和src，将src拷贝到dst，如果dst的长度大于src，则将src中内容按顺序拷贝到dst，位数不够的部分不变
如果src的长度比dst大，只拷贝长度一致的部分，不会给dst扩容

delete
```go
package main
import "fmt"

// 删除map中的数据
func main() {
	mIDMap := make(map[int]string)
	mIDMap[0] = "id-1"
	mIDMap[1] = "id-2"

	delete(mIDMap, 0)
	fmt.Println(mIDMap) //	map[1:id-2
	delete(mIDMap, 1)
	fmt.Println(mIDMap) // map[]
}

```

panic 抛出异常
recover 捕获异常

```go
package main
import "fmt"

func main()  {
	defer func() {
		message := recover()
		fmt.Println("recover panic message:", message) // recover panic message: I am a panic
	}()
	panic("I am a panic")
}
```
// 封装一下
```go
package main
import "fmt"
import "errors"

func main()  {
  receivePanic()
}
func receivePanic()  {
	defer coverPanic()
	panic("I am panic")
	panic(errors.New("I am error"))
	panic(32)
}
func coverPanic()  {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message:", message)
	case error:
		fmt.Println("error message:", message)
	default:
		fmt.Println("unknow message:", message)
	}
	fmt.Println("recover panic message:", message) // recover panic message: I am a panic
}
```

len 方法获取长度，支持 stirng, array, slice, map, chan
cap 获取容量 , slice, array, chan
close 关闭channel（channel比较占资源）


### 结构体 struct
若干字段的集合
面向对象的实现
定义struct
初始化 var 声明然后初始化
```go
package struct_demo

import "fmt"

// Dog 定义 struct
type Dog struct {
	Id int
	Name string
	Age int
}

func TestForStruct()  {
	// 方式1：
	//var dog Dog
	//dog.Name = "tom"
	//dog.Id = 100001
	//dog.Age = 5

	// 方式2
	//dog := Dog{Id: 100002, Name: "tomcat", Age: 3}

	// 方式3
	dog := new(Dog)
	dog.Id = 10005
	dog.Name = "花花"
	dog.Age = 6

	// dog:  {100001 tom 5}
	fmt.Println("dog: ", dog)
}

func (d *Dog)Run()  {
		fmt.Println("id", d.Id, "Dog is running")
}
```


## interface 接口
是一个接口，把一些公关方法抽取出来封装出来，可以被其他的类实现它, 面向对象的多态属性
实现过程看struct_demo和interface_demo

## 协程
并发实现： 协程
多协程间的通信
多协程间的同步

### 启动协程
多核CPU设置，默认启动所有核心
在执行的方法前面加go关键字可以启动协程

设置启动CPU核心数：
```go
var num = runtime.NumCPU()
runtime.GOMAXPROCS(num - 1)
//fmt.Printf("cpu num: %d \n", num)
```

如上，使用runtime包的NumCPU()方法可以获取到当前电脑的cpu数，使用GOMAXPROCS()方法设置启动的协程数

### 协程通信
channel
select操作
```go

var chanInt chan int = make(chan int, 3)
var timeout chan bool = make(chan bool)

// Send 协程发送数据
func Send() {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}

// Receive 协程接收数据
func Receive() {
	//使用select语句接收数据
	for {
		select {
		case num := <- chanInt:
			fmt.Printf("%d \n", num)
		case <- timeout:
			fmt.Println("timeout....")
		}
	}


	/*num := <- chanInt
	fmt.Println("num: ", num)

	num = <- chanInt
	fmt.Println("num: ", num)

	num = <- chanInt
	fmt.Println("num: ", num)*/
}
```

### 协程同步
系统工具 sync.waitgroup
- add(delta int)添加协程记录
- Done()移除协程记录
- Wait()同步等待所有记录的协程全部结束


### 指针
定义指针变量
为指针变量赋值
访问指针变量中指向地址的值
*取值运算；&取地址运算符

指针 数组的组合：
指针数组（类型是数组，数组的每一个元素都是指针）
数组指针（一个数组的指针）

### json

序列化和反序列化
tag
```go
type Server struct {
	ServerName string `json:"server_name"` // `json:"key"` 转化成json后的key是后面的key，映射关系
	ServerIp string `json:"server_ip"`
	ServerPort int `json:"server_port"`
}
```