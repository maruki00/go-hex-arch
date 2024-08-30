package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})

	fmt.Println("SErver running on http://127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)

}
