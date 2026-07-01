package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("IO.WRITER - Interface Penulis Data")
	fmt.Println("---")

	// strings.Builder implements io.Writer
	var builder strings.Builder
	n, err := fmt.Fprint(&builder, "Halo dari fmt.Fprint!")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Bytes written: %d\n", n)
	fmt.Printf("Result: %s\n\n", builder.String())

	fmt.Println("IO.READER - Interface Pembaca Data")
	fmt.Println("---")

	reader := strings.NewReader("Hello, io.Reader!")
	buf := make([]byte, 8)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("\n  EOF reached")
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("  Read %d bytes: %s\n", n, string(buf[:n]))
	}

	fmt.Println()
	fmt.Println("IO.COPY - Salin Data")
	fmt.Println("---")

	src := strings.NewReader("Data yang disalin dari sumber")
	dst := &bytes.Buffer{}

	written, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Written: %d bytes\n", written)
	fmt.Printf("Result: %s\n", dst.String())

	fmt.Println()
	fmt.Println("IO.COPYN - Salin dengan Limit")
	fmt.Println("---")

	src2 := strings.NewReader("Ini adalah teks panjang yang akan dipotong")
	dst2 := &bytes.Buffer{}

	written2, err := io.CopyN(dst2, src2, 10)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Written: %d bytes\n", written2)
	fmt.Printf("Result: %s\n", dst2.String())

	fmt.Println()
	fmt.Println("IO.READALL - Baca Semua Data")
	fmt.Println("---")

	reader3 := strings.NewReader("Baca semua data sekaligus")
	data, err := io.ReadAll(reader3)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Data: %s\n", string(data))

	fmt.Println()
	fmt.Println("BYTES.BUFFER - io.Reader + io.Writer")
	fmt.Println("---")

	buf2 := bytes.NewBufferString("Awal")

	// Write
	buf2.WriteString(" -> Ditambah")
	fmt.Printf("After write: %s\n", buf2.String())

	// Read
	buf3 := make([]byte, 5)
	n, _ = buf2.Read(buf3)
	fmt.Printf("Read %d bytes: %s\n", n, string(buf3))

	// Write more
	fmt.Fprintf(buf2, " (%d items)", 42)
	fmt.Printf("After fmt.Fprintf: %s\n", buf2.String())

	// Reset
	buf2.Reset()
	fmt.Printf("After reset: '%s' (len=%d)\n", buf2.String(), buf2.Len())

	fmt.Println()
	fmt.Println("MULTIREADER - Gabung Banyak Reader")
	fmt.Println("---")

	r1 := strings.NewReader("Bagian 1, ")
	r2 := strings.NewReader("Bagian 2, ")
	r3 := strings.NewReader("Bagian 3")

	multiReader := io.MultiReader(r1, r2, r3)
	result, _ := io.ReadAll(multiReader)
	fmt.Printf("Combined: %s\n", result)

	fmt.Println()
	fmt.Println("MULTIWRITER - Tulis ke Banyak Writer")
	fmt.Println("---")

	w1 := &bytes.Buffer{}
	w2 := &bytes.Buffer{}
	w3 := &bytes.Buffer{}

	multiWriter := io.MultiWriter(w1, w2, w3)
	fmt.Fprint(multiWriter, "Data ditulis ke 3 writer!")

	fmt.Printf("Writer 1: %s\n", w1.String())
	fmt.Printf("Writer 2: %s\n", w2.String())
	fmt.Printf("Writer 3: %s\n", w3.String())

	fmt.Println()
	fmt.Println("READER UNTUK FILE")
	fmt.Println("---")

	// Tulis file dulu
	content := "Baris 1\nBaris 2\nBaris 3\nBaris 4\nBaris 5"
	os.WriteFile("test_io.txt", []byte(content), 0644)

	// Baca file dengan io.Copy
	file, _ := os.Open("test_io.txt")
	defer file.Close()
	defer os.Remove("test_io.txt")

	var fileContent bytes.Buffer
	io.Copy(&fileContent, file)
	fmt.Printf("File content:\n%s\n", fileContent.String())

	fmt.Println()
	fmt.Println("IO.TEEWRITER - Tulis + Salin")
	fmt.Println("---")

	original := strings.NewReader("Data penting")
	teeBuffer := &bytes.Buffer{}

	teeReader := io.TeeReader(original, teeBuffer)

	// Baca dari teeReader, otomatis disalin ke teeBuffer
	data, _ = io.ReadAll(teeReader)
	fmt.Printf("Read: %s\n", string(data))
	fmt.Printf("Tee copy: %s\n", teeBuffer.String())

	fmt.Println()
	fmt.Println("WRITE STRING UNTUK io.Writer")
	fmt.Println("---")

	var buf4 bytes.Buffer

	// Cara 1: fmt.Fprintf
	fmt.Fprintf(&buf4, "Name: %s, Age: %d\n", "Budi", 25)

	// Cara 2: io.WriteString
	io.WriteString(&buf4, "Hello from io.WriteString!\n")

	// Cara 3: Write
	buf4.Write([]byte("Hello from Write!\n"))

	fmt.Printf("Buffer contents:\n%s", buf4.String())

	fmt.Println()
	fmt.Println("IMPLEMENTASI CUSTOM io.Reader")
	fmt.Println("---")

	counter := &CounterReader{count: 0, max: 5}

	buf5 := make([]byte, 20)
	for {
		n, err := counter.Read(buf5)
		if err == io.EOF {
			fmt.Println("  EOF!")
			break
		}
		fmt.Printf("  Read %d bytes: %s\n", n, string(buf5[:n]))
	}

	fmt.Println()
	fmt.Println("IMPLEMENTASI CUSTOM io.Writer")
	fmt.Println("---")

	logWriter := &LogWriter{}

	fmt.Fprint(logWriter, "Application started\n")
	fmt.Fprintf(logWriter, "User %s logged in\n", "admin")
	io.WriteString(logWriter, "Processing request...\n")

	fmt.Println()
	fmt.Println("IO.PIPETE - Hubungkan Reader dan Writer")
	fmt.Println("---")

	// io.Pipe menghubungkan io.Reader dan io.Writer
	// Data yang ditulis ke Writer bisa dibaca dari Reader
	pipeReader, pipeWriter := io.Pipe()

	// Writer goroutine
	go func() {
		defer pipeWriter.Close()
		fmt.Fprintln(pipeWriter, "Data dari goroutine writer")
		fmt.Fprintln(pipeWriter, "Baris kedua")
	}()

	// Baca dari pipe
	pipeData, _ := io.ReadAll(pipeReader)
	fmt.Printf("Data dari pipe:\n%s", pipeData)

	fmt.Println()
	fmt.Println("TIPS")
	fmt.Println("---")
	fmt.Println("- io.Reader    : Baca data (Read, ReadAt, ReadAll)")
	fmt.Println("- io.Writer    : Tulis data (Write, WriteString, WriteAt)")
	fmt.Println("- io.Copy      : Salin reader ke writer")
	fmt.Println("- io.MultiReader/Writer : Gabung banyak reader/writer")
	fmt.Println("- bytes.Buffer : In-memory buffer yang implements keduanya")
	fmt.Println("- strings.Builder : Efficient string concatenation")
}

// CounterReader - custom io.Reader yang mengembalikan angka berurutan
type CounterReader struct {
	count int
	max   int
}

func (r *CounterReader) Read(p []byte) (int, error) {
	if r.count >= r.max {
		return 0, io.EOF
	}

	data := fmt.Sprintf("[%d]", r.count)
	n := copy(p, data)
	r.count++

	return n, nil
}

// LogWriter - custom io.Writer yang menambahkan prefix
type LogWriter struct{}

func (w *LogWriter) Write(p []byte) (int, error) {
	fmt.Printf("[LOG] %s", string(p))
	return len(p), nil
}
