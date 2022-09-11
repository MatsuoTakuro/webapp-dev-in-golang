package ch02

import (
	"context"
	"fmt"
	"time"
)

func Sub() {
	// list_2_3()
	// list_2_5()
	list_2_6()
}

func list_2_3() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("cencelled")
		return
	}
	fmt.Println("not cencelled")
}

func list_2_5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("another goroutine") }()
	fmt.Println("stop")
	fmt.Println(ctx.Err())
	<-ctx.Done()
	fmt.Println(ctx.Err())
	fmt.Println("And the time moves on")
}

func list_2_6() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancelled")
				return
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("not cancelled")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		task <- i
	}

	cancel()
	time.Sleep(time.Second)
}
