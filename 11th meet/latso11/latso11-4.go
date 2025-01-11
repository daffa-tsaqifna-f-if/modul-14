package main

import "fmt"

func main() {
	var x, y string
	var z, i int
	i = 1
	for {
		fmt.Printf("Bunga %d:", i)
		fmt.Scan(&x)
		if x == "SELESAI" {
			break
		}
		y = y + x + " - "
		i++
		z++
	}
	fmt.Print("Pita: ", y, "\nBunga: ", z)
}
