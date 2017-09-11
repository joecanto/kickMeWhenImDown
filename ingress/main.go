// forms.go
package main
import (
	"fmt"
	"net/http"
	"log"
)

type Putdown struct {
	insult string
}

func main() {
	http.HandleFunc("/insult", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, everyone")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	}