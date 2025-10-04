package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "adam"
	names[1] = "bravo"
	names[2] = "Charlie"

	fmt.Println(names)
	fmt.Println(names[1])

	// declare array secara langsung
	var nilai = [...]int{
		90,
		80,
		85,
		70,
	}
	fmt.Println(nilai)

	// edit array
	fmt.Println(len(nilai))
	fmt.Println(nilai[1])

	nilai[3] = 999
	fmt.Println(nilai)
}
