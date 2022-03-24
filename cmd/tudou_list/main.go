package main

import (
	"log"
	"net/http"

	"tudou_list/pkg/handler"
)

func main() {
	handler.Reg_handlers()

	log.Printf("About to listen on 8080. Go to https://0.0.0.0:8080/")
	// err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServeTLS(":8080", "../../tls/server.crt", "../../tls/server.key", nil)

	log.Fatal(err)
}
