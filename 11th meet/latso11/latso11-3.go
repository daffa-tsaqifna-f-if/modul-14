package main

import "fmt"

func main() {
	var m, k, h, u,t string
	var x bool
	m = "merah"
	k = "kuning"
	h = "hijau"
	u = "ungu"
	x = true
	for i := 0; i < 5; i++ {
		fmt.Printf("percobaan %d: ",i+1)
		fmt.Scan(&t)
		if t!=m{
			x=false
		}
		fmt.Scan(&t)
		if t!=k{
			x=false
		}
		fmt.Scan(&t)
		if t!=h{
			x=false
		}
		fmt.Scan(&t)
		if t!=u{
			x=false
		}
	}
	fmt.Print("Berhasil : ",x)
}