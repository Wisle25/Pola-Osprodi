package main

import "fmt"

func main() {
	var length int

	fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}
	half := (length + 1) / 2

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j > 0 && j < half-1 && i > 0 && i < half-1 {
				fmt.Print("  ")
			} else if j > 0 && j < half-1 && i >= half && i < length-1 {
				fmt.Print("  ")
			} else if j >= half && j < length-1 && i > 0 && i < half-1 {
				fmt.Print("  ")
			} else if j >= half && j < length-1 && i >= half && i < length-1 {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println("")
	}
}
