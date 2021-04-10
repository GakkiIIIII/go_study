package main

import (
	"fmt"
	"strconv"
)

func main() {
	strI, strF, strB := "123", "45.63", "false"
	// 将字符串转换为对应的类型
	i, _ := strconv.ParseInt(strI, 10, 64)
	fmt.Printf("i 的类型为%T 值为%[1]d\n", i)

	f, _ := strconv.ParseFloat(strF, 64)
	fmt.Printf("f 的类型为%T 值为%[1]f\n", f)

	b, err := strconv.ParseBool(strB)
	fmt.Printf("b 的类型为%T 值为%[1]t\n", b)
	fmt.Println("转换的错误信息为：", err)

	// 将hello强制转换为int类型，看看会有什么反应
	errorParse := "hello"
	pi, err := strconv.ParseInt(errorParse, 10, 32)
	if err != nil {
		fmt.Println("转换发生了错误～～")
	}
	fmt.Printf("pi 的类型为%T 值为%[1]d\n", pi)

	// 将类型转换为字符串
	strNI := fmt.Sprintf("%d", i)
	fmt.Printf("strNI 的类型为%T 字符串为%[1]q\n", strNI)

	strNF := fmt.Sprintf("%5.2f", f)
	fmt.Printf("strNF 的类型为%T 字符串为%[1]q\n", strNF)

	strNB := strconv.FormatBool(b)
	fmt.Printf("strNB 的类型为%T 字符串为%[1]q\n", strNB)

	formatInt := strconv.FormatInt(i, 2)
	fmt.Printf("将int类型转为字符串，并以二进制形式展现%q\n", formatInt)

}