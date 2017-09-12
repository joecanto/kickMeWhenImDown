// forms.go
package main

import (
	"io"
	"net/http"
)

func insult(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Get lost - loser!")
}

func main() {
	http.HandleFunc("/", insult)
	http.ListenAndServe(":8000", nil)
}