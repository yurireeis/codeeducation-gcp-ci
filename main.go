package main

import "fmt"

func Soma(a int, b int) int {
    return a + b
}

func main() {
	result := Soma(5, 5)
	fmt.Println(result)
}