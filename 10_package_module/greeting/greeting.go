package greeting

import "fmt"

func Hello(name string) {
	fmt.Printf("Halo %s, selamat datang!\n", name)
}

func Goodbye(name string) {
	fmt.Printf("Sampai jumpa %s!\n", name)
}

func secretGreeting(name string) {
	fmt.Printf("ssst, rahasia: %s\n", name)
}
