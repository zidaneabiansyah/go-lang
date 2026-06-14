package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "..."
}

func (a Animal) Breathe() string {
	return "nafas"
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() string {
	return "Guk guk!"
}

type Cat struct {
	Animal
	FurColor string
}

func (c Cat) Speak() string {
	return "Meow~"
}

type Bird struct {
	Animal
	CanFly bool
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

type Server struct {
	Logger
	Host string
	Port int
}

func (s Server) Start() {
	s.Log(fmt.Sprintf("server jalan di %s:%d", s.Host, s.Port))
}

type Admin struct {
	*Logger
	Role string
}

func main() {
	fmt.Println("STRUCT EMBEDDING")
	fmt.Println()

	dog := Dog{Animal: Animal{Name: "Si Doggy"}, Breed: "Bulldog"}
	cat := Cat{Animal: Animal{Name: "Si Kitty"}, FurColor: "Orange"}

	fmt.Printf("%s: %s\n", dog.Name, dog.Speak())
	fmt.Printf("%s: %s\n", cat.Name, cat.Speak())

	fmt.Printf("%s bisa %s\n", dog.Name, dog.Breathe())
	fmt.Printf("%s bisa %s\n", cat.Name, cat.Breathe())

	fmt.Println()
	fmt.Println("METHOD PROMOTION")
	fmt.Println()

	bird := Bird{Animal: Animal{Name: "Si Tweet"}, CanFly: true}
	fmt.Printf("%s: %s\n", bird.Name, bird.Speak())
	fmt.Printf("%s bisa %s\n", bird.Name, bird.Breathe())

	fmt.Println()
	fmt.Println("INTERFACE EMBEDDING")
	fmt.Println()

	file := &File{}
	file.Write("Halo ini isi file")
	fmt.Println("Isi file:", file.Read())

	var rw ReadWriter = file
	rw.Write("data baru via interface")
	fmt.Println("Baca via interface:", rw.Read())

	fmt.Println()
	fmt.Println("EMBEDDING DENGAN METHOD")
	fmt.Println()

	server := Server{
		Logger: Logger{Prefix: "SERVER"},
		Host:   "localhost",
		Port:   8080,
	}
	server.Start()
	server.Log("ini log tambahan")

	fmt.Println()
	fmt.Println("EMBEDDING POINTER")
	fmt.Println()

	admin := Admin{
		Logger: &Logger{Prefix: "ADMIN"},
		Role:   "superadmin",
	}
	admin.Log("mengakses panel admin")
	fmt.Println("Role:", admin.Role)
}
