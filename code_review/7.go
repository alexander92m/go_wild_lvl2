// Что выведет программа? Объяснить вывод программы.
package main

import (
	"fmt"
	"math/rand"
	"time"
)
//функция передает по небуферизованному каналу инты
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) *
			time.Millisecond)
		}
		close(c)
	}()
	return c
}
//функция передает инты из 2 каналов в 3й
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
				case v := <-a:
				c <- v
				case v := <-b:
				c <- v
			}
		}
	}()
	return c
}
//т.к. канал c не закрывается после передачи, то range c получает значение по умолчанию канала, т.е. int 0 
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}