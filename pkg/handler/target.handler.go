package handler

import (
	"io"
	"net/http"
)

func Reg_handlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "11")
	})
}
