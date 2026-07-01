package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=130"`
	Role  string `json:"role"`
}

func (u User) Greet() string {
	return fmt.Sprintf("Halo, nama saya %s", u.Name)
}

func main() {
	fmt.Println("REFLECT TYPEOF - Cek Tipe")
	fmt.Println("---")

	x := 42
	s := "halo"
	f := 3.14
	b := true

	fmt.Printf("x: %v -> %s\n", x, reflect.TypeOf(x))
	fmt.Printf("s: %v -> %s\n", s, reflect.TypeOf(s))
	fmt.Printf("f: %v -> %s\n", f, reflect.TypeOf(f))
	fmt.Printf("b: %v -> %s\n", b, reflect.TypeOf(b))

	fmt.Println()
	fmt.Println("REFLECT VALUEOF - Ambil Value")
	fmt.Println("---")

	y := 100
	val := reflect.ValueOf(y)
	fmt.Printf("Value: %v\n", val)
	fmt.Printf("Type: %s\n", val.Type())
	fmt.Printf("Kind: %s\n", val.Kind())
	fmt.Printf("Int: %d\n", val.Int())

	fmt.Println()
	fmt.Println("REFLECT UNTUK STRUCT")
	fmt.Println("---")

	user := User{Name: "Budi", Email: "budi@gmail.com", Age: 25, Role: "admin"}
	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)

	fmt.Printf("Type: %s\n", userType)
	fmt.Printf("Kind: %s\n", userType.Kind())
	fmt.Printf("NumField: %d\n", userType.NumField())

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		value := userValue.Field(i)
		fmt.Printf("  %s (%s) = %v\n", field.Name, field.Type, value)
	}

	fmt.Println()
	fmt.Println("REFLECT FIELD TAGS (JSON, Validate)")
	fmt.Println("---")

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)

		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")

		fmt.Printf("  %s:\n", field.Name)
		fmt.Printf("    json:     %s\n", jsonTag)
		fmt.Printf("    validate: %s\n", validateTag)
	}

	fmt.Println()
	fmt.Println("REFLECT METHOD")
	fmt.Println("---")

	userType2 := reflect.TypeOf(&user)

	fmt.Printf("NumMethod: %d\n", userType2.NumMethod())

	for i := 0; i < userType2.NumMethod(); i++ {
		method := userType2.Method(i)
		fmt.Printf("  %s%s\n", method.Name, method.Type)
	}

	// Panggil method
	userValue2 := reflect.ValueOf(&user)
	method := userValue2.MethodByName("Greet")
	if method.IsValid() {
		result := method.Call(nil)
		fmt.Printf("  Greet(): %s\n", result[0])
	}

	fmt.Println()
	fmt.Println("REFLECT SET VALUE (Mutable)")
	fmt.Println("---")

	num := 10
	fmt.Printf("Sebelum: %d\n", num)

	numValue := reflect.ValueOf(&num).Elem()
	fmt.Printf("Type: %s, Settable: %v\n", numValue.Type(), numValue.CanSet())

	numValue.SetInt(99)
	fmt.Printf("Sesudah: %d\n", num)

	fmt.Println()
	fmt.Println("REFLECT SET STRUCT FIELD")
	fmt.Println("---")

	user2 := User{Name: "Andi", Email: "andi@test.com", Age: 20}
	fmt.Printf("Sebelum: %+v\n", user2)

	nameField := reflect.ValueOf(&user2).Elem().FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Andi Updated")
	}

	ageField := reflect.ValueOf(&user2).Elem().FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(30)
	}

	fmt.Printf("Sesudah: %+v\n", user2)

	fmt.Println()
	fmt.Println("REFLECT KIND CHECK")
	fmt.Println("---")

	values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}, User{}}

	for _, v := range values {
		t := reflect.TypeOf(v)
		fmt.Printf("  %-10v -> Kind: %-10s Type: %s\n", v, t.Kind(), t)
	}

	fmt.Println()
	fmt.Println("REFLECT UNTUK SLICE/MAP")
	fmt.Println("---")

	slice := []int{10, 20, 30, 40, 50}
	sliceVal := reflect.ValueOf(slice)

	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Length: %d\n", sliceVal.Len())
	fmt.Printf("Cap: %d\n", sliceVal.Cap())

	// Akses elemen per index
	for i := 0; i < sliceVal.Len(); i++ {
		fmt.Printf("  [%d] = %d\n", i, sliceVal.Index(i).Int())
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	mapVal := reflect.ValueOf(m)

	fmt.Printf("\nMap: %v\n", m)
	fmt.Printf("Len: %d\n", mapVal.Len())

	for _, key := range mapVal.MapKeys() {
		val := mapVal.MapIndex(key)
		fmt.Printf("  %s = %d\n", key, val)
	}

	fmt.Println()
	fmt.Println("REFLECT - STRUCT TO MAP")
	fmt.Println("---")

	user3 := User{Name: "Siti", Email: "siti@test.com", Age: 28, Role: "user"}
	result := structToMap(user3)
	fmt.Printf("Struct to Map: %v\n", result)

	fmt.Println()
	fmt.Println("REFLECT - VALIDATE STRUCT")
	fmt.Println("---")

	users := []User{
		{Name: "Budi", Email: "budi@gmail.com", Age: 25},
		{Name: "", Email: "invalid", Age: -5},
		{Name: "Siti", Email: "siti@test.com", Age: 100},
	}

	for _, u := range users {
		err := validateStruct(u)
		status := "VALID"
		if err != nil {
			status = err.Error()
		}
		fmt.Printf("  %+v -> %s\n", u, status)
	}

	fmt.Println()
	fmt.Println("TIPS REFLECT")
	fmt.Println("---")
	fmt.Println("- reflect.TypeOf()  -> dapat tipe")
	fmt.Println("- reflect.ValueOf() -> dapat value")
	fmt.Println("- .Kind()          -> tipe dasar (int, struct, slice)")
	fmt.Println("- .CanSet()        -> cek apakah bisa diubah")
	fmt.Println("- .FieldByName()   -> akses field struct")
	fmt.Println("- .MethodByName()  -> panggil method")
	fmt.Println("- Hati-hati: reflect lambat dan hard to debug")

	_ = strings.TrimSpace("")
}

// structToMap - konversi struct ke map menggunakan reflect
func structToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(s)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if typ.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		result[field.Name] = value.Interface()
	}

	return result
}

// validateStruct - validasi struct berdasarkan tag validate
func validateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if typ.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		validateTag := field.Tag.Get("validate")

		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			switch {
			case rule == "required":
				if value.String() == "" {
					return fmt.Errorf("%s is required", field.Name)
				}
			case rule == "email":
				if !strings.Contains(value.String(), "@") {
					return fmt.Errorf("%s must be valid email", field.Name)
				}
			case strings.HasPrefix(rule, "gte="):
				// Simplified: just check for negative
				if value.Int() < 0 {
					return fmt.Errorf("%s must be positive", field.Name)
				}
			}
		}
	}

	return nil
}
