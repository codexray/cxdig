package main

import (
	"flag"
	"fmt"
)

// Application used to check the command line arguments received when invoked from the "exec" command
func main() {
	id := "<none>"
	name := "<none>"

	flag.StringVar(&id, "id", "", "commit ID")
	flag.StringVar(&name, "name", "", "project name")
	flag.Parse()

	rtn := "path = "
	if len(flag.Args()) != 1 {
		rtn += "<none>"
	} else {
		rtn += flag.Args()[0]
	}

	rtn += " : id = " + id
	rtn += " : name = " + name

	fmt.Println(rtn)
}
