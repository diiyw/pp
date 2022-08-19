package times

import (
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
)

type Unix struct{}

func (u *Unix) Valid(args ...string) bool {
	return args[0] == "unix"
}

func (u *Unix) Run() string {
	return cast.ToString(carbon.Now().Unix())
}
