package main

import (
	"fmt"
	"typotemplate/config/app"
)

func main() {
	app.RequireEnv()

	// html.Tokenizer()

	fmt.Println("Hello World 4!")
}
