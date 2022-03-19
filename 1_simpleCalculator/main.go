package main

import "fmt"

func penjumlahan(a, b int) int {
	return a + b
}

func pengurangan(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func perkalian(a, b int) int {
	return a*b
}

func pembagian(a, b int) int {
	if a < b {
		return b/a
	} else {
		return a/b
	}
}

func main() {

	var mode string
	var a int
	var b int
	var c int

	fmt.Print("Pilih mode (+-x/) : ")
	fmt.Scanln(&mode)
	fmt.Print("a: ")
	fmt.Scanln(&a)
	fmt.Print("b: ")
	fmt.Scanln(&b)

	if mode == "+" {
		c = penjumlahan(a, b)
	} else if mode == "-" {
		c = pengurangan(a,b)
	} else if mode == "x" {
		c = perkalian(a,b)
	} else if mode == "/" {
		c = pembagian(a,b)
	}

	fmt.Println(c)
}
