package main

import "fmt"

func main() {
	var i int
	var f float64
	var str string
	var b bool

	fmt.Println("请输入...")
	fmt.Scanln(&i)

	fmt.Scanln(&f)
	fmt.Scanln(&str)
	fmt.Scanln(&b)

	fmt.Printf("i = %v  f = %v  str = %v  b = %v\n", i, f, str, b)
}
