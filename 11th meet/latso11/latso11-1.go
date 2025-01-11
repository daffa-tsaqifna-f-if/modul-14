package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if i%2 != 0 {
			y++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", y)
}
