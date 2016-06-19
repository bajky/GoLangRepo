package main

import (
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", handleFunction)
	http.ListenAndServe(":8080",nil)
}

func handleFunction(writer http.ResponseWriter, request *http.Request){
	title:= "Hello World"
	template, _ := template.ParseFiles("main.html")
	template.Execute(writer, title)
}
