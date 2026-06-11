package main

import "fmt"

func main() {
	// 1. VARIABEL
	// Deklarasi dengan tipe eksplisit
	var nama string = "Budi"
	var umur int = 25
	var tinggi float64 = 170.5
	var menikah bool = false

	// Deklarasi singkat (type inference) — hanya di dalam fungsi
	pekerjaan := "Programmer"
	kota := "Jakarta"

	fmt.Println("=== VARIABEL ===")
	fmt.Println("Nama:", nama, "| Umur:", umur, "| Tinggi:", tinggi, "| Menikah:", menikah)
	fmt.Println("Pekerjaan:", pekerjaan, "| Kota:", kota)

	// 2. PERCABANGAN (if-else)

	fmt.Println("\n PERCABANGAN")

	nilai := 85

	if nilai >= 90 {
		fmt.Println("Grade: A")
	} else if nilai >= 75 {
		fmt.Println("Grade: B")
	} else if nilai >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// if dengan short statement
	if jumlah := umur + 5; jumlah > 30 {
		fmt.Println("Umur + 5 lebih dari 30")
	}

	// 3. PERULANGAN (for)
	
	fmt.Println("\n PERULANGAN")

	// for seperti C-style
	fmt.Println("For style 1 (C-style):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iterasi ke-%d\n", i)
	}

	// for seperti while
	fmt.Println("\nFor style 2 (while-like):")
	hitungan := 0
	for hitungan < 3 {
		fmt.Printf("Hitungan: %d\n", hitungan)
		hitungan++
	}

	// for range (mirip foreach)
	fmt.Println("\nFor style 3 (range/foreach):")
	buah := []string{"Apel", "Mangga", "Jeruk"}
	for index, value := range buah {
		fmt.Printf("Index %d: %s\n", index, value)
	}

	// break & continue
	fmt.Println("\nBreak & Continue:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip bilangan genap
		}
		if i > 7 {
			break // berhenti di 7
		}
		fmt.Printf("Bilangan ganjil: %d\n", i)
	}
}
