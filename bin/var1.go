package main

import "fmt"

var c, python, java bool //Variables declared without an explicit initial value are given their zero value.

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
