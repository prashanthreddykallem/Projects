package main

import "fmt"
var f = 3 //Short hand assignment is not valid for global declaration
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java,f)
}