package main

import (
	"controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("view")))
	http.Handle("/search", controller.CreatSearchHandler("view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
