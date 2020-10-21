// forms.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type ContactDetails struct {
	Email    string
	Password string
}

func main() {
	var port string

	if len(os.Args) > 1 {
		port = "localhost:" + os.Args[1]
	} else {
		port = "localhost:8080"
	}
	tmpl := template.Must(template.ParseFiles("forms.html"))
	s := &http.Server{
		Addr:    port,
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
		if strings.Index(details.Email, "@") != -1 {
			tmpl.Execute(w, struct{ Success bool }{Success: true})
		} else {
			tmpl.Execute(w, struct{ Success bool }{Success: false})
		}
	})

	// s.ListenAndServe()
	fmt.Println("Running Server on ", "https://"+s.Addr)
	log.Fatal(s.ListenAndServeTLS("deploy/cert.pem", "deploy/key.pem"))
}
