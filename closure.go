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
	hasJPG := makeSuffix("ηζ₯θδΌ")
	fmt.Println(hasJPG(".jpg"))
}
