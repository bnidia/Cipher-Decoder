package main

import "fmt"

func main() {
	const a = 21
	const b = 15
	var g, p int
	var b_key int = 1
	var s_key int = 1
	var a_key int

	fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Println("OK")
	fmt.Scanf("A is %d", &a_key)

	for i := 1; i <= b; i++ {
		b_key = (b_key * g) % p
		s_key = (s_key * a_key) % p
	}
	fmt.Printf("B is %d\n", b_key)
	fmt.Printf("A is %d\n", a_key)
	fmt.Printf("S is %d\n", s_key)
}
