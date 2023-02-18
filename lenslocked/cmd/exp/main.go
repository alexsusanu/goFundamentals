package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
		Bio  string
	}{
		Name: "John",
		Bio:  `<script>alert("you've been hacked")</script>`,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
