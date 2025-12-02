package main

import (
	"fmt"
	"time"
)

// Structs demonstrates struct types, fields, embedding, and methods

func main() {
	basicStructs()
	structMethods()
	structEmbedding()
	structTags()
	structComparison()
	anonymousStructs()
}

// Person represents a person with basic information
type Person struct {
	Name    string
	Age     int
	Email   string
	private int // unexported field
}

// basicStructs demonstrates basic struct usage
func basicStructs() {
	// Struct literal
	p1 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Printf("Person 1: %+v\n", p1)

	// Struct literal (positional)
	p2 := Person{"Bob", 25, "bob@example.com", 0}
	fmt.Printf("Person 2: %+v\n", p2)

	// Zero value struct
	var p3 Person
	fmt.Printf("Zero value Person: %+v\n", p3)

	// Field access
	p1.Name = "Alice Smith"
	fmt.Printf("Updated name: %s\n", p1.Name)

	// Pointer to struct
	p4 := &Person{
		Name:  "Carol",
		Age:   35,
		Email: "carol@example.com",
	}
	fmt.Printf("Pointer to Person: %+v\n", *p4)

	// Struct with pointer fields
	p5 := &Person{
		Name:  "David",
		Age:   28,
		Email: "david@example.com",
	}
	fmt.Printf("Person via pointer: %+v\n", p5)
}

// structMethods demonstrates methods on structs
func structMethods() {
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// Value receiver method
	fmt.Println(p.Greet())

	// Pointer receiver method (can modify)
	p.SetAge(31)
	fmt.Printf("After SetAge: %+v\n", p)

	// Method on pointer
	p2 := &Person{Name: "Bob", Age: 25}
	fmt.Println(p2.Greet())
}

// Greet returns a greeting string
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d years old", p.Name, p.Age)
}

// SetAge sets the person's age
func (p *Person) SetAge(age int) {
	if age < 0 || age > 150 {
		// Secure: validate input
		return
	}
	p.Age = age
}

// structEmbedding demonstrates struct embedding
func structEmbedding() {
	// Base struct
	type Address struct {
		Street  string
		City    string
		Country string
	}

	// Embedded struct
	type Employee struct {
		Person  // embedded
		Address // embedded
		ID      int
		Salary  float64
	}

	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  32,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		ID:     1001,
		Salary: 75000.0,
	}

	// Access embedded fields directly
	fmt.Printf("Employee: %s, Age: %d, City: %s\n", emp.Name, emp.Age, emp.City)

	// Explicit access
	fmt.Printf("Full address: %s, %s, %s\n", emp.Address.Street, emp.Address.City, emp.Address.Country)
}

// structTags demonstrates struct tags
func structTags() {
	// Struct with JSON tags
	type User struct {
		ID       int    `json:"id" db:"user_id"`
		Username string `json:"username" db:"username" validate:"required,min=3"`
		Email    string `json:"email" db:"email" validate:"required,email"`
		Password string `json:"-" db:"password_hash"` // excluded from JSON
	}

	u := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Password: "hashed_password",
	}

	fmt.Printf("User: %+v\n", u)
	fmt.Println("Note: Struct tags are used by encoding/json, database drivers, etc.")
}

// structComparison demonstrates struct comparison
func structComparison() {
	type Point struct {
		X, Y int
	}

	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 3}

	// Structs are comparable if all fields are comparable
	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// Structs with slices/maps are not comparable
	type Container struct {
		Data []int
	}

	c1 := Container{Data: []int{1, 2, 3}}
	c2 := Container{Data: []int{1, 2, 3}}
	// c1 == c2  // compile error: slices are not comparable
	fmt.Printf("c1: %+v, c2: %+v (not comparable)\n", c1, c2)
}

// anonymousStructs demonstrates anonymous structs
func anonymousStructs() {
	// Anonymous struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "Anonymous",
		Age:  99,
	}
	fmt.Printf("Anonymous struct: %+v\n", person)

	// Anonymous struct in map
	config := map[string]interface{}{
		"database": struct {
			Host string
			Port int
		}{
			Host: "localhost",
			Port: 5432,
		},
	}
	fmt.Printf("Config: %+v\n", config)
}

// advancedStructs demonstrates advanced struct patterns
func advancedStructs() {
	acc := Account{Balance: 100.0}

	// Secure: validate before operations
	if err := acc.Withdraw(50.0); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Balance after withdrawal: %.2f\n", acc.Balance)
	}

	if err := acc.Withdraw(100.0); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// Account methods
type Account struct {
	Balance float64
}

// Withdraw withdraws money from account
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// UserProfile demonstrates a more complex struct
type UserProfile struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Settings  map[string]interface{}
}

// NewUserProfile creates a new user profile with defaults
func NewUserProfile(username, email string) *UserProfile {
	now := time.Now()
	return &UserProfile{
		ID:        generateID(),
		Username:  username,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
		Settings:  make(map[string]interface{}),
	}
}

// generateID generates a simple ID (in production, use proper ID generation)
func generateID() int {
	return time.Now().Nanosecond()
}

// UpdateEmail updates the email with validation
func (up *UserProfile) UpdateEmail(newEmail string) error {
	// Secure: validate email format (simplified)
	if newEmail == "" {
		return fmt.Errorf("email cannot be empty")
	}
	up.Email = newEmail
	up.UpdatedAt = time.Now()
	return nil
}
