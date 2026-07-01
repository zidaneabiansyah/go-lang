package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("REGEXP MATCH - Cocokkan Pola")
	fmt.Println("---")

	re := regexp.MustCompile(`\d+`)

	teks := "Saya punya 3 kucing dan 2 anjing"
	fmt.Printf("Teks: %s\n", teks)
	fmt.Printf("Match? %v\n", re.MatchString(teks))

	fmt.Println()
	fmt.Println("FIND ALL - Cari Semua Match")
	fmt.Println("---")

	angka := re.FindAllString(teks, -1)
	fmt.Printf("Semua angka: %v\n", angka)

	// Batasi jumlah match
	angka2 := re.FindAllString(teks, 1)
	fmt.Printf("Max 1 match: %v\n", angka2)

	fmt.Println()
	fmt.Println("FIND INDEX - Posisi Match")
	fmt.Println("---")

	index := re.FindStringIndex(teks)
	fmt.Printf("Index pertama: %v (posisi %d s/d %d)\n", index, index[0], index[1])

	fmt.Println()
	fmt.Println("FIND ALL INDEX - Semua Posisi")
	fmt.Println("---")

	allIndex := re.FindAllStringIndex(teks, -1)
	fmt.Printf("Semua posisi: %v\n", allIndex)

	for _, idx := range allIndex {
		fmt.Printf("  '%s' di posisi %d-%d\n", teks[idx[0]:idx[1]], idx[0], idx[1])
	}

	fmt.Println()
	fmt.Println("SUBMATCH - Grup dalam Pola")
	fmt.Println("---")

	re2 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	email := "Hubungi saya di budi@gmail.com atau siti@yahoo.co.id"

	submatch := re2.FindStringSubmatch(email)
	fmt.Printf("Full match: %s\n", submatch[0])
	fmt.Printf("Username: %s\n", submatch[1])
	fmt.Printf("Domain: %s\n", submatch[2])
	fmt.Printf("TLD: %s\n", submatch[3])

	fmt.Println()
	fmt.Println("NAMED GROUP - Grup Bernama")
	fmt.Println("---")

	re3 := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
 tanggal := "Tanggal lahir: 1990-05-15, tanggal nikah: 2020-12-25"

	// Cari semua tanggal
	allDates := re3.FindAllString(tanggal, -1)
	fmt.Printf("Semua tanggal: %v\n", allDates)

	// Ambil group berdasarkan index
	submatch3 := re3.FindStringSubmatch(tanggal)
	names := re3.SubexpNames()
	for i, name := range names {
		if i > 0 {
			fmt.Printf("  %s = %s\n", name, submatch3[i])
		}
	}

	fmt.Println()
	fmt.Println("REPLACE - Ganti Teks")
	fmt.Println("---")

	teks2 := "Halo 123 dunia 456"

	// Ganti semua angka
	re4 := regexp.MustCompile(`\d+`)
	result := re4.ReplaceAllString(teks2, "[ANGKA]")
	fmt.Printf("Replace all: %s\n", result)

	// Ganti dengan fungsi
	result2 := re4.ReplaceAllStringFunc(teks2, func(s string) string {
		return fmt.Sprintf("[%s!]", s)
	})
	fmt.Printf("Replace func: %s\n", result2)

	// Ganti email dengan asterisk
	re5 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	emails := "Kontak: budi@gmail.com, siti@yahoo.com, andi@outlook.com"
	censored := re5.ReplaceAllString(emails, "***@$2.$3")
	fmt.Printf("Censored: %s\n", censored)

	fmt.Println()
	fmt.Println("SPLIT - Pecah Teks")
	fmt.Println("---")

	teks3 := "satu,dua; tiga|empat  lima"
	re6 := regexp.MustCompile(`[,;|\s]+`)
	parts := re6.Split(teks3, -1)
	fmt.Printf("Teks: %s\n", teks3)
	fmt.Printf("Hasil split: %v\n", parts)

	fmt.Println()
	fmt.Println("VALIDASI EMAIL")
	fmt.Println("---")

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	emails2 := []string{
		"budi@gmail.com",
		"siti@yahoo.co.id",
		"invalid-email",
		"user@.com",
		"test@test",
		"admin@company.org",
	}

	for _, e := range emails2 {
		valid := emailRegex.MatchString(e)
		status := "VALID"
		if !valid {
			status = "INVALID"
		}
		fmt.Printf("  %-25s -> %s\n", e, status)
	}

	fmt.Println()
	fmt.Println("VALIDASI NOMOR TELEPON")
	fmt.Println("---")

	phoneRegex := regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,10}$`)

	phones := []string{
		"081234567890",
		"+62812345678",
		"628567890123",
		"08123",
		"021123456",
		"08123456789012345",
	}

	for _, p := range phones {
		valid := phoneRegex.MatchString(p)
		status := "VALID"
		if !valid {
			status = "INVALID"
		}
		fmt.Printf("  %-20s -> %s\n", p, status)
	}

	fmt.Println()
	fmt.Println("VALIDASI URL")
	fmt.Println("---")

	urlRegex := regexp.MustCompile(`^https?://[a-zA-Z0-9\-]+\.[a-zA-Z]{2,}(/.*)?$`)

	urls := []string{
		"https://www.google.com",
		"http://example.com/path/to/page",
		"ftp://files.example.com",
		"www.google.com",
		"https://",
	}

	for _, u := range urls {
		valid := urlRegex.MatchString(u)
		status := "VALID"
		if !valid {
			status = "INVALID"
		}
		fmt.Printf("  %-40s -> %s\n", u, status)
	}

	fmt.Println()
	fmt.Println("VALIDASI IP ADDRESS")
	fmt.Println("---")

	ipRegex := regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)

	ips := []string{
		"192.168.1.1",
		"10.0.0.1",
		"255.255.255.0",
		"999.999.999.999",
		"192.168.1",
		"abc.def.ghi.jkl",
	}

	for _, ip := range ips {
		valid := ipRegex.MatchString(ip)
		status := "VALID"
		if !valid {
			status = "INVALID"
		}
		fmt.Printf("  %-20s -> %s\n", ip, status)
	}

	fmt.Println()
	fmt.Println("ESCAPED PATTERN - Pola khusus")
	fmt.Println("---")

	// Cari URL dalam teks
	text := "Kunjungi https://go.dev atau http://example.com/path?q=search"

	re7 := regexp.MustCompile(`https?://[^\s]+`)
	urls2 := re7.FindAllString(text, -1)
	fmt.Printf("URL ditemukan: %v\n", urls2)

	// Cari hashtag
	tweet := "Belajar #golang itu #fun dan #productive banget!"
	re8 := regexp.MustCompile(`#\w+`)
	hashtags := re8.FindAllString(tweet, -1)
	fmt.Printf("Hashtags: %v\n", hashtags)

	// Cari mention
	posts := "Thanks @golang dan @github untuk platformnya"
	re9 := regexp.MustCompile(`@\w+`)
	mentions := re9.FindAllString(posts, -1)
	fmt.Printf("Mentions: %v\n", mentions)

	fmt.Println()
	fmt.Println("TIPS REGEXP")
	fmt.Println("---")
	fmt.Println(`\d     = digit [0-9]`)
	fmt.Println(`\w     = word char [a-zA-Z0-9_]`)
	fmt.Println(`\s     = whitespace`)
	fmt.Println(`.      = karakter apapun`)
	fmt.Println(`+      = 1 atau lebih`)
	fmt.Println(`*      = 0 atau lebih`)
	fmt.Println(`?      = 0 atau 1`)
	fmt.Println(`{n,m}  = n sampai m kali`)
	fmt.Println(`^      = awal string`)
	fmt.Println(`$      = akhir string`)
	fmt.Println(`[]     = karakter class`)
	fmt.Println(`()     = grup/capture`)
	fmt.Println(`|      = OR`)
	fmt.Println(`\b     = word boundary`)

	_ = strings.TrimSpace("")
}
