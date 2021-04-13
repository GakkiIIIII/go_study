// defer使用
package main

import "fmt"

func test() {
	a, b := 10, 20

	// 当有多个defer时，会根据声明顺序入栈。而且会保存当前变量的一个副本(值类型和引用类型都是一样的效果)
	defer fmt.Printf("a的值为:%d\n", a) // 3
	defer fmt.Printf("b的值为:%d\n", b) // 2
	a++
	b++
	res := a + b
	fmt.Printf("res的值为:%d\n", res) // 1
}

func test2() {
	arr := []int{1, 2, 3, 4}
	defer fmt.Printf("arr数组：%v\n", arr)
	arr = append(arr, 5)
	fmt.Println(arr)
}
func main() {
	// test()
	test2()
}
