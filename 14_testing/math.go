package main

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

type User struct {
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) Validate() error {
	if u.Name == "" {
		return ErrEmptyName
	}
	if u.Age < 0 {
		return ErrNegativeAge
	}
	return nil
}
