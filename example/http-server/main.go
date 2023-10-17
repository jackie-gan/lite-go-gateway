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
	fmt.Println("server started on port 8099...")
	http.ListenAndServe(":8099", nil)
}
