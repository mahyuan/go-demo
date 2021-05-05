package point_demo

import "fmt"

func TestPointArr()  {
	// 指针数组
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组:", pointArr) //  [0xc00001a0c0 0xc00001a0c8]

	// 数组指针
	arr := [...]int{3,4,5}
	arrPoint := &arr
	fmt.Println("数组指针:", arrPoint) // &[3 4 5]
}

// TestPoint 指针操作
func  TestPoint()  {
	var count int = 30
	var countPoint *int
	var countPoint1 *int
	countPoint = &count
	fmt.Printf("count 变量地址: %x \n", &count)
	if countPoint != nil {
		fmt.Printf("countPoint 变量存储的地址: %x \n", countPoint)
	}
	fmt.Printf("countPoint指针指向地址的值: %d \n", *countPoint)
	fmt.Printf("countPoint1 变量存储的地址: %x \n", countPoint1)
}
