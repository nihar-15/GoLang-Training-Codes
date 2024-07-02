// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Rose": "Red"}
	fmt.Println(flowerColor["Sunflower"])
	fmt.Println(flowerColor["Roose"]) // no error

	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}
	fmt.Println("Initial map:", capital)

	capital["US"] = "Washigton DC"

	capital["India"] = "Delhi"
	fmt.Println("Updated map:", capital)

	delete(capital, "Nepal")
	//student := map[int]string{1:"Joe"}

	fmt.Println("Updated map:", capital)

	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	for num, square := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", num, square)
	}

	for num, square := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", num, square)
	}

	for num, square := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", num, square)
	}

	for num, square := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", num, square)
	}

	student := make(map[int]string)

	student[1] = "Lily"
	student[2] = "Hari"

	fmt.Println(student)
	fmt.Printf("%T\n", student)

	for num := range student {
		fmt.Println(num)
	}

	//m := map[string]Vertex{
	//	"Bell Labs": {40.68, -74.3999},
	//	"Google":    {37.42, -122.08045},
	//}

	//fmt.Println(m)

	for _, capital := range capital {
		fmt.Println(capital)
	}

	m := make(map[string]int)

	m["answer"] = 42

	fmt.Println("Value", m["answer"])

	m["answer"] = 48
	fmt.Println("Value", m["answer"])

	delete(m, "answer")
	fmt.Println("Value", m["answer"])

	val, isPresent := m["answer"]
	fmt.Println("Value", val, " Is Present", isPresent)

}
