package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return y, z, x
}

func main() {
	a, b, c := swap("hello", "world", "extra")
	fmt.Println(a, b, c)
}
