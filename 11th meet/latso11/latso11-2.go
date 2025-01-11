package main

import "fmt"

func main() {
	var x int
	var y bool
	fmt.Scan(&x)
	for i := 2; i < x; i++ {
		if x%i == 0 {
			y = false
			break
		} else {
			y = true
		}
	}
	if y == true {
		fmt.Print("Prima")
	} else {
		fmt.Print("bukan prima")
	}
}
