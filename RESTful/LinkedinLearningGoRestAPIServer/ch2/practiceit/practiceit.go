package main

import (
	"github.com/gabuh/backend"
)

func main() {
	a := backend.App{}
	a.Port = ":8080"
	a.Initialize()
	a.Run()

}
