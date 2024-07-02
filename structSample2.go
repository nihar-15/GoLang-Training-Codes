// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type Student struct {
	name string
}

type Rectangle func(int, int) int

type rectPara struct {
	length  int
	breadth int
	color   string

	rect Rectangle
}

func main() {
	fmt.Println("Hello, 世界")

	var Book1 Books
	//var Book2 Books

	Book1.title = "Go"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go programming"
	Book1.book_id = 98989
	printBook(Book1)

	result := rectPara{
		length:  10,
		breadth: 3,
		color:   "blue",
		rect: func(l, b int) int {
			return l * b
			//return 2*l +2*b
		},
	}

	result1 := rectPara{
		length:  10,
		breadth: 3,
		color:   "blue",
		rect: func(l, b int) int {

			return 2*l + 2*b
		},
	}

	fmt.Println("Color of Rectangle", result.color)
	fmt.Println("Area", result.rect(8, 9))
	fmt.Println("Color of Rectangle", result1.color)
	fmt.Println("Area", result1.rect(8, 9))

	numbers := []int{1, 3, 4, 6}
	fmt.Println("Numbers", numbers)

	var nums = make([]int, 3, 5)
	fmt.Printf("len =%d, cap=%d,", len(nums), cap(nums))

	arr := [8]int{1, 87, 3, 5, 6, 7, 9}
	sliceNums := arr[4:7]

	fmt.Printf("Type of sliceNums: %T\n", sliceNums)
	fmt.Printf("Type of arr: %T\n", arr)

	//slice functions
	s := []int{7, 1}
	s = append(s, 17, 1)
	fmt.Println(s)

	s = append(s, sliceNums...)
	fmt.Println(s)

	//sub slicing

	fmt.Println(s[1:3])

	// missing lower bound implies 0
	fmt.Println("Missing lower bound", s[:3])

	// missing upper bound implies len
	fmt.Println("Missing upper  bound", s[3:])

	even := []int{2, 10, 4, 7}
	odd := []int{12, 14, 17}

	copy(odd, even) // Copy elements from even to odd

	fmt.Println(odd)

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	intSlice := arr1[2:5]

	for index, ele := range intSlice {
		fmt.Println(index, ele)
	}

	for _, ele := range intSlice {
		fmt.Println(ele)
	}
	intSlice[1] = 190
	println("After change")
	for index := range intSlice {
		fmt.Println(intSlice[index])
	}
	//keys := []int{3, 2, 8, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)

	slice := []string{"honesty", "is", "the", "best", "policy"}
	sort.Strings(slice)
	fmt.Println(slice)

	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	slice2 := []int{30, 10, 20, 40, 50, 70, 60, 80}

	status := sort.IntsAreSorted(slice1)

	if status == true {
		println("Slice is sorted")
	} else {
		println("Slice is not  sorted")
	}

	status = sort.IntsAreSorted(slice2)
	if status == true {
		println("Slice2 is sorted")
	} else {
		println("Slice2 is not  sorted")
	}

	slice3 := []string{"honesty", "is", "the", "best", "policy"}

	status = sort.StringsAreSorted(slice3)
	if status == true {
		println("Slice 3 is sorted")
	} else {
		println("Slice 3 is not  sorted") // also we have Float64sAreSorted
	}

}

func printBook(book Books) {
	fmt.Println(book.title)
	fmt.Println(book.author)
}
