package main

import (
	"fmt"
	iago "go-ia/concepts"
	iagoPC "go-ia/patterns/creational"
	iagoPS "go-ia/patterns/structural"
)


func main() {
	fmt.Println("Hello World!")
	iago.ArrayExample()
	iago.NamelessFunctionExample()
	iago.SliceExample()
	iago.IterateExample()
	iago.MapExample()
	iago.StructExample()
	iago.StructWithMethodsExample()
	iago.ListExample()
	iago.QueueExample()
	iago.BoundedQueueExample()
	iago.BSTExample()
	//iago.FanInExample()
	//iago.FanOutExample()
	//iago.WorkerPoolExample()
	iagoPC.SingletonExample()
	iagoPC.BuilderExample()
	iagoPC.FactoryExample()
	iagoPC.AbstractFactoryExample()
	iagoPS.AdapterExample()
}
