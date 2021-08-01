package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var ch = make(chan int, 10)					//定义全局变量：一个带缓冲的通道
var stop = false							//定义全局标志变量

func Producer(w1 *sync.WaitGroup){			//生产者
	for !stop{
		a := rand.Intn(100)
		ch <- a
		time.Sleep( 1 * time.Second)
	}
	w1.Done()
}

func Consumer(w2 *sync.WaitGroup){			//消费者
	for v := range ch{
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
	w2.Done()
}

func main(){
	var w1 sync.WaitGroup
	var w2 sync.WaitGroup
	w1.Add(5)
	w2.Add(5)
	for i := 0; i < 5; i++{
		go Producer(&w1)
		go Consumer(&w2)
	}
	time.Sleep(1 * time.Second)
	stop = true
	w1.Wait()
	close(ch)
	w2.Wait()
	return
}