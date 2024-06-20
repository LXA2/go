package main

import "fmt"

var a=make([]int,7,7)

var m = make(map[string]int)


func main(){
	a[0]=5
	a = []int{5,6,7,8,9,1}
	b := a[1:]
	copy(a,b)
	fmt.Println("length:",len(a))
	fmt.Println("capacity:",cap(a))
	fmt.Println("content(a):",a)
	fmt.Println(b)
	for i , q := range a {
		fmt.Println("i",i,"q",q)
		//fmt.Println(q)
	}


	m["a"]=123
	fmt.Println(m)
}
