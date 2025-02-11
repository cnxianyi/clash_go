package main

import (
	"clash_go/config"
	"clash_go/tp"
)

func main() {
	tp.SetTagMap()
	config.Setup()
}
