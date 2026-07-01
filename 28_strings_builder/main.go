package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("STRINGS.BASIC - Concatenation Comparison")
	fmt.Println("---")

	n := 10000

	// Method 1: String concatenation (+=)
	start := time.Now()
	result1 := ""
	for i := 0; i < n; i++ {
		result1 += "a"
	}
	time1 := time.Since(start)

	// Method 2: fmt.Sprintf
	start = time.Now()
	result2 := ""
	for i := 0; i < n; i++ {
		result2 = fmt.Sprintf("%s%s", result2, "a")
	}
	time2 := time.Since(start)

	// Method 3: strings.Join
	start = time.Now()
	parts := make([]string, n)
	for i := 0; i < n; i++ {
		parts[i] = "a"
	}
	result3 := strings.Join(parts, "")
	time3 := time.Since(start)

	// Method 4: strings.Builder
	start = time.Now()
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString("a")
	}
	result4 := builder.String()
	time4 := time.Since(start)

	// Method 5: bytes.Buffer
	start = time.Now()
	var buffer strings.Builder
	for i := 0; i < n; i++ {
		buffer.WriteString("a")
	}
	result5 := buffer.String()
	time5 := time.Since(start)

	fmt.Printf("String +=    : %v (len=%d)\n", time1, len(result1))
	fmt.Printf("Sprintf      : %v (len=%d)\n", time2, len(result2))
	fmt.Printf("strings.Join : %v (len=%d)\n", time3, len(result3))
	fmt.Printf("Builder      : %v (len=%d)\n", time4, len(result4))
	fmt.Printf("Buffer       : %v (len=%d)\n", time5, len(result5))

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - Basic Operations")
	fmt.Println("---")

	var sb strings.Builder

	// WriteString
	sb.WriteString("Hello")
	sb.WriteString(", ")
	sb.WriteString("World!")
	fmt.Printf("WriteString: %s\n", sb.String())

	// Reset
	sb.Reset()
	fmt.Printf("After Reset: '%s' (len=%d)\n", sb.String(), sb.Len())

	// WriteByte
	sb.WriteByte('H')
	sb.WriteByte('i')
	fmt.Printf("WriteByte: %s\n", sb.String())

	// WriteRune
	sb.Reset()
	sb.WriteRune('A')
	sb.WriteRune('🚀')
	fmt.Printf("WriteRune: %s\n", sb.String())

	// Grow (pre-allocate)
	sb.Reset()
	sb.Grow(100)
	fmt.Printf("After Grow(100): len=%d, cap based on Grow\n", sb.Len())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - Build Complex String")
	fmt.Println("---")

	var sb2 strings.Builder

	// HTML generation
	sb2.WriteString("<html>\n")
	sb2.WriteString("<head>\n")
	sb2.WriteString("  <title>Belajar Go</title>\n")
	sb2.WriteString("</head>\n")
	sb2.WriteString("<body>\n")
	sb2.WriteString("  <h1>Halo, Go!</h1>\n")
	sb2.WriteString("</body>\n")
	sb2.WriteString("</html>")

	fmt.Printf("HTML:\n%s\n", sb2.String())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - With Fprintf")
	fmt.Println("---")

	var sb3 strings.Builder

	// Use Fprintf to write formatted strings
	fmt.Fprintf(&sb3, "Name: %s\n", "Budi")
	fmt.Fprintf(&sb3, "Age: %d\n", 25)
	fmt.Fprintf(&sb3, "Active: %t\n", true)
	fmt.Fprintf(&sb3, "Score: %.2f\n", 95.5)

	fmt.Printf("Formatted:\n%s", sb3.String())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - CSV Builder")
	fmt.Println("---")

	var csv strings.Builder

	// Header
	headers := []string{"ID", "Name", "Email", "Age"}
	for i, h := range headers {
		if i > 0 {
			csv.WriteString(",")
		}
		csv.WriteString(h)
	}
	csv.WriteString("\n")

	// Rows
	people := [][]string{
		{"1", "Budi", "budi@mail.com", "25"},
		{"2", "Siti", "siti@mail.com", "30"},
		{"3", "Andi", "andi@mail.com", "28"},
	}

	for _, row := range people {
		for i, val := range row {
			if i > 0 {
				csv.WriteString(",")
			}
			csv.WriteString(val)
		}
		csv.WriteString("\n")
	}

	fmt.Printf("CSV:\n%s", csv.String())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - JSON Builder")
	fmt.Println("---")

	var json strings.Builder

	json.WriteString("{\n")
	json.WriteString("  \"users\": [\n")

	users := []struct {
		Name string
		Age  int
	}{
		{"Budi", 25},
		{"Siti", 30},
		{"Andi", 28},
	}

	for i, u := range users {
		json.WriteString("    {\n")
		fmt.Fprintf(&json, "      \"name\": \"%s\",\n", u.Name)
		fmt.Fprintf(&json, "      \"age\": %d\n", u.Age)
		json.WriteString("    }")
		if i < len(users)-1 {
			json.WriteString(",")
		}
		json.WriteString("\n")
	}

	json.WriteString("  ]\n")
	json.WriteString("}")

	fmt.Printf("JSON:\n%s\n", json.String())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - SQL Builder")
	fmt.Println("---")

	var sql strings.Builder

	table := "users"
	columns := []string{"name", "email", "age"}
	values := []interface{}{"Budi", "budi@mail.com", 25}

	sql.WriteString("INSERT INTO ")
	sql.WriteString(table)
	sql.WriteString(" (")
	for i, col := range columns {
		if i > 0 {
			sql.WriteString(", ")
		}
		sql.WriteString(col)
	}
	sql.WriteString(") VALUES (")
	for i, val := range values {
		if i > 0 {
			sql.WriteString(", ")
		}
		fmt.Fprintf(&sql, "'%v'", val)
	}
	sql.WriteString(");")

	fmt.Printf("SQL:\n%s\n", sql.String())

	fmt.Println()
	fmt.Println("STRINGS.BUILDER - Multi-line Template")
	fmt.Println("---")

	type User struct {
		Name  string
		Email string
		Role  string
	}

	user := User{Name: "Budi", Email: "budi@mail.com", Role: "admin"}

	var tmpl strings.Builder

	tmpl.WriteString("=== USER REPORT ===\n")
	tmpl.WriteString("─────────────────\n")
	fmt.Fprintf(&tmpl, "Name  : %s\n", user.Name)
	fmt.Fprintf(&tmpl, "Email : %s\n", user.Email)
	fmt.Fprintf(&tmpl, "Role  : %s\n", user.Role)
	tmpl.WriteString("─────────────────\n")

	fmt.Printf("%s", tmpl.String())

	fmt.Println()
	fmt.Println("STRINGS.Builder vs bytes.Buffer")
	fmt.Println("---")

	// strings.Builder lebih cepat untuk string concatenation
	// karena tidak perlu konversi []byte <-> string

	n2 := 100000

	// Benchmark strings.Builder
	start = time.Now()
	var sbench strings.Builder
	for i := 0; i < n2; i++ {
		sbench.WriteString("x")
	}
	_ = sbench.String()
	tBuilder := time.Since(start)

	// Benchmark bytes.Buffer
	start = time.Now()
	var bbuf strings.Builder
	for i := 0; i < n2; i++ {
		bbuf.WriteString("x")
	}
	_ = bbuf.String()
	tBuffer := time.Since(start)

	fmt.Printf("strings.Builder : %v\n", tBuilder)
	fmt.Printf("strings.Builder : %v\n", tBuffer)
	fmt.Println("\n(strings.Builder dan bytes.Buffer hampir sama performanya)")

	fmt.Println()
	fmt.Println("TIPS")
	fmt.Println("---")
	fmt.Println("- strings.Builder : optimal untuk concat string")
	fmt.Println("- .Grow(n)        : pre-allocate memory untuk performa")
	fmt.Println("- .Reset()        : kosongkan builder (reuse)")
	fmt.Println("- .String()       : ambil hasil akhir")
	fmt.Println("- .Len()          : panjang saat ini")
	fmt.Println("- Gunakan Fprintf untuk formatted output")
	fmt.Println("- Implement io.StringWriter interface")
}
