package main

import "fmt"

var length int

func main() {
	fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}

	for y := 0; y < length/2; y++ {
		for x := 0; x < length; x++ {
			if y == x {
				fmt.Print("* ")
			} else if x == length/2 {
				fmt.Print("* ")
			} else if x == length-1-y {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}

	for i := 0; i < length; i++ {
		fmt.Print("* ")
	}
	fmt.Println("")

	for b := 0; b < length/2; b++ {
		for a := 0; a < length; a++ {
			if a == length/2-1-b {
				fmt.Print("* ")
			} else if a == (length+1)/2-1 {
				fmt.Print("* ")
			} else if a == (length+1)/2+b {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}
