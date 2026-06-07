package main

import (
	"templeruins/persistence/files"
	"templeruins/web"
)

func main() {
	files.EmailRead()
	web.Init()
}
