package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   //numbers1 := make([]int, len(numbers), (cap(numbers))*2)
   var numbers1 []int
   numbers1=append(numbers)
   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)
   fmt.Println("-----------------------------------------------")

   fmt.Println("numbers:",numbers)
   fmt.Println("numbers1:",numbers1)

   numbers[1]=22
   fmt.Println("numbers:",numbers)
   fmt.Println("numbers1:",numbers1)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}