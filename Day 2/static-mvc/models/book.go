package models

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var data = []Book{
	{Id: 1, Title: "The Alchemist", Author: "Paulo Coelho"},
	{Id: 2, Title: "The Kite Runner", Author: "Khaled Hosseini"},
}

func GetAllBooks() []Book {
	return data
}

func GetBookByID(id int) interface{} {
	for _, book := range data {
		if book.Id == id {
			return book
		}
	}
	return nil
}

func CreateBook(book Book) []Book {
	data = append(data, book)
	return data
}

func UpdateBook(id int, book Book) []Book {
	for i, u := range data {
		if u.Id == id {
			data[i] = book
		}
	}
	return data
}

func DeleteBook(id int) []Book {
	for i, book := range data {
		if book.Id == id {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return data
}
