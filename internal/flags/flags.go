package flags

import (
	"log"

	"github.com/pborman/getopt/v2"
)

const (
	expectedArgsCount = 2
)

type Flags struct {
	Host    string
	Port    string
	Timeout int
}

func Parse() *Flags {
	timeout := getopt.IntLong("timeout", 't', 10, "timeout for the connection in seconds")
	getopt.Parse()

	flags := Flags{
		Timeout: *timeout,
	}

	args := getopt.Args()
	if len(args) < expectedArgsCount {
		log.Fatal("host and(or) port are not specified")
	}

	flags.Host = args[0]
	flags.Port = args[1]

	return &flags
}
