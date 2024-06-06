package main

import (
	"GO/src"
)

func main() {
	serverInstance := src.Server()
	serverInstance.Run()
}
