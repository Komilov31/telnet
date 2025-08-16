package main

import (
	"github.com/Komilov31/telnet/internal/flags"
	"github.com/Komilov31/telnet/internal/telnet"
)

func main() {
	flags := flags.Parse()
	telnet := telnet.New(flags)

	telnet.ProcessProgram()
}
