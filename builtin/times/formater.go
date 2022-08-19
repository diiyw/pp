package times

import (
	"regexp"

	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
)

type UnixToString struct {
	c *carbon.Carbon
}

func (u *UnixToString) Valid(args ...string) bool {
	if len(args[0]) == 10 {
		if matched, _ := regexp.MatchString(`^\d+$`, args[0]); matched {
			c, err := carbon.CreateFromTimestamp(cast.ToInt64(args[0]), "Asia/Shanghai")
			if err != nil {
				return false
			}
			u.c = c
			return true
		}
	}
	return false
}

func (u *UnixToString) Run() string {
	return u.c.Format(carbon.DefaultFormat)
}
