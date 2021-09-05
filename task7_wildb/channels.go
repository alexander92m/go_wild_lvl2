package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<- chan interface{})
	or = func(channels ...<- chan interface{}) {
		for i, ch := range channels {
			fmt.Printf("tipe of channels[%d]=%T\n", i, <-ch)
			
		}
		
	}


    sig := func(after time.Duration) <- chan interface{} {
        c := make(chan interface{})
        go func() {
            defer close(c)
            time.Sleep(after)
    	}()
    	return c
	}

	start := time.Now()
	or (
		
		sig(1*time.Second),
		sig(2*time.Second),
		sig(5*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
	fmt.Println(1)

}