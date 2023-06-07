package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hahahah"))
	})
	http.ListenAndServe(":8888", nil)
}
