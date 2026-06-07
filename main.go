package main

import (
	"templeruins/web"
	"templeruins/persistence/files"
)

func main() {
	files.EmailRead()
	web.Init()
}
