package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

type Department struct {
	Name      string
	Employees []Employee
}

func (e Employee) getInfo() string {
	return fmt.Sprintf("%s (%d tahun) - Rp%.0f", e.Name, e.Age, e.Salary)
}

func (e *Employee) raiseSalary(percent float64) {
	e.Salary += e.Salary * percent / 100
}

func (e *Employee) birthday() {
	e.Age++
}

func printSep(title string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println("---")
}

func main() {
	printSep("STRUCT DASAR")

	budi := Employee{"Budi", 25, 5000000}
	siti := Employee{Name: "Siti", Age: 30, Salary: 7000000}
	var andi Employee
	andi.Name = "Andi"
	andi.Age = 28
	andi.Salary = 6000000

	fmt.Println(budi.getInfo())
	fmt.Println(siti.getInfo())
	fmt.Println(andi.getInfo())

	printSep("POINTER BASIC")

	x := 10
	p := &x
	fmt.Printf("x  = %d\n", x)
	fmt.Printf("p  = %p\n", p)
	fmt.Printf("*p = %d\n", *p)

	*p = 20
	fmt.Printf("setelah *p = 20, x = %d\n", x)

	printSep("POINTER KE STRUCT")

	ptr := &budi
	ptr.Salary = 5500000
	fmt.Println(budi.getInfo())

	printSep("METHOD POINTER RECEIVER")

	budi.raiseSalary(100)
	fmt.Println(budi.getInfo())

	budi.birthday()
	budi.birthday()
	fmt.Println(budi.getInfo())

	printSep("STRUCT FIELD ADDRESS")

	fmt.Printf("Alamat Name:   %p\n", &budi.Name)
	fmt.Printf("Alamat Age:    %p\n", &budi.Age)
	fmt.Printf("Alamat Salary: %p\n", &budi.Salary)

	printSep("NESTED STRUCT")

	itDept := Department{
		Name: "IT",
		Employees: []Employee{
			{Name: "Rudi", Age: 27, Salary: 8000000},
			{Name: "Dewi", Age: 26, Salary: 7500000},
		},
	}

	for _, emp := range itDept.Employees {
		fmt.Println("-", emp.getInfo())
	}

	printSep("PASS BY VALUE VS POINTER")

	emp := Employee{"Eko", 22, 4000000}

	changeName(emp)
	fmt.Println("setelah changeName:", emp.Name)

	changeNamePtr(&emp)
	fmt.Println("setelah changeNamePtr:", emp.Name)

	emp2 := emp
	emp2.Name = "Eko Kopi"
	fmt.Println("emp.Name:", emp.Name)
	fmt.Println("emp2.Name:", emp2.Name)

	emp3 := &emp
	emp3.Name = "Eko Banget"
	fmt.Println("emp.Name:", emp.Name)
	fmt.Println("emp3.Name:", emp3.Name)

	fmt.Println("\nnew -> new pointer")
	newEmp := new(Employee)
	newEmp.Name = "Baru"
	newEmp.Age = 20
	newEmp.Salary = 3000000
	fmt.Println(newEmp.getInfo())
}

func changeName(e Employee) {
	e.Name = "Tidak Berubah"
}

func changeNamePtr(e *Employee) {
	e.Name = "Berubah"
}
