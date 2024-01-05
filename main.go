package main

import (
	"fmt"
	"net/http"
	"nihal-innsof/file-upload/templates/layout"
	"os"

	"github.com/a-h/templ"
)

func main() {
	index := layout.Index()
	http.Handle("/", templ.Handler(index))
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, fileInfo, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fileName := fileInfo.Filename
		if file != nil {
			fmt.Println("File recieved")
			fmt.Println("File name:", fileName)
		}
		f, err := os.Open("./uploads/fileName")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	})
	http.ListenAndServe(":8080", nil)
}
