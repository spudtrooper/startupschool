package main

import (
	"flag"

	"github.com/spudtrooper/startupschool/startupschool"
)

var (
	data = flag.String("data", "data", "directory to store data")
)

func main() {
	flag.Parse()
	if err := startupschool.Report(*data); err != nil {
		panic(err.Error())
	}
}
