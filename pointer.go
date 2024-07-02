// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var intData = 20
	var intPointer *int

	intPointer = &intData

	fmt.Println("What intData stores:", intData)
	fmt.Println("address of intDta :", &intData)
	fmt.Println("What intPointer  stores:", intPointer)

	*intPointer = 39
	fmt.Println("What intData stores now:", intData)

	var name = "John"
	var ptr *string

	ptr = &name

	fmt.Println("Value of pointer ", ptr)
	fmt.Println("Address of  variable ", &name)
	fmt.Println("Value of variable ", *ptr)

}
