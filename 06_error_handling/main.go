package main

import (
	"errors"
	"fmt"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("gak bisa bagi dengan 0")
	}
	return a / b, nil
}

type DivideError struct {
	A, B int
	Msg  string
}

func (e DivideError) Error() string {
	return fmt.Sprintf("divide error: %d / %d -> %s", e.A, e.B, e.Msg)
}

func divideCustom(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideError{A: a, B: b, Msg: "penyebut 0, gak sah"}
	}
	return a / b, nil
}

type ValidationError struct {
	Field string
	Value interface{}
	Rule  string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal: %s = %v, syarat: %s", e.Field, e.Value, e.Rule)
}

func validateUser(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "name", Value: name, Rule: "wajib diisi"}
	}
	if age < 17 {
		return ValidationError{Field: "age", Value: age, Rule: "minimal 17 tahun"}
	}
	return nil
}

type MyError struct {
	Code    int
	Message string
	Err     error
}

func (e MyError) Error() string {
	return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
}

func (e MyError) Unwrap() error {
	return e.Err
}

func readConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return MyError{Code: 404, Message: "file gak ketemu", Err: err}
	}
	return nil
}

func riskyOperation(flag bool) (string, error) {
	if flag {
		return "berhasil", nil
	}
	return "", errors.New("sesuatu error bro")
}

func main() {
	fmt.Println("ERROR HANDLING BASIC")
	fmt.Println("---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	fmt.Println("CUSTOM ERROR")
	fmt.Println("---")

	_, err = divideCustom(5, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("VALIDATION ERROR")
	fmt.Println("---")

	err = validateUser("", 20)
	if err != nil {
		fmt.Println(err)
	}

	err = validateUser("Budi", 15)
	if err != nil {
		fmt.Println(err)
	}

	err = validateUser("Budi", 25)
	if err != nil {
		fmt.Println("gak boleh sampai sini")
	} else {
		fmt.Println("validasi sukses")
	}

	fmt.Println()
	fmt.Println("ERROR WRAPPING")
	fmt.Println("---")

	err = readConfig("/tmp/gak-ada-file.conf")
	if err != nil {
		fmt.Println(err)

		var myErr MyError
		if errors.As(err, &myErr) {
			fmt.Printf("Kode error: %d\n", myErr.Code)
		}
	}

	fmt.Println()
	fmt.Println("ERROR SENTINEL & IS")
	fmt.Println("---")

	_, err = riskyOperation(false)
	if errors.Is(err, errors.New("sesuatu error bro")) {
		fmt.Println("ketemu error yang sama (by message)")
	}

	ErrNotFound := errors.New("not found")
	err = fmt.Errorf("wrapping: %w", ErrNotFound)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("ketemu ErrNotFound di wrapped error")
	}

	fmt.Println()
	fmt.Println("PANIC & RECOVER")
	fmt.Println("---")

	handlePanic()

	fmt.Println()
	fmt.Println("DEFER")
	fmt.Println("---")

	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("error buka file:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 50)
	n, _ := file.Read(buf)
	fmt.Printf("Baca file: %s\n", string(buf[:n]))
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("terjadi panic, tapi diamankan:", r)
		}
	}()

	fmt.Println("mau panic...")
	panic("ada yang salah parah!")
}
