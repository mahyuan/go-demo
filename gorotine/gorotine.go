package gorotine

import (
	"fmt"
	"sync"
	"time"
)

var WG sync.WaitGroup

// Read 读取数据
func Read() {
	for i := 0; i < 11; i++ {
		WG.Add(1)
	}
}

// Write 写入数据
func  Write()  {
	for i:=0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done", i)
		WG.Done()
	}
}

/*
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


	//num := <- chanInt
	//fmt.Println("num: ", num)
	//
	//num = <- chanInt
	//fmt.Println("num: ", num)
	//
	//num = <- chanInt
	//fmt.Println("num: ", num)
}
*/
/*
func Loop()  {
	for i := 0; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Printf("%d ", i)
	}
}

func Loop1()  {
	for i := 0; i < 11; i++ {
		time.Sleep(time.Microsecond * 500)
		fmt.Printf("%d ", i)
	}
}*/