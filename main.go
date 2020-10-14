package main

import "fmt"

func main() {
	var tamano int
	var s []int
	var v int
	var suma int

	fmt.Scan(&tamano)

	for i := 0; i < tamano; i++ {
		fmt.Scan(&v)
		s = append(s, v)
	}
	for _, v := range s {
		suma += v
	}
	fmt.Println(suma)
}
