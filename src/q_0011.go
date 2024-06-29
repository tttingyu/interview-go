package main

import (
	"fmt"
	"sync"
)

func main() {
	print1 := make(chan bool)
	print2 := make(chan bool)
	print3 := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-print1
			fmt.Print("1")
			print2 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-print2
			fmt.Print("2")
			print3 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-print3
			fmt.Print("3\n")
			if i < 9 { // 最后一次循环不再发送到print1
				print1 <- true
			}
		}
	}()

	print1 <- true // 启动第一轮循环

	wg.Wait()
	close(print1)
	close(print2)
	close(print3)
}
