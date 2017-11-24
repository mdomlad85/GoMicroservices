package util

import "net/http"

type Render interface {
	Render(http.ResponseWriter) error
}

var (
	_ Render     = JSON{}
	_ Render     = IndentedJSON{}
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
