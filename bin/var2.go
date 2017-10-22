package main

import "fmt"

var i, j int = 1, 2 //global declaration

func main() {
	var c, python, java = true, false, "no!" //local declaration
	fmt.Println(i, j, c, python, java)
}