package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(name string, delay time.Duration) string {
	time.Sleep(delay)
	return fmt.Sprintf("%s selesai (%v)", name, delay)
}

func slowOperationWithContext(ctx context.Context, name string, delay time.Duration) (string, error) {
	select {
	case <-time.After(delay):
		return fmt.Sprintf("%s selesai (%v)", name, delay), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	fmt.Println("CONTEXT BACKGROUND & TODO")
	ctx := context.Background()
	fmt.Printf("Background: %v\n", ctx)

	todo := context.TODO()
	fmt.Printf("TODO: %v\n", todo)

	fmt.Println()
	fmt.Println("CONTEXT WITH CANCEL")
	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("goroutine lain: batalkan context!")
		cancel()
	}()

	result, err := slowOperationWithContext(ctx, "tugas berat", 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("CONTEXT WITH TIMEOUT")
	fmt.Println()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel2()

	result, err = slowOperationWithContext(ctx2, "tugas timeout", 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("CONTEXT WITH DEADLINE")
	fmt.Println()

	deadline := time.Now().Add(200 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()

	result, err = slowOperationWithContext(ctx3, "tugas deadline", 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("CONTEXT WITH VALUE")
	fmt.Println()

	ctx4 := context.WithValue(context.Background(), "user_id", 123)
	ctx4 = context.WithValue(ctx4, "role", "admin")
	processRequest(ctx4)

	fmt.Println()
	fmt.Println("CONTEXT BERANTAI (parent-child)")
	fmt.Println()

	parentCtx, parentCancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer parentCancel()

	childCtx, childCancel := context.WithTimeout(parentCtx, 200*time.Millisecond)
	defer childCancel()

	result, err = slowOperationWithContext(childCtx, "child", 300*time.Millisecond)
	if err != nil {
		fmt.Printf("Child error: %s\n", err)
	} else {
		fmt.Println(result)
	}

	result, err = slowOperationWithContext(parentCtx, "parent", 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Parent error: %s\n", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("CONTOH GRACEFUL CANCELLATION")
	fmt.Println()

	ctx5, cancel5 := context.WithCancel(context.Background())

	go worker(ctx5, "worker-1")
	go worker(ctx5, "worker-2")

	time.Sleep(150 * time.Millisecond)
	fmt.Println("main: stop semua worker!")
	cancel5()

	time.Sleep(100 * time.Millisecond)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("user_id")
	role := ctx.Value("role")
	fmt.Printf("Memproses request: user_id=%v, role=%v\n", userID, role)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: berhenti karena %s\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: bekerja...\n", name)
			time.Sleep(50 * time.Millisecond)
		}
	}
}
