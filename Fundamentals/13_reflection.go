package main

import (
	"fmt"
	"reflect"
)

// Reflection demonstrates reflection capabilities (use sparingly)

func main() {
	basicReflection()
	typeInspection()
	valueInspection()
	valueModification()
	structReflection()
	reflectionBestPractices()
}

// basicReflection demonstrates basic reflection
func basicReflection() {
	var x int = 42

	// Get type
	t := reflect.TypeOf(x)
	fmt.Printf("Type: %v, Kind: %v\n", t, t.Kind())

	// Get value
	v := reflect.ValueOf(x)
	fmt.Printf("Value: %v, Type: %v\n", v, v.Type())

	// Type assertions vs reflection
	var i interface{} = 42

	// Type assertion (compile-time)
	if val, ok := i.(int); ok {
		fmt.Printf("Type assertion: %d\n", val)
	}

	// Reflection (runtime)
	val := reflect.ValueOf(i)
	fmt.Printf("Reflection: %v (kind: %v)\n", val, val.Kind())
}

// typeInspection demonstrates type inspection
func typeInspection() {
	// Inspect struct type
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	t := reflect.TypeOf(p)

	fmt.Printf("Struct name: %s\n", t.Name())
	fmt.Printf("Number of fields: %d\n", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("  Field %d: %s (type: %v, tag: %s)\n",
			i, field.Name, field.Type, field.Tag)
	}

	// Inspect function type
	funcType := reflect.TypeOf(add)
	fmt.Printf("Function: %v\n", funcType)
	fmt.Printf("  Input parameters: %d\n", funcType.NumIn())
	fmt.Printf("  Output parameters: %d\n", funcType.NumOut())

	for i := 0; i < funcType.NumIn(); i++ {
		fmt.Printf("    Param %d: %v\n", i, funcType.In(i))
	}
}

// add is a simple function for reflection
func add(a, b int) int {
	return a + b
}

// valueInspection demonstrates value inspection
func valueInspection() {
	// Inspect struct values
	type Point struct {
		X, Y int
	}

	point := Point{X: 10, Y: 20}
	v := reflect.ValueOf(point)

	fmt.Printf("Point: %v\n", point)
	fmt.Printf("Number of fields: %d\n", v.NumField())

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("  Field %d: %v (type: %v)\n", i, field, field.Type())
	}

	// Check if value is zero
	var zeroPoint Point
	zeroVal := reflect.ValueOf(zeroPoint)
	fmt.Printf("Is zero: %t\n", isZero(zeroVal))

	// Inspect slice
	slice := []int{1, 2, 3}
	sliceVal := reflect.ValueOf(slice)
	fmt.Printf("Slice length: %d\n", sliceVal.Len())
	for i := 0; i < sliceVal.Len(); i++ {
		fmt.Printf("  [%d] = %v\n", i, sliceVal.Index(i))
	}
}

// isZero checks if a value is zero
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map, reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}

// valueModification demonstrates value modification
func valueModification() {
	// Modify through pointer
	x := 42
	v := reflect.ValueOf(&x).Elem()
	fmt.Printf("Original: %d\n", x)

	v.SetInt(100)
	fmt.Printf("Modified: %d\n", x)

	// Modify struct fields
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	vp := reflect.ValueOf(&p).Elem()

	nameField := vp.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Bob")
	}

	ageField := vp.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(25)
	}

	fmt.Printf("Modified person: %+v\n", p)

	// Create new values
	newInt := reflect.New(reflect.TypeOf(0)).Elem()
	newInt.SetInt(42)
	fmt.Printf("New int: %d\n", newInt.Int())
}

// structReflection demonstrates struct-specific reflection
func structReflection() {
	type User struct {
		ID       int    `json:"id" db:"user_id"`
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"email"`
		password string `json:"-"` // private field
	}

	user := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		password: "secret",
	}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	fmt.Println("Struct reflection:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		// Check if field is exported
		if field.PkgPath == "" {
			fmt.Printf("  %s: %v\n", field.Name, fieldVal)

			// Extract tags
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				fmt.Printf("    JSON tag: %s\n", jsonTag)
			}
		} else {
			fmt.Printf("  %s: <unexported>\n", field.Name)
		}
	}
}

// reflectionBestPractices demonstrates best practices
func reflectionBestPractices() {
	// 1. Use reflection sparingly (performance cost)
	// 2. Prefer type assertions when possible
	// 3. Check IsValid() and CanSet() before operations
	// 4. Handle errors properly
	// 5. Use for generic operations (JSON, database, etc.)

	fmt.Println("Reflection best practices:")
	fmt.Println("  - Use sparingly (performance cost)")
	fmt.Println("  - Prefer type assertions")
	fmt.Println("  - Always check IsValid() and CanSet()")
	fmt.Println("  - Handle errors properly")

	// Example: Safe field access
	safeFieldAccess()
}

// safeFieldAccess demonstrates safe field access
func safeFieldAccess() {
	type Config struct {
		Host string
		Port int
	}

	config := Config{Host: "localhost", Port: 8080}
	v := reflect.ValueOf(&config).Elem()

	// Secure: Check before accessing
	hostField := v.FieldByName("Host")
	if hostField.IsValid() && hostField.CanSet() {
		if hostField.Kind() == reflect.String {
			hostField.SetString("0.0.0.0")
		}
	}

	fmt.Printf("Config: %+v\n", config)
}

// callFunction demonstrates calling functions via reflection
func callFunction() {
	// Get function value
	fn := reflect.ValueOf(add)

	// Prepare arguments
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}

	// Call function
	results := fn.Call(args)

	if len(results) > 0 {
		fmt.Printf("Result: %v\n", results[0].Int())
	}
}

// createInstance demonstrates creating instances via reflection
func createInstance() {
	// Create slice
	sliceType := reflect.SliceOf(reflect.TypeOf(0))
	slice := reflect.MakeSlice(sliceType, 3, 3)

	for i := 0; i < 3; i++ {
		slice.Index(i).SetInt(int64(i + 1))
	}

	fmt.Printf("Created slice: %v\n", slice.Interface())

	// Create map
	mapType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	m := reflect.MakeMap(mapType)

	key := reflect.ValueOf("one")
	value := reflect.ValueOf(1)
	m.SetMapIndex(key, value)

	fmt.Printf("Created map: %v\n", m.Interface())
}
