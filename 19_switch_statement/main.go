package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("SWITCH BASIC")
	fmt.Println("---")

	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Hari Senin, mulai kerja!")
	case "Selasa":
		fmt.Println("Hari Selasa")
	case "Rabu":
		fmt.Println("Hari Rabu")
	case "Kamis":
		fmt.Println("Hari Kamis")
	case "Jumat":
		fmt.Println("Hari Jumat, weekend sebentar lagi!")
	case "Sabtu", "Minggu": // multiple values dalam satu case
		fmt.Println("Weekend! Waktu istirahat!")
	default:
		fmt.Println("Hari gak dikenal")
	}

	fmt.Println()
	fmt.Println("SWITCH TANPA EXPRESSION (mirip if-else)")
	fmt.Println("---")

	nilai := 85

	switch {
	case nilai >= 90:
		fmt.Println("Grade: A")
	case nilai >= 80:
		fmt.Println("Grade: B")
	case nilai >= 70:
		fmt.Println("Grade: C")
	case nilai >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	fmt.Println()
	fmt.Println("SWITCH DENGAN SHORT STATEMENT")
	fmt.Println("---")

	switch jam := time.Now().Hour(); {
	case jam < 6:
		fmt.Printf("Jam %d: Dini hari\n", jam)
	case jam < 12:
		fmt.Printf("Jam %d: Pagi\n", jam)
	case jam < 17:
		fmt.Printf("Jam %d: Siang\n", jam)
	case jam < 21:
		fmt.Printf("Jam %d: Sore\n", jam)
	default:
		fmt.Printf("Jam %d: Malam\n", jam)
	}

	fmt.Println()
	fmt.Println("SWITCH DENGAN TYPE ASSERTION")
	fmt.Println("---")

	data := []interface{}{42, "halo", 3.14, true, nil}

	for _, val := range data {
		switch v := val.(type) {
		case int:
			fmt.Printf("int: %d (dua kali lipat: %d)\n", v, v*2)
		case string:
			fmt.Printf("string: \"%s\" (panjang: %d)\n", v, len(v))
		case float64:
			fmt.Printf("float64: %.2f\n", v)
		case bool:
			fmt.Printf("bool: %t\n", v)
		case nil:
			fmt.Println("nil: gak ada nilai")
		default:
			fmt.Printf("tipe gak dikenal: %T\n", v)
		}
	}

	fmt.Println()
	fmt.Println("SWITCH DENGAN FALLTHROUGH")
	fmt.Println("---")

	level := 2

	switch level {
	case 1:
		fmt.Println("Level 1 - Junior")
		fallthrough // paksa lanjut ke case berikutnya TANPA cek kondisi
	case 2:
		fmt.Println("Level 2 - Mid")
		fallthrough
	case 3:
		fmt.Println("Level 3 - Senior")
	case 4:
		fmt.Println("Level 4 - Lead")
	}

	fmt.Println()
	fmt.Println("SWITCH SEBAGAI LOOKUP TABLE")
	fmt.Println("---")

	bulan := 2

	namaBulan := ""
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Invalid"
	}
	fmt.Printf("Bulan ke-%d: %s\n", bulan, namaBulan)

	fmt.Println()
	fmt.Println("SWITCH UNTUK STATUS CODE")
	fmt.Println("---")

	statusCode := 404

	switch statusCode {
	case 200:
		fmt.Println("OK - Request berhasil")
	case 201:
		fmt.Println("Created - Resource berhasil dibuat")
	case 204:
		fmt.Println("No Content - Tidak ada content")
	case 400:
		fmt.Println("Bad Request - Request salah")
	case 401:
		fmt.Println("Unauthorized - Perlu autentikasi")
	case 403:
		fmt.Println("Forbidden - Akses ditolak")
	case 404:
		fmt.Println("Not Found - Resource gak ketemu")
	case 500:
		fmt.Println("Internal Server Error - Server error")
	default:
		fmt.Printf("Status code gak dikenal: %d\n", statusCode)
	}

	fmt.Println()
	fmt.Println("SWITCH UNTUK KONVERSI SATUAN")
	fmt.Println("---")

	satuan := "km"
	jarak := 10.0

	var hasil float64
	switch satuan {
	case "km":
		hasil = jarak * 1000
		fmt.Printf("%.0f %s = %.0f meter\n", jarak, satuan, hasil)
	case "mile":
		hasil = jarak * 1609.344
		fmt.Printf("%.0f %s = %.0f meter\n", jarak, satuan, hasil)
	case "yard":
		hasil = jarak * 0.9144
		fmt.Printf("%.0f %s = %.0f meter\n", jarak, satuan, hasil)
	case "feet":
		hasil = jarak * 0.3048
		fmt.Printf("%.0f %s = %.0f meter\n", jarak, satuan, hasil)
	default:
		fmt.Println("Satuan gak didukung")
	}

	fmt.Println()
	fmt.Println("SWITCH DENGAN MULTI-CASE & SHORTCUT")
	fmt.Println("---")

	karakter := 'A'

	switch {
	case karakter >= 'A' && karakter <= 'Z':
		fmt.Printf("'%c' adalah huruf besar\n", karakter)
	case karakter >= 'a' && karakter <= 'z':
		fmt.Printf("'%c' adalah huruf kecil\n", karakter)
	case karakter >= '0' && karakter <= '9':
		fmt.Printf("'%c' adalah angka\n", karakter)
	default:
		fmt.Printf("'%c' adalah karakter spesial\n", karakter)
	}

	fmt.Println()
	fmt.Println("TIP: Switch di Go TIDAK pernah fallthrough secara default")
	fmt.Println("(berbeda dengan C/Java yang butuh break)")
}
