package builtin

import (
	"github.com/diiyw/pp/builtin/times"
)

type Command interface {
	Valid(args ...string) bool
	Run() string
}

var Commands = []Command{
	new(times.UnixToString),
	new(times.Unix),
	new(times.Now),
}
