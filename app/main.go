package main

import (
	"fmt"
	"typotemplate/config/app"
	"typotemplate/html"
)

func main() {
	app.RequireEnv()

	html.Tokenizer()

	fmt.Println("Hello World 4!")
}
