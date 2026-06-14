package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("TIME")
	fmt.Println()

	now := time.Now()
	fmt.Printf("Sekarang: %v\n", now)
	fmt.Printf("Tahun: %d, Bulan: %s, Tanggal: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Jam: %d, Menit: %d, Detik: %d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Hari: %s\n", now.Weekday())
	fmt.Printf("Unix timestamp: %d\n", now.Unix())

	fmt.Println()
	fmt.Println("FORMAT & PARSE TIME")
	fmt.Println()

	layout := "2006-01-02 15:04:05"
	formatted := now.Format(layout)
	fmt.Printf("Formatted: %s\n", formatted)

	parsed, _ := time.Parse(layout, "2024-12-25 10:30:00")
	fmt.Printf("Parsed: %v\n", parsed)

	customFormat := now.Format("Monday, 02 Jan 2006")
	fmt.Printf("Custom: %s\n", customFormat)

	fmt.Println()
	fmt.Println("DURATION")
	fmt.Println()

	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	duration := time.Since(start)
	fmt.Printf("Durasi: %v\n", duration)
	fmt.Printf("Milidetik: %d\n", duration.Milliseconds())

	timeout := 5 * time.Second
	fmt.Printf("Timeout: %v\n", timeout)

	d, _ := time.ParseDuration("2h30m")
	fmt.Printf("Parsed duration: %v = %d menit\n", d, int(d.Minutes()))

	deadline := time.Now().Add(30 * time.Minute)
	fmt.Printf("Deadline: %v\n", deadline)

	fmt.Println()
	fmt.Println("TIMER & TICKER")
	fmt.Println()

	timer := time.NewTimer(100 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer selesai")

	ticker := time.NewTicker(50 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Tick %d: %s\n", i, <-ticker.C)
		}
		done <- true
	}()
	<-done
	ticker.Stop()

	fmt.Println()
	fmt.Println("PERBANDINGAN WAKTU")
	fmt.Println()

	t1 := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("t1: %v\n", t1)
	fmt.Printf("t2: %v\n", t2)
	fmt.Printf("t1 before t2: %v\n", t1.Before(t2))
	fmt.Printf("t1 after t2: %v\n", t1.After(t2))
	fmt.Printf("t1 equal t2: %v\n", t1.Equal(t2))
	fmt.Printf("Selisih: %v\n", t2.Sub(t1))

	fmt.Println()
	fmt.Println("STRINGS")
	fmt.Println()

	s := "  Halo, Selamat Belajar Golang!  "
	fmt.Printf("Original: '%s'\n", s)
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(s))
	fmt.Printf("ToLower: %s\n", strings.ToLower(s))
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(s))
	fmt.Printf("Contains 'Belajar': %v\n", strings.Contains(s, "Belajar"))
	fmt.Printf("HasPrefix '  Halo': %v\n", strings.HasPrefix(s, "  Halo"))
	fmt.Printf("HasSuffix 'Golang!  ': %v\n", strings.HasSuffix(s, "Golang!  "))

	parts := strings.Split("apple,banana,grape", ",")
	fmt.Printf("Split: %v\n", parts)

	joined := strings.Join([]string{"go", "rust", "python"}, ", ")
	fmt.Printf("Join: %s\n", joined)

	replaced := strings.ReplaceAll("foo bar foo baz", "foo", "xxx")
	fmt.Printf("ReplaceAll: %s\n", replaced)

	trimmed := strings.Trim("!!!hello!!!", "!")
	fmt.Printf("Trim: '%s'\n", trimmed)

	fmt.Println()
	fmt.Println("STRCONV")
	fmt.Println()

	num := 255
	str := strconv.Itoa(num)
	fmt.Printf("Itoa: %s\n", str)

	val, _ := strconv.Atoi("1024")
	fmt.Printf("Atoi: %d\n", val)

	f := 3.14159
	fStr := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("FormatFloat: %s\n", fStr)

	parsedF, _ := strconv.ParseFloat("2.71828", 64)
	fmt.Printf("ParseFloat: %.5f\n", parsedF)

	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool: %v\n", b)

	fmt.Printf("FormatBool: %s\n", strconv.FormatBool(false))

	fmt.Println()
	fmt.Println("BENCHMARK SEDERHANA")
	fmt.Println()

	inputs := []string{"42", "100", "255", "1024", "99999"}
	start = time.Now()
	for i := 0; i < 100000; i++ {
		for _, input := range inputs {
			strconv.Atoi(input)
		}
	}
	fmt.Printf("100k x Atoi: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < 100000; i++ {
		for _, input := range inputs {
			fmt.Sscanf(input, "%d", new(int))
		}
	}
	fmt.Printf("100k x Sscanf: %v\n", time.Since(start))
}
