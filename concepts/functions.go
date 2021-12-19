package concepts

import "fmt"

func NamelessFunctionExample() {
	func(){
		z := 2 + 3
		fmt.Println(z)
	}()

}
