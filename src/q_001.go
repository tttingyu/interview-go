package main

//交替打印数字字母
import (
	"fmt"
	"sync"
)

func main() {
	letter := make(chan bool)
	number := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 1
		for {
			fmt.Print(i)
			i++
			<-letter
			if i > 27 { // 假设处理到 27
				break
			}
			number <- true
		}
		close(number)
	}()

	go func() {
		defer wg.Done()
		str := "abcdefghigklmnopqrstuvwxyz"
		for _, char := range str {
			letter <- true
			fmt.Printf("%c", char)
			<-number
		}
		close(letter) // 关闭 number channel，告知第一个 goroutine 不会再有数据接收
	}()

	wg.Wait()
}
