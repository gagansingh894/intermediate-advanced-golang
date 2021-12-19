package concepts

import "fmt"

func ArrayExample() {
	var myArray [3]int32
	fmt.Println("My initial array", myArray)
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println("My array after setting values", myArray)

	// alternate way od defining array
	mySecondArray := [3]int32{2, 4, 5}
	fmt.Println("My second array is", mySecondArray)
}

func SliceExample() {
	var mySlice []int32
	fmt.Printf("%T", mySlice)
	fmt.Println()
	fmt.Println("Capacity is: ", cap(mySlice))
	fmt.Println("Length is: ", len(mySlice))
	mySlice = append(mySlice, 1)
	fmt.Println("Capacity is: ", cap(mySlice))
	fmt.Println("Length is: ", len(mySlice))
	mySlice = append(mySlice, 2)
	fmt.Println("Capacity is: ", cap(mySlice))
	fmt.Println("Length is: ", len(mySlice))
	mySlice = append(mySlice, 3)
	fmt.Println("Capacity is: ", cap(mySlice))
	fmt.Println("Length is: ", len(mySlice))
	mySliceTwo := []string{"Gagan", "Sahib", "Savy"}
	fmt.Println(mySliceTwo)
}

func MapExample() {
	names := []string{"Gagan", "Sahib", "Savy"}
	marks := []int32{20, 30, 40}
	studentMarks := make(map[string]int32)
	fmt.Println("My map before assignment", studentMarks)
	for i:=0; i<=len(names)-1; i++ {
		studentMarks[names[i]] = marks[i]
	}
	fmt.Println("My map after assignment", studentMarks)
	value, err := studentMarks["Gagan"]
	fmt.Println(err, value)
	value, err = studentMarks["Bitcoin"]
	fmt.Println(err, value)

	for key, value := range studentMarks {
		fmt.Println(key, value)
	}
}

func IterateExample() {
	myArray := []int32{1,2,3}
	for i, v := range myArray {
		fmt.Println("Index is: ", i, " Value is: ", v)
	}
}