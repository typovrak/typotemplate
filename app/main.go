package main

import (
	"fmt"
	"typotemplate/config/app"
)

func main() {
	app.RequireEnv()

	fmt.Println("Hello World 4!")
}
