package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		//default:
		//	fmt.Println("hello")
		case <-ctx.Done():
			fmt.Println("超时了")
			return ctx.Err()
		}
	}
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go worker(ctx, &wg)
	//}
	//
	//time.Sleep(4 * time.Second)
	//cancel()
	//
	////context.WithCancel()
	//
	//wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	ctx1, _ := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel1()

	fmt.Println("sjkjvdk")
	cancel()
	select {
	case <-time.After(30 * time.Second):
		fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	case <-ctx1.Done():
		fmt.Println(ctx1.Err(), "ppppppp") // prints "context deadline exceeded"

	}
	fmt.Println("djvnlau")
	<-ctx.Done()
	fmt.Println("dsnvfu")

}
