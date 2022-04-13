package main

import (
	"flag"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/startupschool/graphql"
)

var (
	operationName = flag.String("operation_name", "", "graphql OperationName")
	query         = flag.String("query", "", "graphql Query")
)

func main() {
	flag.Parse()
	check.Check(*operationName != "", check.CheckMessage("--operation_name required"))
	check.Check(*query != "", check.CheckMessage("--query required"))
	api, err := graphql.MakeAPIFromFlags()
	check.Err(err)
	check.Err(api.Query(*operationName, *query))
}
