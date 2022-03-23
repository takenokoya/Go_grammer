package main

import "fmt"

// 一番最初に呼ばれる特別な関数
func init() {
	fmt.Println("Init")
}

// 通常の関数はmain関数から呼ぶ必要がある
func buzz() {
	fmt.Println("Buzz")
}

// main関数
func main() {
	buzz()
	fmt.Println("Hello, World!", "test")
}
