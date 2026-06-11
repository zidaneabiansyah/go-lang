package main

import (
	"fmt"
	"strings"
)

// 1. FUNGSI SEDERHANA
// Fungsi tanpa parameter dan tanpa return
func sapa() {
	fmt.Println("Halo, selamat belajar Go!")
}

// Fungsi dengan parameter
func sapaOrang(nama string) {
	fmt.Printf("Halo %s, selamat belajar Go!\n", nama)
}

// Fungsi dengan parameter dan return value
func tambah(a int, b int) int {
	return a + b
}

// Tipe parameter sama bisa diringkas
func kurang(a, b int) int {
	return a - b
}

// 2. MULTIPLE RETURN VALUE

// Fungsi mengembalikan dua nilai
func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// 3. NAMED RETURN VALUE

// Return value sudah diberi nama
func hitungLingkaran(r float64) (luas, keliling float64) {
	luas = 3.14 * r * r
	keliling = 2 * 3.14 * r
	return // return kosong, otomatis mengembalikan luas dan keliling
}

// 4. VARIADIC FUNCTION

// ...int berarti bisa menerima 0 atau lebih argumen int
func jumlahAngka(angka ...int) int {
	total := 0
	for _, val := range angka {
		total += val
	}
	return total
}

// 5. FUNCTION SEBAGAI VALUE

// Fungsi bisa disimpan di variabel
func operasiMatematika(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 6. CLOSURE

func buatCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

// 7. METHOD

type Orang struct {
	Nama string
	Umur int
}

// Method dengan value receiver
func (o Orang) perkenalan() {
	fmt.Printf("Halo, nama saya %s, umur %d tahun\n", o.Nama, o.Umur)
}

// Method dengan pointer receiver (bisa mengubah data)
func (o *Orang) ulangTahun() {
	o.Umur++
}

// MAIN
func main() {
	fmt.Println("FUNGSI SEDERHANA")
	sapa()
	sapaOrang("Budi")

	hasilTambah := tambah(10, 5)
	fmt.Println("10 + 5 =", hasilTambah)

	hasilKurang := kurang(10, 5)
	fmt.Println("10 - 5 =", hasilKurang)

	fmt.Println("\n MULTIPLE RETURN VALUE")
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", hasil)
	}

	_, err = bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n NAMED RETURN VALUE")
	l, k := hitungLingkaran(7)
	fmt.Printf("Lingkaran r=7: Luas=%.2f, Keliling=%.2f\n", l, k)

	fmt.Println("\n VARIADIC FUNCTION ")
	fmt.Println("Jumlah:", jumlahAngka(1, 2, 3, 4, 5))
	fmt.Println("Jumlah (kosong):", jumlahAngka())

	fmt.Println("\n FUNCTION SEBAGAI VALUE ")
	// Fungsi sebagai argumen
	hasilKali := operasiMatematika(6, 7, func(a, b int) int {
		return a * b
	})
	fmt.Println("6 * 7 =", hasilKali)

	// Fungsi disimpan di variabel
	kali := func(a, b int) int { return a * b }
	fmt.Println("4 * 5 =", kali(4, 5))

	fmt.Println("\nCLOSURE ")
	counter1 := buatCounter()
	fmt.Println("Counter 1:", counter1())
	fmt.Println("Counter 1:", counter1())
	counter2 := buatCounter()
	fmt.Println("Counter 2:", counter2())
	fmt.Println("Counter 1 lagi:", counter1())

	fmt.Println("\n METHOD ")
	orang := Orang{Nama: "Budi", Umur: 25}
	orang.perkenalan()
	orang.ulangTahun()
	orang.perkenalan()

	fmt.Println("\n FUNGSI BUILT-IN YANG SERING DIPAKAI")
	teks := "  Hello, Go!  "
	fmt.Println("len:", len(teks))
	fmt.Println("strings.TrimSpace:", "'"+strings.TrimSpace(teks)+"'")
	fmt.Println("strings.ToUpper:", strings.ToUpper(teks))
	fmt.Println("strings.Contains:", strings.Contains(teks, "Go"))
}
