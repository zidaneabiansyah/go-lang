package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string {
	return "Guk guk!"
}

type Cat struct{ Name string }

func (c Cat) Speak() string {
	return "Meow~"
}

type Robot struct{ Name string }

func (r Robot) Speak() string {
	return "BEEP BOOP"
}

func greet(s Speaker) {
	fmt.Printf("%T bilang: %s\n", s, s.Speak())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type EmptyInterface interface{}

func describe(i interface{}) {
	fmt.Printf("(%T, %v)\n", i, i)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v*2)
	case string:
		fmt.Println("string:", "\""+v+"\"")
	case bool:
		fmt.Println("bool:", !v)
	case Speaker:
		fmt.Println("speaker:", v.Speak())
	default:
		fmt.Println("gatau tipe ini")
	}
}

type Employee interface {
	GetSalary() float64
	GetRole() string
}

type FullTime struct {
	Name   string
	Salary float64
}

func (f FullTime) GetSalary() float64 {
	return f.Salary
}

func (f FullTime) GetRole() string {
	return "Full-time"
}

type Freelance struct {
	Name      string
	HourRate  float64
	HoursWork int
}

func (fr Freelance) GetSalary() float64 {
	return fr.HourRate * float64(fr.HoursWork)
}

func (fr Freelance) GetRole() string {
	return "Freelance"
}

func printPayroll(emps []Employee) {
	total := 0.0
	for _, e := range emps {
		fmt.Printf("%T - %s: Rp%.0f\n", e, e.GetRole(), e.GetSalary())
		total += e.GetSalary()
	}
	fmt.Printf("Total: Rp%.0f\n", total)
}

func main() {
	fmt.Println("INTERFACE BASIC")
	fmt.Println("---")

	dog := Dog{"Si Doggy"}
	cat := Cat{"Si Kitty"}
	robot := Robot{"R2-D2"}

	greet(dog)
	greet(cat)
	greet(robot)

	fmt.Println()
	fmt.Println("MULTIPLE INTERFACE")
	fmt.Println("---")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	shapes := []Shape{rect, circle}
	for _, s := range shapes {
		fmt.Printf("%T: area=%.2f, perimeter=%.2f\n", s, s.Area(), s.Perimeter())
	}

	fmt.Println()
	fmt.Println("EMPTY INTERFACE / ANY")
	fmt.Println("---")

	var x interface{} = 42
	describe(x)

	x = "halo"
	describe(x)

	x = true
	describe(x)

	x = dog
	describe(x)

	fmt.Println()
	fmt.Println("TYPE ASSERTION")
	fmt.Println("---")

	val := interface{}(42)
	num, ok := val.(int)
	fmt.Println(num, ok)

	str, ok := val.(string)
	fmt.Println(str, ok)

	fmt.Println()
	fmt.Println("TYPE SWITCH")
	fmt.Println("---")

	checkType(10)
	checkType("hello")
	checkType(false)
	checkType(dog)
	checkType(3.14)

	fmt.Println()
	fmt.Println("POLYMORPHISM")
	fmt.Println("---")

	budi := FullTime{"Budi", 5000000}
	rudi := Freelance{"Rudi", 100000, 40}

	printPayroll([]Employee{budi, rudi})
}
