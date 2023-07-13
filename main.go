package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", middleware2Handler(middleware1Handler(http.HandlerFunc(IndexController))))
}

func IndexController(writer http.ResponseWriter, request *http.Request) {
	m := make(map[string]interface{})
	m["code"] = 0
	m["message"] = "success"
	m["data"] = nil
	marshal, err := json.Marshal(m)
	if err != nil {
		return
	}
	_, _ = writer.Write(marshal)
}

func middleware1Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1")
		next.ServeHTTP(w, r)
	})
}

func middleware2Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware2")
		next.ServeHTTP(w, r)
	})
}
