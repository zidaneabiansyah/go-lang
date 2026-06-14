package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Inventory struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("FILE I/O")

	data := []byte("halo ini file baris 1\nini baris kedua\n")
	err := os.WriteFile("example.txt", data, 0644)
	check(err)
	fmt.Println("file berhasil ditulis")

	read, err := os.ReadFile("example.txt")
	check(err)
	fmt.Printf("isi file:\n%s\n", read)

	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	_, err = file.WriteString("baris ketiga (append)\n")
	check(err)
	file.Close()
	fmt.Println("append berhasil")

	read2, _ := os.ReadFile("example.txt")
	fmt.Printf("setelah append:\n%s\n", read2)

	fmt.Println()
	fmt.Println("JSON MARSHAL (encode Go -> JSON)")

	budi := Person{Name: "Budi", Age: 25, Email: "budi@mail.com"}
	jsonData, err := json.Marshal(budi)
	check(err)
	fmt.Printf("JSON: %s\n", jsonData)

	jsonIndent, err := json.MarshalIndent(budi, "", "  ")
	check(err)
	fmt.Printf("JSON indented:\n%s\n", jsonIndent)

	fmt.Println()
	fmt.Println("JSON UNMARSHAL (decode JSON -> Go)")

	jsonStr := `{"name":"Siti","age":30,"email":"siti@mail.com"}`
	var siti Person
	err = json.Unmarshal([]byte(jsonStr), &siti)
	check(err)
	fmt.Printf("Person: %+v\n", siti)

	unknownJSON := `{"name":"Andi","age":28,"email":"andi@mail.com","phone":"0812345"}`
	var andi Person
	json.Unmarshal([]byte(unknownJSON), &andi)
	fmt.Printf("Andi (extra field ignored): %+v\n", andi)

	fmt.Println()
	fmt.Println("JSON ARRAY")

	peopleJSON := `[
		{"name":"Alice","age":22,"email":"alice@mail.com"},
		{"name":"Bob","age":27,"email":"bob@mail.com"}
	]`
	var people []Person
	json.Unmarshal([]byte(peopleJSON), &people)
	fmt.Printf("people: %+v\n", people)

	for _, p := range people {
		fmt.Printf("- %s, %d tahun, %s\n", p.Name, p.Age, p.Email)
	}

	fmt.Println()
	fmt.Println("JSON KE FILE")

	products := []Product{
		{ID: 1, Name: "Laptop", Price: 15000000, Stock: 10},
		{ID: 2, Name: "Mouse", Price: 250000, Stock: 50},
		{ID: 3, Name: "Keyboard", Price: 750000, Stock: 30},
	}

	inv := Inventory{Products: products, Total: len(products)}

	jsonBytes, _ := json.MarshalIndent(inv, "", "  ")
	err = os.WriteFile("inventory.json", jsonBytes, 0644)
	check(err)
	fmt.Println("inventory.json berhasil dibuat")

	fmt.Println()
	fmt.Println("BACA JSON DARI FILE")

	fileBytes, err := os.ReadFile("inventory.json")
	check(err)

	var loaded Inventory
	json.Unmarshal(fileBytes, &loaded)
	fmt.Printf("Total produk: %d\n", loaded.Total)
	for _, p := range loaded.Products {
		fmt.Printf("  %d. %s - Rp%.0f (stok: %d)\n", p.ID, p.Name, p.Price, p.Stock)
	}

	fmt.Println()
	fmt.Println("STREAMING JSON (encoder/decoder)")

	f, err := os.Create("stream.json")
	check(err)
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")

	for _, p := range products {
		err := encoder.Encode(p)
		check(err)
	}
	fmt.Println("stream.json berhasil — tiap baris 1 JSON object")

	f2, err := os.Open("stream.json")
	check(err)
	defer f2.Close()

	decoder := json.NewDecoder(f2)
	fmt.Println("baca stream.json:")
	for {
		var prod Product
		err := decoder.Decode(&prod)
		if err != nil {
			break
		}
		fmt.Printf("  %+v\n", prod)
	}

	fmt.Println()
	fmt.Println("MAP KE JSON / JSON KE MAP")

	dataMap := map[string]interface{}{
		"nama":   "Eko",
		"umur":   35,
		"hobi":   []string{"coding", "membaca"},
		"aktif":  true,
		"saldo":  50000.5,
	}
	mapJSON, _ := json.MarshalIndent(dataMap, "", "  ")
	fmt.Printf("map -> JSON:\n%s\n", mapJSON)

	arbitraryJSON := `{"nama":"Rudi","nilai":95,"lulus":true}`
	var result map[string]interface{}
	json.Unmarshal([]byte(arbitraryJSON), &result)
	fmt.Printf("JSON -> map: %+v\n", result)
	for k, v := range result {
		fmt.Printf("  %s: %v (tipe: %T)\n", k, v, v)
	}

	err = os.Remove("example.txt")
	check(err)
	err = os.Remove("inventory.json")
	check(err)
	err = os.Remove("stream.json")
	check(err)
}
