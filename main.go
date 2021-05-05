package main

import (
	"errors"
	"fmt"
	"go_demo1/interface_demo"
	"go_demo1/point_demo"
	"reflect"
)

func main() {
	// 指针数组
	point_demo.TestPointArr()


	// 指针创建
	//point_demo.TestPoint()


	// 测试协程间通信
	//gorotine.Read()
	//go gorotine.Write()
	//gorotine.WG.Wait()
	//fmt.Println("All Done")
	//time.Sleep(time.Second * 60)

	// 协程发送数据和接收数据
	//go gorotine.Send()
	//go gorotine.Receive()
	//time.Sleep(time.Second * 10)


	/*
	// 协程
	// 使用go关键字启动协程
	num := runtime.NumCPU()
	fmt.Printf("cpu num: %d \n", num)

	runtime.GOMAXPROCS(num - 1)
	go gorotine.Loop()
	go gorotine.Loop1()
	time.Sleep(time.Second * 5)
	*/

	// test for struct
	//testForStruct()
	//struct_demo.TestForStruct()

	// test for Run()
	//dog := new(struct_demo.Dog)
	//dog.Id = 10005
	//dog.Name = "花花"
	//dog.Age = 6
	//dog.Color = "Red"
	//dog.Run()
	//dog.Eat()

	// test for interface
	//var b interface_demo.Behaver
	//b = new(struct_demo.Cat)
	//b.Run()

/*	dog := new(struct_demo.Dog)
	cat := new(struct_demo.Cat)
	//dog.Run()
	//cat.Run()
	action(dog)
	action(cat)*/


	//Show2()
	//show1.Show1()

	//    fmt.Println("Hello, World!")
	//    show2.shows()

	/*var a = [...]int{ 100, 101, 102 }
	fmt.Println("a[0]", a[0])
	fmt.Println("a[1]", a[1])
	fmt.Println("a[2]", a[2])

	var b = &a
	fmt.Println("b[0]", b[0])
	fmt.Println("b[1]", b[1])
	fmt.Println("b[2]", b[2])

	for i, v := range b {
		fmt.Println(i, v)
	}


	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}

	var c = [...]int{2: 3, 1: 2}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}*/

	//makeSlice()
	//makeMap()
	//makeChan()

	//newMap()

	//appendElementForSlice()

	//copySlice()

	//deleteFromMap()

	//recoverPanic()
	//receivePanic()

	//getLen()

	//closeChan()
}

func action(b interface_demo.Behaver) string  {
	b.Run()
	b.Eat()
	return ""
}
func getLen()  {
	mMap := make(map[int]string)
	mMap[0] = "cat"
	mMap[1] = "dog"
	mMap[4] = "ddd"

	fmt.Println("len(mMap):", len(mMap))
	//fmt.Println("cap(mMap):", cap(mMap))

	mSlice := make([]string, 4)
	mSlice[0] = "slice-1"
	mSlice[1] = "slice-2"
	//append(mSlice, "append-item")
	fmt.Println("len(mSlice):", len(mSlice))
	fmt.Println("cap(mSlice):", cap(mSlice))
}


// 创建和关闭chan
func  closeChan()  {
	mChan := make(chan int, 1)
	mChan <- 1
	defer close(mChan)
	// 业务代码
}

// panic抛出异常，recover捕获异常
func recoverPanic()  {
	defer func() {
		message := recover()
		fmt.Println("recover panic message:", message) // recover panic message: I am a panic
	}()
	panic("I am a panic")
}

// 封装一下
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




// 删除map中的数据
func deleteFromMap()  {
	mIDMap := make(map[int]string)
	mIDMap[0] = "id-1"
	mIDMap[1] = "id-2"

	delete(mIDMap, 0)
	fmt.Println(mIDMap) //	map[1:id-2
	delete(mIDMap, 1)
	fmt.Println(mIDMap) // map[]
}


func copySlice() {
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

func  appendElementForSlice()  {
	mIDSlice := make([]string, 2 )
	mIDSlice[0] = "id-1"
	mIDSlice[1] = "id-2"

	mIDSlice = append(mIDSlice, "id-3")
	fmt.Println("test append for slice", mIDSlice) // test append for slice [id-1 id-2 id-3]
	fmt.Println("len of mIDSlice", len(mIDSlice))
	fmt.Println("cap=", cap(mIDSlice))
}

func  newMap()  {
	mNewMap := new (map[int]string)
	mMakeMap := make(map[int]string)
	fmt.Println("mNewMap", reflect.TypeOf(mNewMap))   // *map[int]string
	fmt.Println("mMakeMap", reflect.TypeOf(mMakeMap)) // map[int]string
}

//	make 创建切片
func makeSlice()  {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"
	fmt.Println("mSlice", mSlice) // mSlice [dog cat tiger]

}

// 创建map
func makeMap()  {
	mMap := make(map[int]string, 3)
	mMap[10] = "dog"
	mMap[101] = "tiger"
	mMap[102] = "cat"
	fmt.Println("mMap", mMap) //mMap map[10:dog 101:tiger 102:cat]

}

// 创建channel类型, 返回的是引用类型
func makeChan()  {
	// type, size, interType ; size 是缓存的尺寸, 没有缓存的chan，不用传第二个参数size
	mChan := make(chan int)
	close(mChan)


	fmt.Println("mChan", mChan) //mChan 0xc000064060

}


/*import (
	"github.com/gin-gonic/gin"
	"net/http"
)*/

/*
import (

	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	 //example curl for /admin with basicauth header
	 //  Zm9vOmJhcg== is base64("foo:bar")
	 //
		//curl -X POST \
	 // 	http://localhost:8080/admin \
	 // 	-H 'authorization: Basic Zm9vOmJhcg==' \
	 // 	-H 'content-type: application/json' \
	 // 	-d '{"value":"bar"}'
	 //
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
*/


//package main

//func main() {
////    fmt.Println("Hello, World!")
////    show2.shows()
//
//	var a = [...]int{ 100, 101, 102 }
//	fmt.Println("a[0]", a[0])
//	fmt.Println("a[1]", a[1])
//	fmt.Println("a[2]", a[2])
//}

/*
func fn() {
    fmt.Println("hello")
}


package main


import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    fmt.Println("Please visit http://127.0.0.1:12345/")
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        s := fmt.Sprintf("你好, 世界! -- Time: %s", time.Now().String())
        fmt.Fprintf(w, "%v\n", s)
        log.Printf("%v\n", s)
    })
    if err := http.ListenAndServe(":12345", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    fn()

}

func fn() {
    //fmt.Println("Please visit http://127.0.0.1:12345/")

    http.HandleFunc("/fn", func(w http.ResponseWriter, req *http.Request) {
        s1 := fmt.Sprintf("你好啊，老乡.... == Time: %s", time.Now().String())
        fmt.Fprintf(w, "%v\n", s1)
        log.Printf("%v\n", s1)
    })
    if err := http.ListenAndServe(":12345", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}*/