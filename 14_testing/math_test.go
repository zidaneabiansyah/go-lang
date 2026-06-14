package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	expected := 12
	if result != expected {
		t.Errorf("Multiply(3, 4) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; expected 5", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err != ErrDivideByZero {
		t.Errorf("Divide(10, 0) expected ErrDivideByZero, got %v", err)
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{100, true},
	}

	for _, tt := range tests {
		result := IsEven(tt.input)
		if result != tt.expected {
			t.Errorf("IsEven(%d) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		err      error
	}{
		{"factorial 0", 0, 1, nil},
		{"factorial 1", 1, 1, nil},
		{"factorial 5", 5, 120, nil},
		{"factorial 10", 10, 3628800, nil},
		{"negative input", -1, 0, ErrNegativeInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Factorial(tt.input)
			if err != tt.err {
				t.Errorf("Factorial(%d) error = %v; expected %v", tt.input, err, tt.err)
			}
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUserIsAdult(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected bool
	}{
		{"adult 25", User{Name: "Budi", Age: 25}, true},
		{"minor 17", User{Name: "Andi", Age: 17}, false},
		{"exactly 18", User{Name: "Siti", Age: 18}, true},
		{"age 0", User{Name: "Bayi", Age: 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.user.IsAdult()
			if result != tt.expected {
				t.Errorf("%s IsAdult() = %v; expected %v", tt.user.Name, result, tt.expected)
			}
		})
	}
}

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name string
		user User
		err  error
	}{
		{"valid user", User{Name: "Budi", Age: 25}, nil},
		{"empty name", User{Name: "", Age: 25}, ErrEmptyName},
		{"negative age", User{Name: "Budi", Age: -5}, ErrNegativeAge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if err != tt.err {
				t.Errorf("Validate() error = %v; expected %v", err, tt.err)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}
