package commands

import (
	"github.com/diiyw/pp/commands/times"
)

type Command interface {
	Valid(str string) bool
	Run() string
}

var Commands = []Command{
	new(times.UnixToString),
	new(times.Unix),
}