package main

import "fmt"

var length int

func atas (){
    for i:= 0; i < length/2 + 1; i++{
        for j:= 0; j < i ; j++{
            fmt.Print("  ")
        }
        for j:= i; j < length-i; j++{
            fmt.Print("* ")
        }
        for j:= length - i; j < length + 1; j++{
            fmt.Print("  ")
        }
        fmt.Println("")
    }
}

func bawah (){
    for i:= 1; i < length/2 + 1; i++{
        for j:= 0; j < length/2 - i; j++{
            fmt.Print("  ")
        } 
        for j:= length/2 - i; j <= length/2 + i; j++{
            fmt.Print("* ")
        }
        for j:= length/2 + i; j < length; j++{
            fmt.Print("  ")
        }
        fmt.Println("")
    }
}

func main(){
    fmt.Scan(&length)
	if length%2 == 0 {
		length--
	}

    atas()
    bawah()
}
