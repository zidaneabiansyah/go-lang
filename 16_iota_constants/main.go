package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	days := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	if d < 0 || int(d) >= len(days) {
		return "Unknown"
	}
	return days[d]
}

type OrderStatus int

const (
	StatusPending OrderStatus = iota
	StatusProcessing
	StatusShipped
	StatusDelivered
	StatusCancelled
)

func (s OrderStatus) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusProcessing:
		return "Diproses"
	case StatusShipped:
		return "Dikirim"
	case StatusDelivered:
		return "Terkirim"
	case StatusCancelled:
		return "Dibatalkan"
	default:
		return "Unknown"
	}
}

func (s OrderStatus) IsActive() bool {
	return s == StatusPending || s == StatusProcessing || s == StatusShipped
}

type Permission int

const (
	Read Permission = 1 << iota
	Write
	Execute
	Delete
	Admin = Read | Write | Execute | Delete
)

func (p Permission) Has(perm Permission) bool {
	return p&perm != 0
}

func (p Permission) String() string {
	var result string
	if p.Has(Read) {
		result += "R"
	} else {
		result += "-"
	}
	if p.Has(Write) {
		result += "W"
	} else {
		result += "-"
	}
	if p.Has(Execute) {
		result += "X"
	} else {
		result += "-"
	}
	if p.Has(Delete) {
		result += "D"
	} else {
		result += "-"
	}
	return result
}

type Size int

const (
	_ Size = iota
	Small
	Medium
	Large
	ExtraLarge
)

func (s Size) String() string {
	switch s {
	case Small:
		return "Kecil"
	case Medium:
		return "Sedang"
	case Large:
		return "Besar"
	case ExtraLarge:
		return "Extra Besar"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("IOTA BASIC — ENUM HARI")
	for d := Sunday; d <= Saturday; d++ {
		fmt.Printf("  Day(%d): %s\n", d, d)
	}
	fmt.Printf("  Sunday type: %T, value: %d\n", Sunday, Sunday)

	fmt.Println()
	fmt.Println("IOTA UNTUK ORDER STATUS")
	statuses := []OrderStatus{StatusPending, StatusProcessing, StatusShipped, StatusDelivered, StatusCancelled}
	for _, s := range statuses {
		fmt.Printf("  %d: %s (active: %v)\n", s, s, s.IsActive())
	}

	orderStatus := StatusShipped
	fmt.Printf("\n  order status sekarang: %s\n", orderStatus)
	fmt.Printf("  masih aktif? %v\n", orderStatus.IsActive())

	fmt.Println()
	fmt.Println("IOTA BITMASK — PERMISSIONS")
	var userPerm Permission = Read | Write
	fmt.Printf("  User permission: %s\n", userPerm)
	fmt.Printf("  Can Read? %v\n", userPerm.Has(Read))
	fmt.Printf("  Can Execute? %v\n", userPerm.Has(Execute))

	adminPerm := Admin
	fmt.Printf("  Admin permission: %s\n", adminPerm)
	fmt.Printf("  Admin can Delete? %v\n", adminPerm.Has(Delete))

	var guestPerm Permission = Read
	fmt.Printf("  Guest permission: %s\n", guestPerm)

	fmt.Println()
	fmt.Println("IOTA DENGAN SKIP (underscore)")
	for s := Small; s <= ExtraLarge; s++ {
		fmt.Printf("  Size(%d): %s\n", s, s)
	}

	fmt.Println()
	fmt.Println("TIP: iota reset ke 0 di setiap block const baru")
}
