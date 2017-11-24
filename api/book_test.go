package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const json_test_str  = `{"title":"Cloud Native Go","author":"M.D.","isbn":"012334454345"}`

func TestBookToJSON(t *testing.T)  {
	book := Book{
		Title: "Cloud Native Go",
		Author: "M.D.",
		ISBN: "012334454345",
	}
	json := book.ToJSON()

	assert.Equal(t, json_test_str,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T)  {
	json := []byte(json_test_str)
	book := FromJSON(json)
	assert.Equal(t, Book{ Title: "Cloud Native Go", Author: "M.D.", ISBN: "012334454345", },
	book, "Book JSON unmarshalling wrong.")
}