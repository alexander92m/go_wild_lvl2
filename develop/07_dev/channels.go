package main

import (
	"fmt"
	"time"
)
//or сливает каналы в 1 до получения сигнала по одному из них
func or(channels ...<- chan interface{}) <- chan interface{} {
	out := make(chan interface{})
	stop := make(chan interface{}, len(channels))
	done := make(chan interface{})
	
	for _, ch := range channels {
    	go func(ch <- chan interface{}) {
			for {
				var t interface{}
				select {
				case <-stop :
					return
				case t = <-ch :
					done<- t
					out <- t
				default :
					continue
				}
			}
		}(ch)
	}
	<-done
	close(stop)
    return out
}

func main() {
    sig := func(after time.Duration) <- chan interface{} {
        c := make(chan interface{})
        go func() {
            defer close(c)
            time.Sleep(after)
    	}()
    	return c
	}
	start := time.Now()
	<- or (	
		sig(5*time.Second),
		sig(4*time.Second),
		sig(3*time.Second),
		sig(2*time.Second),
		sig(6*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
	fmt.Println(1)

}