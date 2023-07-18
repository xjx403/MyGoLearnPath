package main

import "fmt"

type Student struct {
	id int
	name string
	gender string
	score [3]int
}


func main() {

	fmt.Println("*******string:******")
	var str string = "E人"
	fmt.Printf("str: %s, len = %d\n", str, len(str))
	str += "e"
	fmt.Printf("str: %s, len = %d\n", str, len(str))
	
	fmt.Printf("\n*******pointer******\n")
	var inta int = 9
	pointerShow(&inta)

	fmt.Printf("\n*******struct and method******\n")
	var student1 Student
	var score = [3]int{90, 100, 100}
	student1.setStudent(1,"小明","man",score)
	fmt.Println("Inited object:",student1)

	fmt.Printf("\n*******function show******\n")
	var x, y int = 11, 23
	fmt.Printf("Max (%d, %d) function End:%d\n", x, y, maxOfTwo(x, y))

	fmt.Printf("\n*******slipe show******\n")
	slipeShow()
}

func maxOfTwo(x int, y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}

/*
	指针演示
*/
func pointerShow(p *int) {
	if p == nil {
		fmt.Println("Input Point is NIL, please check!")
	}else {
		fmt.Printf("Before change the value is: %d\n", *p)
		*p++;
		fmt.Printf("After change the value is: %d\n", *p)
	}
}

/*
	结构体方法，本质上还是函数，要想修改结构体，还是需要传指针
*/
func (stu *Student) setStudent(id int, name string, gender string, score [3]int) {
	stu.id = id
	stu.name = name
	stu.gender = gender
	for i := 0; i < len(stu.score); i++ {
		stu.score[i] = score[i]
	}
} 

func printSlipe(x []int) {
	fmt.Printf("slipe len is %d, capatity is %d, : %v \n", len(x), cap(x), x)
}
func slipeShow() {
	slipeDemo1 := make([]int, 3, 10)
	printSlipe(slipeDemo1)
	slipeDemo1 = append(slipeDemo1, 1)
	printSlipe(slipeDemo1)
	slipeDemo1 = append(slipeDemo1, 2, 3, 4)
	printSlipe(slipeDemo1)

	var slipeDemo2 = make([]int, len(slipeDemo1), cap(slipeDemo1)*2)
	copy(slipeDemo2, slipeDemo1)
	printSlipe(slipeDemo2)
}
