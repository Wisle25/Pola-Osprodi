package main

import "fmt"

var length int
var half int

func atas() {
	for i := 0; i < half; i++ {
		for j := 0; j < length; j++ {
			if j < half-i {
				fmt.Print("* ")
			} else if j >= half-1+i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func bawah() {
	for i := 1; i < half; i++ {
		for j := 0; j < length; j++ {
			if j < 0+1+i {
				fmt.Print("* ")
			} else if j > length-2-i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}
	half = (length + 1) / 2

	atas()
	bawah()
}
