package main

import (
	"fmt"
	"time"
)

func Print_odd(i int){
	j := i * 10 - 9
	for j < i * 10{
		fmt.Println(j)
		j += 2
	}
}

func Print_even(i int){
	j := (i - 1) * 10
	for j < i * 10{
		fmt.Println(j)
		j += 2
	}
}
func main(){
	for i := 1; i <= 10; i++{
		go Print_even(i)
		go Print_odd(i)
		time.Sleep(time.Second * 1)
	}
	return
}
