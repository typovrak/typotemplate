package main

import (
	"fmt"
	"typotemplate/config/app"
)

func main() {
	app.RequireEnv()

	// html.Minifier()

	fmt.Println("Hello World !")
}
