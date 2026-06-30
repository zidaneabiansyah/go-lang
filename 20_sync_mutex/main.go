package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("MASALAH TANPA MUTEX (DATA RACE)")
	fmt.Println("---")

	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DATA RACE: multiple goroutine akses bareng
		}()
	}
	wg.Wait()
	fmt.Printf("Hasil tanpa mutex (bisa salah): %d (expected 1000)\n\n", counter)

	fmt.Println("SOLUSI DENGAN sync.Mutex")
	fmt.Println("---")

	counter2 := 0
	var mu sync.Mutex
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter2++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Hasil dengan mutex (benar): %d\n\n", counter2)

	fmt.Println("MUTEX UNTUK PROTEKSI MAP")
	fmt.Println("---")

	saldo := map[string]int{
		"alice": 1000,
		"bob":   2000,
	}
	var mapMu sync.Mutex

	var wg3 sync.WaitGroup

	// Goroutine: tambah saldo alice
	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mapMu.Lock()
			saldo["alice"] += 100
			mapMu.Unlock()
		}()
	}

	// Goroutine: tambah saldo bob
	for i := 0; i < 3; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mapMu.Lock()
			saldo["bob"] += 200
			mapMu.Unlock()
		}()
	}

	wg3.Wait()
	fmt.Printf("Saldo Alice: %d\n", saldo["alice"])
	fmt.Printf("Saldo Bob: %d\n\n", saldo["bob"])

	fmt.Println("sync.RWMutex (READ-WRITE MUTEX)")
	fmt.Println("---")

	type Config struct {
		data map[string]string
		rw   sync.RWMutex
	}

	config := &Config{
		data: map[string]string{
			"app_name": "BelajarGo",
			"version":  "1.0.0",
			"env":      "development",
		},
	}

	var wg4 sync.WaitGroup

	// Reader: bisa dibaca barengan (tidak saling block)
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			config.rw.RLock()
			defer config.rw.RUnlock()
			fmt.Printf("  Reader %d: app_name=%s, version=%s\n",
				id, config.data["app_name"], config.data["version"])
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	// Writer: harus tunggu reader selesai
	wg4.Add(1)
	go func() {
		defer wg4.Done()
		config.rw.Lock()
		defer config.rw.Unlock()
		config.data["version"] = "1.1.0"
		fmt.Println("  Writer: version updated ke 1.1.0")
	}()

	// Reader lagi
	for i := 0; i < 3; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			config.rw.RLock()
			defer config.rw.RUnlock()
			fmt.Printf("  Reader %d (post-update): version=%s\n",
				id, config.data["version"])
		}(i)
	}

	wg4.Wait()
	fmt.Println()

	fmt.Println("PATTERN: MUTEX DENGAN METHOD")
	fmt.Println("---")

	wallet := NewWallet("Budi")
	wallet.Deposit(1000)
	wallet.Deposit(500)
	wallet.Withdraw(300)
	fmt.Printf("Saldo %s: Rp%d\n\n", wallet.owner, wallet.Balance())

	fmt.Println("PATTERN: MUTEX DENGAN DEFER")
	fmt.Println("---")

	cache := &SafeCache{items: make(map[string]string)}

	var wg5 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()
			key := fmt.Sprintf("key-%d", id%3)
			cache.Set(key, fmt.Sprintf("value-%d", id))
			val, ok := cache.Get(key)
			if ok {
				fmt.Printf("  Worker %d: %s = %s\n", id, key, val)
			}
		}(i)
	}

	wg5.Wait()
	fmt.Printf("Cache size: %d\n\n", len(cache.items))

	fmt.Println("RACE CONDITION DEMO")
	fmt.Println("---")

	fmt.Println("Jalankan dengan: go run -race ./20_sync_mutex/")
	fmt.Println("Flag -race akan mendeteksi data race secara otomatis")
}

// Wallet - contoh penggunaan mutex dalam struct
type Wallet struct {
	owner  string
	saldo  int
	mu     sync.Mutex
}

func NewWallet(owner string) *Wallet {
	return &Wallet{owner: owner}
}

func (w *Wallet) Deposit(amount int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.saldo += amount
	fmt.Printf("  Deposit Rp%d -> saldo Rp%d\n", amount, w.saldo)
}

func (w *Wallet) Withdraw(amount int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.saldo < amount {
		return fmt.Errorf("saldo tidak cukup: Rp%d < Rp%d", w.saldo, amount)
	}

	w.saldo -= amount
	fmt.Printf("  Withdraw Rp%d -> saldo Rp%d\n", amount, w.saldo)
	return nil
}

func (w *Wallet) Balance() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.saldo
}

// SafeCache - contoh thread-safe cache
type SafeCache struct {
	items map[string]string
	mu    sync.RWMutex
}

func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.items[key]
	return val, ok
}

func (c *SafeCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *SafeCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
