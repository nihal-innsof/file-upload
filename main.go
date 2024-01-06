package main

import (
	"fmt"
	"io"
	"net/http"
	"nihal-innsof/file-upload/templates/components"
	"nihal-innsof/file-upload/templates/layout"
	"os"

	"github.com/a-h/templ"
)

const downloadDirectory = "./uploads/"

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
		f, err := os.Create("./uploads/" + fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		io.Copy(f, file)
		successComponent := components.Response("Upload success")
		successComponent.Render(r.Context(), w)
	})
	http.HandleFunc("/downloads", func(w http.ResponseWriter, r *http.Request) {
		entries, err := os.ReadDir(downloadDirectory)
		var resultList []string
		if err != nil {
			panic(err)
		}
		for _, item := range entries {
			resultList = append(resultList, item.Name())
		}
		itemList := components.ItemList(resultList)
		itemList.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:8080", nil)
}
