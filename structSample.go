// You can edit this code!
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	

	var Book1 Books
	//var Book2 Books

	Book1.title = "Go"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go programming"
	Book1.book_id = 98989
	printBook(Book1)

}

func printBook(book Books) {
	fmt.Println(book.title)
	fmt.Println(book.author)
}
