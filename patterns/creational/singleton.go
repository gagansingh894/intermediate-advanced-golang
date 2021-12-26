package creational

import "fmt"

// Singleton
type singleton struct {
	count int
}

type Singleton interface {
	AddOne() int
	GetCount() int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count ++
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}

func SingletonExample() {
	counter := GetInstance()
	fmt.Printf("%T\n", counter)
	fmt.Printf("Count is %d\n", counter.GetCount())
	counter.AddOne()
	fmt.Printf("Count is %d\n", counter.GetCount())
	counter.AddOne()
	fmt.Printf("Count is %d\n", counter.GetCount())
	counter2 := GetInstance()
	fmt.Printf("Count is %d\n", counter2.GetCount())

}
