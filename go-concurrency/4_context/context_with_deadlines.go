package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performTaskByExecedingTheDeadline(ctx)
	go performTaskByCanclingTheContext(ctx, cancel)

	select {
	case <-ctx.Done():
		if err := ctx.Err(); err == context.DeadlineExceeded {
			fmt.Println("Task timed out")
		} else if err == context.Canceled {
			fmt.Println("Task canceled")
		}

	}
}

func performTaskByExecedingTheDeadline(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}

func performTaskByCanclingTheContext(ctx context.Context) {
	time.Sleep(1 * time.Second)
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}
