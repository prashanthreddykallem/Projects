package main

import "fmt"

func main() {
	//START OMIT
	sum := 1
	for sum <= 100 {
		sum += sum
	}
	fmt.Println(sum)
	//END OMIT
}
