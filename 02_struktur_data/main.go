package main

import "fmt"

func main() {
	// 1. ARRAY - Kumpulan data dengan ukuran tetap
	fmt.Println("ARRAY")

	// Deklarasi array
	var angka [5]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50
	fmt.Println("Array angka:", angka)

	// Inisialisasi langsung
	buah := [3]string{"Apel", "Mangga", "Jeruk"}
	fmt.Println("Array buah:", buah)
	fmt.Println("Panjang array:", len(buah))

	// Array multi dimensi
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Matrix 2x2:", matrix)

	// Iterasi array
	fmt.Println("\nIterasi array:")
	for i, val := range buah {
		fmt.Printf("Index %d: %s\n", i, val)
	}

	// 2. SLICE
	fmt.Println("\n SLICE (ukuran dinamis) ")

	// Membuat slice dari literal
	nama := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Slice nama:", nama)
	fmt.Println("Panjang:", len(nama), "Kapasitas:", cap(nama))

	// append — menambah elemen
	nama = append(nama, "Diana")
	fmt.Println("Setelah append:", nama)

	// slicing — memotong [start:end]
	sub := nama[1:3] // index 1 sampai sebelum 3
	fmt.Println("Slice [1:3]:", sub)

	// make — membuat slice dengan kapasitas tertentu
	angka2 := make([]int, 3, 5) // panjang 3, kapasitas 5
	angka2[0] = 100
	angka2[1] = 200
	angka2[2] = 300
	fmt.Println("Slice dengan make:", angka2, "len:", len(angka2), "cap:", cap(angka2))

	// copy
	kopi := make([]string, len(nama))
	copy(kopi, nama)
	fmt.Println("Copy slice:", kopi)

	// 3. MAP - Kumpulan key-value
	fmt.Println("\n MAP ")

	// Deklarasi map
	umur := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Charlie": 35,
	}
	fmt.Println("Map umur:", umur)
	fmt.Println("Umur Alice:", umur["Alice"])

	// Menambah/mengubah data
	umur["Diana"] = 28
	fmt.Println("Setelah tambah Diana:", umur)

	// Cek key exists
	value, exists := umur["Eve"]
	fmt.Printf("Eve: value=%d, exists=%t\n", value, exists)

	// Hapus data
	delete(umur, "Charlie")
	fmt.Println("Setelah hapus Charlie:", umur)

	// Iterasi map
	fmt.Println("\nIterasi map:")
	for key, val := range umur {
		fmt.Printf("%s -> %d tahun\n", key, val)
	}
}
