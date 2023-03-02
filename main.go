package main

import "go_image/app"

func main() {
	application := &app.App{}
	application.Init()
	application.Run()
}
