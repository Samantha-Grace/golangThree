package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "<script>alert('hi')</script>"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
