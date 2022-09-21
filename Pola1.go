package main

import "fmt"

var length int

func atas() {
	for i := 0; i <= length/2; i++ {
		for j := 0; j < length; j++ {
			if i%2 == 0 && j%2 != 0 && j > i && j <= length-i {
				fmt.Print("  ")
			} else if i%2 != 0 && j%2 == 0 && j > i && j <= length-i {
				fmt.Print("  ")
			} else if j < i {
				fmt.Print("  ")
			} else if j > length-i {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println("")
	}
}

func bawah() {
	if length%4 == 3 {
		for i := 0; i < length/2; i++ {
			if i%2 == 0 {
				for j := 0; j < length; j++ {
					if j%2 == 0 && j > length/2-3-i && j <= length/2+2+i {
						fmt.Print("* ")
					} else {
						fmt.Print("  ")
					}
				}
			} else if i%2 != 0 {
				for j := 0; j < length; j++ {
					if j%2 != 0 && j > length/2-2-i && j < length/2+2+i {
						fmt.Print("* ")
					} else {
						fmt.Print("  ")
					}
				}
			}
			fmt.Println("")
		}

	} else if length%4 == 1 {
		for i := 1; i < length/2 + 1; i++ {
			if i%2 == 0 {
				for j := 0; j < length; j++ {
					if j%2 == 0 && j > length/2-2-i && j <= length/2+1+i {
						fmt.Print("* ")
					} else {
						fmt.Print("  ")
					}
				}
			} else if i%2 != 0 {
				for j := 0; j < length; j++ {
					if j%2 != 0 && j > length/2-2-i && j < length/2+2+i {
						fmt.Print("* ")
					} else {
						fmt.Print("  ")
					}
				}
			}
			fmt.Println("")
		}
	}
}

func main() {
	fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}

	atas()
	bawah()
}
