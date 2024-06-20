package main

import (
	"fmt"
)

func main(){
	var c string
    n,err := fmt.Scanln(&c)
	fmt.Println(c)
     if err != nil {
         fmt.Println("Error:", err,n)
     } else {
         fmt.Println("Input:", c,n)
     }
}