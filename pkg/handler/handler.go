package handler

import (
	"net/http"
)

func RegHandlers() {
	http.HandleFunc("/", homeHandler)
}
