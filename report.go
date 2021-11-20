package main

import (
	"flag"

	"github.com/spudtrooper/startupschool/startupschool"
)

func main() {
	flag.Parse()
	if err := startupschool.Report(); err != nil {
		panic(err.Error())
	}
}
