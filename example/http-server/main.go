package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is Server Handler")
}

func main() {
	http.HandleFunc("/test", handler)
	if err := http.ListenAndServe(":8099", nil); err == nil {
		fmt.Println("server started on port 8099...")
	}
}
