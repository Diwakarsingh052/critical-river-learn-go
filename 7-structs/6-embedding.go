package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

// Composition

type BookAuthor struct {
	// with embedding we get automatic implementation of interface
	// if User implements the interface, then BookAuthor will also implement the interface in case of embedding
	User // embedding // anonymous field // field name would be same as the struct name
	//u     User // not embedding
	books []string
}

func (b *BookAuthor) UpdateEmail(email string) {
	b.email = email
}

func (b *BookAuthor) AddBook(book string) {
	b.books = append(b.books, book)
}

func main() {
	// bookAuthor can access the user fields and methods
	b := BookAuthor{
		User: User{
			name:  "bob",
			email: "bob@email.com",
		},
		books: []string{"book1", "book2"},
	}

	b.name = "bobby"
	b.UpdateEmail("")
	b.User.UpdateEmail("")
	b.AddBook("book3")
	fmt.Println(b)
}
