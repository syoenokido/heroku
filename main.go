package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templataHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func (t *templataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 =
			template.Must(template.ParseFiles(filepath.Join("template",
				t.filename)))
	})
	t.temp1.Execute(w, nil)
}

func main() {
	http.Handle("/", &templataHandler{filename: "chat.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
