package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexController)
	_ = http.ListenAndServe(":8888", nil)
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
