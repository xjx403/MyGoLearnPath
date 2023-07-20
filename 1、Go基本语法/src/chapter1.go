package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID int 
	Name string
	Gender string
	Score [3]int
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
	var Score = [3]int{90, 100, 100}
	student1.setStudent(1,"小明","man",Score)
	fmt.Println("IniteScored object:",student1)
	//序列化
	jsons, errs := json.Marshal(&student1)
	if errs != nil {
		fmt.Println("json marshal error；", errs)
	} 
	fmt.Println("json data:", string(jsons))

	fmt.Printf("\n*******function show******\n")
	var x, y int = 11, 23
	fmt.Printf("Max (%d, %d) function End:%d\n", x, y, maxOfTwo(x, y))

	slipeShow()

	mapShow()

	forShow()
}

func preHeadInfor(message string) {
	fmt.Printf("\n*******%s******\n",message)
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
	stu.ID = id
	stu.Name = name
	stu.Gender = gender
	for i := 0; i< len(stu.Score); i++ {
		stu.Score[i] = score[i]
	}
} 

func printSlipe(x []int) {
	fmt.Printf("slipe len is %d, capatity is %d, : %v \n", len(x), cap(x), x)
}

func slipeShow() {
	preHeadInfor("slipe show")
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

func mapShow() {
	preHeadInfor("map show")
	var m1 map[int]string
	m1 = make(map[int]string)
	m1[1] = "java"
	fmt.Println("map1:", m1)

	var m2 map[int]string = map[int]string{}
	m2[1] = "go"
	fmt.Println("map2:", m2)

	m3 := map[int]string{
		1:"C++",
		2:"C",
	}
	fmt.Println("map3:", m3)

	fmt.Println("");
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age" : "30",
		"hobby" : [] string{"读书","爬山"},
	}
	fmt.Println("map data:",res)

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("map to json data:", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json to map data:", res2)
	fmt.Println("")

	//编辑和删除
	m3[3] = "汇编"
	fmt.Println("add map3:", m3)
	m3[1] = "CPlusPlus"
	fmt.Println("modify map3:", m3)
	delete(m3, 2)
	fmt.Println("delete map3:", m3)
}

func forShow() {
	preHeadInfor("for show")
	strs := []string{"小明","小花"}
	for i, s := range strs {
		fmt.Println(i, s)
	}

	testMap := map[string]string{
		"ke1":"v1",
		"k2":"v2",
	}
	for key, value := range testMap {
		fmt.Printf("Key is %s - value is: %s\n", key, value)
	}
	for key := range testMap {
		fmt.Printf("key is %s\n", key)
	}
	for _, value := range testMap {
		fmt.Printf("value is %s\n", value)
	}
}
