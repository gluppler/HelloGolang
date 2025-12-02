package main

import (
	"fmt"
	"reflect"
)

// Advanced Reflection demonstrates advanced reflection techniques

func main() {
	dynamicFunctionCalls()
	structTagParsing()
	typeValidation()
	dynamicStructCreation()
	reflectionPerformance()
}

// dynamicFunctionCalls demonstrates calling functions dynamically
func dynamicFunctionCalls() {
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
		fmt.Printf("Dynamic call result: %v\n", results[0].Int())
	}

	// Call with variadic arguments
	sumFn := reflect.ValueOf(sum)
	variadicArgs := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
		reflect.ValueOf(3),
	}
	variadicResults := sumFn.Call(variadicArgs)
	fmt.Printf("Variadic call result: %v\n", variadicResults[0].Int())
}

// add performs addition
func add(a, b int) int {
	return a + b
}

// sum sums variadic arguments
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// structTagParsing demonstrates parsing struct tags
func structTagParsing() {
	type User struct {
		ID       int    `json:"id" db:"user_id" validate:"required"`
		Username string `json:"username" db:"username" validate:"required,min=3"`
		Email    string `json:"email" db:"email" validate:"required,email"`
		Password string `json:"-" db:"password_hash"`
	}

	user := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Password: "secret",
	}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	fmt.Println("Struct tags:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		validateTag := field.Tag.Get("validate")

		fmt.Printf("  %s:\n", field.Name)
		fmt.Printf("    Value: %v\n", fieldVal)
		if jsonTag != "" {
			fmt.Printf("    JSON tag: %s\n", jsonTag)
		}
		if dbTag != "" {
			fmt.Printf("    DB tag: %s\n", dbTag)
		}
		if validateTag != "" {
			fmt.Printf("    Validate tag: %s\n", validateTag)
		}
	}
}

// typeValidation demonstrates type validation
func typeValidation() {
	// Validate struct fields
	type Config struct {
		Host string `validate:"required"`
		Port int    `validate:"required,min=1,max=65535"`
	}

	config := Config{
		Host: "localhost",
		Port: 8080,
	}

	err := validateStruct(config)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Validation passed")
	}
}

// validateStruct validates a struct using reflection
func validateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// Handle pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	// Must be struct
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		// Check required
		if validateTag == "required" {
			if isZero(fieldVal) {
				return fmt.Errorf("field %s is required", field.Name)
			}
		}

		// Additional validation logic here
	}

	return nil
}

// isZero checks if value is zero
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

// dynamicStructCreation demonstrates creating structs dynamically
func dynamicStructCreation() {
	// Create struct type dynamically
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
		},
	}

	structType := reflect.StructOf(fields)

	// Create instance
	instance := reflect.New(structType).Elem()
	instance.Field(0).SetString("Alice")
	instance.Field(1).SetInt(30)

	fmt.Printf("Dynamic struct: %+v\n", instance.Interface())

	// Create slice dynamically
	sliceType := reflect.SliceOf(structType)
	slice := reflect.MakeSlice(sliceType, 0, 10)

	// Add elements
	for i := 0; i < 3; i++ {
		elem := reflect.New(structType).Elem()
		elem.Field(0).SetString(fmt.Sprintf("Person%d", i))
		elem.Field(1).SetInt(int64(20 + i))
		slice = reflect.Append(slice, elem)
	}

	fmt.Printf("Dynamic slice: %+v\n", slice.Interface())
}

// reflectionPerformance demonstrates reflection performance considerations
func reflectionPerformance() {
	// Reflection has performance overhead
	// Use sparingly in hot paths

	type Data struct {
		Value int
	}

	data := Data{Value: 42}

	// Direct access (fast)
	directValue := data.Value
	_ = directValue

	// Reflection access (slower)
	v := reflect.ValueOf(data)
	reflectedValue := v.FieldByName("Value").Int()
	_ = reflectedValue

	fmt.Println("Reflection performance:")
	fmt.Println("  - Direct access: O(1), very fast")
	fmt.Println("  - Reflection: Slower, use for dynamic operations only")
	fmt.Println("  - Cache reflect.Type and reflect.Value when possible")
}
