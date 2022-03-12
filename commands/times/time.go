package times

import (
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
)

type Unix int8

func (u *Unix) Valid(raw string) bool {
	return raw == "unix"
}

func (u *Unix) Run() string {
	return cast.ToString(carbon.Now().Unix())
}
