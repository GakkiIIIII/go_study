package main

import (
	"fmt"
	"strings"
)

func makeSuffix(file string) func(string) string {
	return func(suffix string) string {
		if strings.HasSuffix(file, suffix) {
			return file
		} else {
			return file + suffix
		}
	}
}

func main() {
	hasJPG := makeSuffix("生日聚会")
	fmt.Println(hasJPG(".jpg"))
}
