package struct_demo

import "fmt"

//func main()  {
//	TestForStruct()
//}

// Animal 父类
type Animal struct {
	Color string
}

// Dog 继承自Animal
type Dog struct {
	Animal
	Id int
	Name string
	Age int
}

// Dog 定义 struct
//type Dog struct {
//	Id int
//	Name string
//	Age int
//}


// Cate
type  Cat struct {
	Animal
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

func (d *Dog) Run() string {
	fmt.Println("id", d.Id, "Dog is running")
	return "Dog is running"
}


func (d *Dog) Eat() string{
	fmt.Println("animal dog is eating")
	return "animal dog is eating"
}

func (d *Cat) Run() string {
	fmt.Println("id", d.Id, "Cat is running")
	return "Cat is running"
}

func (d *Cat) Eat() string{
	fmt.Println("animal cat is eating")
	return "animal cat is eating"
}