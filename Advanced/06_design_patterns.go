package main

import (
	"fmt"
	"sync"
)

// Design Patterns demonstrates common design patterns in Go

func main() {
	singletonPattern()
	factoryPattern()
	builderPattern()
	observerPattern()
	strategyPattern()
	adapterPattern()
}

// singletonPattern demonstrates singleton pattern
func singletonPattern() {
	type Singleton struct {
		value int
	}
	
	var (
		instance *Singleton
		once     sync.Once
	)
	
	getInstance := func() *Singleton {
		once.Do(func() {
			instance = &Singleton{value: 42}
		})
		return instance
	}
	
	inst1 := getInstance()
	inst2 := getInstance()
	
	fmt.Printf("Singleton instances equal: %t\n", inst1 == inst2)
	fmt.Printf("Singleton value: %d\n", inst1.value)
}

// Product interface for factory pattern
type Product interface {
	Use() string
}

// ConcreteProductA implements Product
type ConcreteProductA struct{}

// Use returns product A description
func (p *ConcreteProductA) Use() string {
	return "Product A"
}

// ConcreteProductB implements Product
type ConcreteProductB struct{}

// Use returns product B description
func (p *ConcreteProductB) Use() string {
	return "Product B"
}

// factoryPattern demonstrates factory pattern
func factoryPattern() {
	// Factory
	createProduct := func(productType string) Product {
		switch productType {
		case "A":
			return &ConcreteProductA{}
		case "B":
			return &ConcreteProductB{}
		default:
			return nil
		}
	}
	
	productA := createProduct("A")
	productB := createProduct("B")
	
	fmt.Printf("Product A: %s\n", productA.Use())
	fmt.Printf("Product B: %s\n", productB.Use())
}

// Query represents a database query
type Query struct {
	table      string
	conditions []string
	orderBy    string
	limit      int
}

// QueryBuilder builds queries
type QueryBuilder struct {
	query Query
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// Table sets the table name
func (qb *QueryBuilder) Table(table string) *QueryBuilder {
	qb.query.table = table
	return qb
}

// Where adds a condition
func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.query.conditions = append(qb.query.conditions, condition)
	return qb
}

// OrderBy sets the order by field
func (qb *QueryBuilder) OrderBy(field string) *QueryBuilder {
	qb.query.orderBy = field
	return qb
}

// Limit sets the limit
func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.query.limit = n
	return qb
}

// Build returns the query
func (qb *QueryBuilder) Build() Query {
	return qb.query
}

// builderPattern demonstrates builder pattern
func builderPattern() {
	query := NewQueryBuilder().
		Table("users").
		Where("age > 18").
		Where("active = true").
		OrderBy("name").
		Limit(10).
		Build()
	
	fmt.Printf("Query: %+v\n", query)
}

// observerPattern demonstrates observer pattern
func observerPattern() {
	// Observer interface
	type Observer interface {
		Update(message string)
	}
	
	// Subject
	type Subject struct {
		observers []Observer
		mu        sync.Mutex
	}
	
	attach := func(s *Subject, observer Observer) {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.observers = append(s.observers, observer)
	}
	
	notify := func(s *Subject, message string) {
		s.mu.Lock()
		observers := make([]Observer, len(s.observers))
		copy(observers, s.observers)
		s.mu.Unlock()
		
		for _, observer := range observers {
			observer.Update(message)
		}
	}
	
	subject := &Subject{}
	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}
	
	attach(subject, observer1)
	attach(subject, observer2)
	
	notify(subject, "State changed")
}

// ConcreteObserver implements Observer
type ConcreteObserver struct {
	name string
}

// Update implements Observer interface
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("Observer %s received: %s\n", o.name, message)
}

// SortStrategy interface
type SortStrategy interface {
	Sort(data []int) []int
}

// BubbleSort implements SortStrategy
type BubbleSort struct{}

// Sort implements SortStrategy
func (s *BubbleSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort implements SortStrategy
type QuickSort struct{}

// Sort implements SortStrategy
func (s *QuickSort) Sort(data []int) []int {
	// Simplified quick sort
	result := make([]int, len(data))
	copy(result, data)
	// In production, implement full quicksort
	return result
}

// Sorter is the context for strategy pattern
type Sorter struct {
	strategy SortStrategy
}

// setStrategy sets the sorting strategy
func setStrategy(s *Sorter, strategy SortStrategy) {
	s.strategy = strategy
}

// sort sorts data using the current strategy
func sort(s *Sorter, data []int) []int {
	return s.strategy.Sort(data)
}

// strategyPattern demonstrates strategy pattern
func strategyPattern() {
	data := []int{5, 2, 8, 1, 9}
	
	sorter := &Sorter{}
	setStrategy(sorter, &BubbleSort{})
	sorted1 := sort(sorter, data)
	fmt.Printf("Bubble sorted: %v\n", sorted1)
	
	setStrategy(sorter, &QuickSort{})
	sorted2 := sort(sorter, data)
	fmt.Printf("Quick sorted: %v\n", sorted2)
}

// adapterPattern demonstrates adapter pattern
func adapterPattern() {
	// Target interface
	type Target interface {
		Request() string
	}
	
	adaptee := newAdaptee()
	adapter := newAdapter(adaptee)
	
	var target Target = adapter
	fmt.Printf("Adapted request: %s\n", target.Request())
}

// Adaptee represents the adaptee in adapter pattern
type Adaptee struct {
	specificRequest func() string
}

// newAdaptee creates a new adaptee
func newAdaptee() *Adaptee {
	return &Adaptee{
		specificRequest: func() string {
			return "Specific request"
		},
	}
}

// Adapter adapts Adaptee to Target interface
type Adapter struct {
	adaptee *Adaptee
}

// newAdapter creates a new adapter
func newAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request implements Target interface
func (a *Adapter) Request() string {
	return a.adaptee.specificRequest()
}
