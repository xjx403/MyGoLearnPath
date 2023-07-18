package main

import "fmt"

func main() {

	fmt.Println("*******string:******")
	var str string = "E人"
	fmt.Printf("str: %s, len = %d\n", str, len(str))
	str += "e"
	fmt.Printf("str: %s, len = %d\n", str, len(str))
	
	fmt.Printf("\n*******pointer******\n")
	var inta int = 9
	pointerShow(&inta)
}

func pointerShow(p *int) {
	if p == nil {
		fmt.Println("Input Point is NIL, please check!")
	}else {
		fmt.Printf("Before change the value is: %d\n", *p)
		*p++;
		fmt.Printf("After change the value is: %d\n", *p)
	}
}
