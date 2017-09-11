// forms.go
package main
import (
	"html/template"
	"net/http"
)

type Putdown struct {
	insult string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/forms.html"))

	http.HandleFunc("/insult", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := Putdown{
			insult: r.FormValue("insult"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}