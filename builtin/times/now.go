package times

import (
	"github.com/uniplaces/carbon"
)

type Now int8

func (n *Now) Valid(args ...string) bool {
	return args[0] == "now"
}

func (n *Now) Run() string {
	c, err := carbon.NowInLocation("Asia/Shanghai")
	if err != nil {
		return ""
	}
	return c.Format(carbon.DefaultFormat)
}
