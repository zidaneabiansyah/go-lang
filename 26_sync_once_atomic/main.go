package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type LazyConfig struct {
	DBHost string
	DBPort string
}

func (c *LazyConfig) GetConfig() *LazyConfig {
	if c.DBHost == "" {
		c.DBHost = "localhost"
		c.DBPort = "5432"
	}
	return c
}

type Database struct {
	ConnStr string
}

var dbInstance *Database
var dbOnce sync.Once

func GetDatabase() *Database {
	dbOnce.Do(func() {
		dbInstance = &Database{ConnStr: "postgres://localhost:5432/mydb"}
		fmt.Println("  Database connection created (only once!)")
	})
	return dbInstance
}

func main() {
	fmt.Println("SYNC.ONCE - Jalankan Sekali Saja")
	fmt.Println("---")

	var once sync.Once
	var wg sync.WaitGroup

	// 10 goroutine coba inisialisasi
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf("  Worker %d: Inisialisasi dilakukan HANYA SEKALI!\n", id)
				time.Sleep(50 * time.Millisecond)
			})
			fmt.Printf("  Worker %d: Selesai\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("  Semua worker selesai")

	fmt.Println()
	fmt.Println("SYNC.ONCE - Lazy Initialization")
	fmt.Println("---")

	config := &LazyConfig{}

	var wg2 sync.WaitGroup

	// Simulasi multiple request yang butuh config
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			cfg := config.GetConfig()
			fmt.Printf("  Request %d: DB_HOST=%s\n", id, cfg.DBHost)
		}(i)
	}

	wg2.Wait()

	fmt.Println()
	fmt.Println("SYNC.ONCE - Thread-Safe Singleton")
	fmt.Println("---")

	var wg3 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			db := GetDatabase()
			fmt.Printf("  Worker %d: DB Connection = %s\n", id, db.ConnStr)
		}(i)
	}

	wg3.Wait()

	fmt.Println()
	fmt.Println("ATOMIC - Counter Tanpa Lock")
	fmt.Println("---")

	var counter int64 = 0
	var wg4 sync.WaitGroup

	// 1000 goroutine increment
	for i := 0; i < 1000; i++ {
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg4.Wait()
	fmt.Printf("Counter: %d (expected 1000)\n", counter)

	fmt.Println()
	fmt.Println("ATOMIC - Read & Write")
	fmt.Println("---")

	var status int64 = 0

	// Writer
	go func() {
		time.Sleep(100 * time.Millisecond)
		atomic.StoreInt64(&status, 1)
		fmt.Println("  Status updated to 1")
	}()

	// Reader (cek berkala)
	for i := 0; i < 5; i++ {
		val := atomic.LoadInt64(&status)
		fmt.Printf("  Status: %d\n", val)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("ATOMIC - Compare And Swap (CAS)")
	fmt.Println("---")

	var value int64 = 10

	// Sukses
	swapped := atomic.CompareAndSwapInt64(&value, 10, 20)
	fmt.Printf("CAS(10->20): %v, value now: %d\n", swapped, value)

	// Gagal (value sudah 20, bukan 10)
	swapped = atomic.CompareAndSwapInt64(&value, 10, 30)
	fmt.Printf("CAS(10->30): %v, value still: %d\n", swapped, value)

	fmt.Println()
	fmt.Println("ATOMIC - Increment & Decrement")
	fmt.Println("---")

	var hits int64 = 0

	var wg5 sync.WaitGroup

	// Simulasi 1000 concurrent hits
	for i := 0; i < 1000; i++ {
		wg5.Add(1)
		go func() {
			defer wg5.Done()
			atomic.AddInt64(&hits, 1)
		}()
	}

	wg5.Wait()
	fmt.Printf("Total hits: %d\n", hits)

	// Decrement
	atomic.AddInt64(&hits, -100)
	fmt.Printf("Setelah -100: %d\n", hits)

	fmt.Println()
	fmt.Println("ATOMIC - Load & Store")
	fmt.Println("---")

	var flag int32 = 0

	// Multiple reader
	var wg6 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg6.Add(1)
		go func(id int) {
			defer wg6.Done()
			for {
				if atomic.LoadInt32(&flag) == 1 {
					fmt.Printf("  Reader %d: detected flag=1, stopping\n", id)
					return
				}
				time.Sleep(30 * time.Millisecond)
			}
		}(i)
	}

	// Set flag setelah 100ms
	time.Sleep(100 * time.Millisecond)
	atomic.StoreInt32(&flag, 1)
	fmt.Println("  Main: flag set to 1")

	wg6.Wait()

	fmt.Println()
	fmt.Println("ATOMIC VALUE - Simpan Tipe Apa Saja")
	fmt.Println("---")

	var configValue atomic.Value

	// Simpan config awal
	configValue.Store(map[string]string{
		"host": "localhost",
		"port": "8080",
	})

	// Baca config
	cfg := configValue.Load().(map[string]string)
	fmt.Printf("Config awal: %v\n", cfg)

	// Update config
	configValue.Store(map[string]string{
		"host": "production.com",
		"port": "443",
	})

	cfg = configValue.Load().(map[string]string)
	fmt.Printf("Config updated: %v\n", cfg)

	fmt.Println()
	fmt.Println("PRAKTIK: SYNC.ONCE UNTUK DATABASE")
	fmt.Println("---")

	type DBPool struct {
		connections []string
		once        sync.Once
	}

	pool := &DBPool{}

	var wg7 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg7.Add(1)
		go func(id int) {
			defer wg7.Done()
			pool.once.Do(func() {
				fmt.Printf("  Worker %d: Membuat connection pool...\n", id)
				pool.connections = []string{"conn1", "conn2", "conn3"}
				time.Sleep(50 * time.Millisecond)
			})
			fmt.Printf("  Worker %d: Pool size = %d\n", id, len(pool.connections))
		}(i)
	}

	wg7.Wait()

	fmt.Println()
	fmt.Println("PRAKTIK: ATOMIC UNTUK RATE LIMITER")
	fmt.Println("---")

	var requests int64 = 0
	var wg8 sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg8.Add(1)
		go func(id int) {
			defer wg8.Done()
			current := atomic.AddInt64(&requests, 1)
			if current <= 5 {
				fmt.Printf("  Request %d: ALLOWED (count=%d)\n", id, current)
			} else {
				fmt.Printf("  Request %d: REJECTED (count=%d)\n", id, current)
			}
		}(i)
	}

	wg8.Wait()

	fmt.Println()
	fmt.Println("TIPS")
	fmt.Println("---")
	fmt.Println("- sync.Once    : Pastikan fungsi hanya jalan 1x (singleton, init)")
	fmt.Println("- atomic.Add   : Counter tanpa lock (lebih cepat dari Mutex)")
	fmt.Println("- atomic.Load  : Baca value thread-safe")
	fmt.Println("- atomic.Store : Tulis value thread-safe")
	fmt.Println("- atomic.CompareAndSwap : Lock-free update bersyarat")
	fmt.Println("- atomic.Value : Simpan tipe apa saja secara thread-safe")
}
