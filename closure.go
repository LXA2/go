package main

import "fmt"

func price(initialBalance float32) (func(float32),func(float32),func() (float32,float32,float32)) {
    balance := initialBalance
    sum := float32(0)

    consume := func(price float32){
        sum += price
    }

    discount := func(percent float32){
        sum *= (percent / 100)
    }

    calculate := func() (float32,float32,float32){
        balance -= sum
        return initialBalance, balance, sum
    }
    return consume, discount, calculate
}

func main() {
    initialBalance := float32(100)
    consume,_,calculate := price(initialBalance)
    consume(30)
    iniBal,bal,sum := calculate()
    fmt.Println("Initial balance:",iniBal,"  Current balance:",bal,"  Sum:",sum)
    
    initialBalance1 := float32(100)
    consume1,discount1,calculate1 := price(initialBalance1)
    consume1(30)
    consume1(10)
    discount1(50)
    iniBal1,bal1,sum1 := calculate1()
    fmt.Println("Initial balance:",iniBal1,"  Current balance:",bal1,"  Sum:",sum1)
    fmt.Println("Initial balance:",iniBal,"  Current balance:",bal,"  Sum:",sum)//again
}