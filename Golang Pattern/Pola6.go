package main

import "fmt"

var length int

func main() {
	fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			if j == i {
				fmt.Print("* ")
			} else if j == length-i+1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
