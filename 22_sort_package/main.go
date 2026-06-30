package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("sort.Ints - Sort Integer Slice")
	fmt.Println("---")

	angka := []int{42, 17, 89, 3, 55, 26}
	fmt.Println("Sebelum:", angka)

	sort.Ints(angka)
	fmt.Println("Sesudah:", angka)

	fmt.Println()
	fmt.Println("sort.Strings - Sort String Slice")
	fmt.Println("---")

	nama := []string{"Siti", "Budi", "Andi", "Diana", "Charlie"}
	fmt.Println("Sebelum:", nama)

	sort.Strings(nama)
	fmt.Println("Sesudah:", nama)

	fmt.Println()
	fmt.Println("sort.Float64s - Sort Float Slice")
	fmt.Println("---")

	nilai := []float64{3.14, 1.41, 2.72, 0.58, 1.73}
	fmt.Println("Sebelum:", nilai)

	sort.Float64s(nilai)
	fmt.Println("Sesudah:", nilai)

	fmt.Println()
	fmt.Println("sort.Slice - Custom Sort (ascending/descending)")
	fmt.Println("---")

	umur := []int{25, 30, 18, 35, 22}
	fmt.Println("Sebelum:", umur)

	// Descending (besar ke kecil)
	sort.Slice(umur, func(i, j int) bool {
		return umur[i] > umur[j]
	})
	fmt.Println("Descending:", umur)

	// Ascending (kecil ke besar)
	sort.Slice(umur, func(i, j int) bool {
		return umur[i] < umur[j]
	})
	fmt.Println("Ascending:", umur)

	fmt.Println()
	fmt.Println("sort.Slice - Sort Struct by Field")
	fmt.Println("---")

	type Mahasiswa struct {
		Nama  string
		Umur  int
		IPK   float64
	}

	mahasiswa := []Mahasiswa{
		{"Budi", 22, 3.5},
		{"Siti", 20, 3.8},
		{"Andi", 23, 3.2},
		{"Diana", 21, 3.9},
		{"Charlie", 22, 3.5},
	}

	// Sort by Nama (ascending)
	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].Nama < mahasiswa[j].Nama
	})
	fmt.Println("Sort by Nama:")
	for _, m := range mahasiswa {
		fmt.Printf("  %s (umur: %d, IPK: %.1f)\n", m.Nama, m.Umur, m.IPK)
	}

	// Sort by IPK (descending - terbesar dulu)
	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].IPK > mahasiswa[j].IPK
	})
	fmt.Println("\nSort by IPK (terbesar dulu):")
	for _, m := range mahasiswa {
		fmt.Printf("  %s - IPK: %.1f\n", m.Nama, m.IPK)
	}

	// Sort by Umur, lalu by Nama
	sort.Slice(mahasiswa, func(i, j int) bool {
		if mahasiswa[i].Umur == mahasiswa[j].Umur {
			return mahasiswa[i].Nama < mahasiswa[j].Nama
		}
		return mahasiswa[i].Umur < mahasiswa[j].Umur
	})
	fmt.Println("\nSort by Umur, lalu Nama:")
	for _, m := range mahasiswa {
		fmt.Printf("  %s (umur: %d)\n", m.Nama, m.Umur)
	}

	fmt.Println()
	fmt.Println("sort.SliceStable - Stable Sort (pertahankan urutan awal)")
	fmt.Println("---")

	type Produk struct {
		Nama  string
		Harga int
	}

	produk := []Produk{
		{"Laptop", 15000000},
		{"Mouse", 250000},
		{"Keyboard", 750000},
		{"Monitor", 3500000},
		{"Mouse Gaming", 500000},
	}

	// Stable sort by Harga - Mouse dan Mouse Gaming tetap urutan awal
	sort.SliceStable(produk, func(i, j int) bool {
		return produk[i].Harga < produk[j].Harga
	})
	fmt.Println("Sort by Harga (stable):")
	for _, p := range produk {
		fmt.Printf("  %s - Rp%d\n", p.Nama, p.Harga)
	}

	fmt.Println()
	fmt.Println("sort.Search - Binary Search")
	fmt.Println("---")

	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	target := 50

	idx := sort.Search(len(data), func(i int) bool {
		return data[i] >= target
	})

	if idx < len(data) && data[idx] == target {
		fmt.Printf("Ditemukan %d di index %d\n", target, idx)
	} else {
		fmt.Printf("%d gak ditemukan\n", target)
	}

	target2 := 55
	idx2 := sort.Search(len(data), func(i int) bool {
		return data[i] >= target2
	})
	fmt.Printf("Index untuk %d (insertion point): %d\n", target2, idx2)

	fmt.Println()
	fmt.Println("sort.Interface - Implementasi Manual")
	fmt.Println("---")

	nama2 := []string{"Charlie", "Bob", "Alice", "David", "Eve"}
	fmt.Println("Sebelum:", nama2)

	// Sort by panjang nama
	sort.Slice(nama2, func(i, j int) bool {
		return len(nama2[i]) < len(nama2[j])
	})
	fmt.Println("Sort by panjang nama:", nama2)

	fmt.Println()
	fmt.Println("SORT DENGAN MULTI-CRITERIA")
	fmt.Println("---")

	type Karyawan struct {
		Nama     string
		Dept     string
		Gaji     int
		TahunMasuk int
	}

	karyawan := []Karyawan{
		{"Budi", "IT", 8000000, 2020},
		{"Siti", "HR", 7000000, 2019},
		{"Andi", "IT", 9000000, 2018},
		{"Diana", "Marketing", 7500000, 2021},
		{"Eko", "IT", 8500000, 2020},
		{"Fani", "HR", 7200000, 2020},
	}

	// Sort: Dept ASC, Gaji DESC
	sort.Slice(karyawan, func(i, j int) bool {
		if karyawan[i].Dept == karyawan[j].Dept {
			return karyawan[i].Gaji > karyawan[j].Gaji
		}
		return karyawan[i].Dept < karyawan[j].Dept
	})

	fmt.Println("Sort by Dept (ASC), lalu Gaji (DESC):")
	currentDept := ""
	for _, k := range karyawan {
		if k.Dept != currentDept {
			fmt.Printf("\n  [%s]\n", k.Dept)
			currentDept = k.Dept
		}
		fmt.Printf("    %s - Rp%d (%d)\n", k.Nama, k.Gaji, k.TahunMasuk)
	}

	fmt.Println()
	fmt.Println("TIPS")
	fmt.Println("---")
	fmt.Println("- sort.Ints/String/Float64s: sort ASC cepat")
	fmt.Println("- sort.Slice: custom sort, paling fleksibel")
	fmt.Println("- sort.SliceStable: pertahankan urutan asli untuk elemen equal")
	fmt.Println("- sort.Search: binary search, data HARUS sudah ter-sort")
	fmt.Println("- fmt.Println(strings.Repeat(\"-\", 30)):", strings.Repeat("-", 30))
}
