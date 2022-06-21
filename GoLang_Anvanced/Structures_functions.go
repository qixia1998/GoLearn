package main

import "fmt"

type Bookx struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Bookx  /* Declare Book1 of type Book */
	var Book2 Bookx	 /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Jason Stathom"
	Book1.subject = "Go Programming Video Course"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zanis Ali Rhan"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	printBook(Book1)

	/* print Book2 info */
	printBook(Book2)

}

func printBook ( book Bookx ) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}