package api

import (
	"encoding/json"
	"net/http"
	"github.com/mdomlad85/GoMicroservices/util"
	"io/ioutil"
)

// Book type with Name, Author and ISBN
type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	// define the book
}

func (b Book) ToJSON() []byte  {
	ToJSON, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}
	return  ToJSON
}

func FromJSON(data []byte) Book  {
	book := Book{}
	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}
	return book
}

var books = map[string]Book{
	"030324003242": Book{ Title: "The Hitchhiker's guide to the Galaxy", Author: "Douglas Adams", ISBN: "030324003242", },
	"000000000000": Book{ Title: "Cloud Native Go", Author: "M.D.", ISBN: "000000000000", },
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		booksArr := AllBooks()
		err := util.WriteJSON(w, booksArr)

		if err != nil {
			panic(err)
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)

		if created {
			w.Header().Add("Location", "/api/books/" + isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)

		if found {
			util.WriteJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)

		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}

}

func AllBooks() []Book {
	booksArr := make([]Book, len(books))
	i := 0
	for _, book := range books {
		booksArr[i] = book
		i++
	}
	return booksArr
}

func GetBook(isbn string) (Book, bool) {
	book, ok :=  books[isbn]
	return book, ok
}

func CreateBook(book Book) (string, bool) {
	_, ok := books[book.ISBN]
	if !ok && len(book.ISBN) > 0 {
		books[book.ISBN] = book

		return book.ISBN, true
	}
	return "", false
}

func UpdateBook(isbn string, book Book) bool {
	_, ok := books[isbn]
	if ok {
		books[isbn] = book

		return true
	}
	return false
}

func DeleteBook(isbn string) {
	delete(books, isbn)
}