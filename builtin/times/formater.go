package times

import (
	"regexp"

	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
)

type UnixToString struct {
	c *carbon.Carbon
}

func (u *UnixToString) Valid(raw string) bool {
	if len(raw) == 10 {
		if matched, _ := regexp.MatchString(`^\d+$`, raw); matched {
			c, err := carbon.CreateFromTimestamp(cast.ToInt64(raw), "Asia/Shanghai")
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
