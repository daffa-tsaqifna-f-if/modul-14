package main

import "fmt"

func main() {
	var x int
	var y, z, t float64
	t = 1
	fmt.Print("Nilai K = ")
	fmt.Scan(&x)
	for i := 0; i <= x; i++ {
		y = float64((4*i)+2) * float64((4*i)+2)
		z = float64((4*i)+1) * float64((4*i)+3)
		t = t * (y / z)
	}
	fmt.Printf("Nilai akar 2 = %.10f", t)
}
