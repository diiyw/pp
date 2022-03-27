package times

import (
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
)

type Unix int8

func (u *Unix) Valid(args ...string) bool {
	return args[1] == "unix"
}

func (u *Unix) Run() string {
	return cast.ToString(carbon.Now().Unix())
}
