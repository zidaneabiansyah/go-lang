package mathutils

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

var internalCounter = 0

type Calculator struct {
	History []int
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	c.History = append(c.History, total)
	internalCounter++
	return total
}

func GetInternalCounter() int {
	return internalCounter
}
