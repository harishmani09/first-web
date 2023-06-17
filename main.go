package main

import (
	"fmt"
	"net/http"
)

const webPort = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(webPort, nil)
}
