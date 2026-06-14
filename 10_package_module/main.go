package main

import (
	"fmt"
	"os"

	"belajar-golang/10_package_module/greeting"
	"belajar-golang/10_package_module/mathutils"
)

func init() {
	fmt.Println("init di main.go jalan pertama kali")
}

func init() {
	fmt.Println("init kedua — bisa lebih dari satu init")
}

func main() {
	fmt.Println("=== PACKAGE & MODULE ===")
	fmt.Println()

	greeting.Hello("Budi")
	greeting.Goodbye("Budi")

	fmt.Println()

	calc := mathutils.NewCalculator()
	r1 := calc.Add(10, 20, 30)
	r2 := calc.Add(5, 15)
	fmt.Printf("Hasil 1: %d\n", r1)
	fmt.Printf("Hasil 2: %d\n", r2)
	fmt.Printf("Riwayat: %v\n", calc.History)
	fmt.Printf("Total kalkulasi: %d\n", mathutils.GetInternalCounter())

	fmt.Println()
	fmt.Println("Mengakses environment variable:")
	fmt.Println("USER:", os.Getenv("USER"))
	fmt.Println("HOME:", os.Getenv("HOME"))

	fmt.Println()
	fmt.Println("Catatan:")
	fmt.Println("- Nama function yg diawali huruf besar = exported (public)")
	fmt.Println("- Nama function yg diawali huruf kecil = unexported (private)")
	fmt.Println("- Sama berlaku untuk variable, struct, field, dll")
	fmt.Println("- init() jalan otomatis sebelum main(), bisa lebih dari satu")
}
