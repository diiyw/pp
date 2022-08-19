package system

import (
	"fmt"
	"os"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

type System struct {
	info string
}

func (s *System) Valid(args ...string) bool {
	if len(args) > 1 {
		s.info = args[1]
	}
	return args[0] == "system"
}

func (s *System) Run() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Description"})
	switch s.info {
	case "cpu":
		c, _ := cpu.Info()
		v := c[0]
		refValue := reflect.ValueOf(v)
		refType := reflect.TypeOf(v)
		for i := 0; i < refType.NumField(); i++ {
			fieldType := refType.Field(i)
			t.AppendSeparator()
			t.AppendRow(table.Row{fieldType.Name, fmt.Sprintf("%v", refValue.Field(i))})
		}
	default:
		platform, family, version, _ := host.PlatformInformation()
		t.AppendRow(table.Row{"Platform", platform})
		t.AppendSeparator()
		t.AppendRow(table.Row{"Family", family})
		t.AppendSeparator()
		t.AppendRow(table.Row{"Version", version})
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     "Name",
			WidthMin: 6,
			WidthMax: 18,
		},
		{
			Name:     "Description",
			WidthMin: 6,
			WidthMax: 64,
		},
	})
	t.Render()
	return ""
}
