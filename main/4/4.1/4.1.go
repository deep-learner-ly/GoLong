package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		return
	}
}
func main() {
	//http.HandleFunc("/", handler)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	return
	//}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, r.URL.String())
	//})
	//http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.URL.Query() {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	})
	http.ListenAndServe(":8080", nil)
}
