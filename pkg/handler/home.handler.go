package handler

import (
	"io"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "11")
}
