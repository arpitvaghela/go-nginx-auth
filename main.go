// forms.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
)

type ContactDetails struct {
	Email    string
	Password string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Some validation
		if strings.Index(details.Email, "/") == -1 {
			tmpl.Execute(w, struct{ Success bool }{Success: true})
		} else {
			tmpl.Execute(w, struct{ Success bool }{Success: false})
		}
	})

	// s.ListenAndServe()
	http2.ConfigureServer(s, &http2.Server{})
	log.Fatal(s.ListenAndServeTLS("deploy/cert.pem", "deploy/key.pem"))
}
