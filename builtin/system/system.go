package system

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
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
	switch s.info {
	case "cpu":
		t.AppendHeader(table.Row{"Name", "Description"})
		c, _ := cpu.Info()
		v := c[0]
		refValue := reflect.ValueOf(v)
		refType := reflect.TypeOf(v)
		for i := 0; i < refType.NumField(); i++ {
			fieldType := refType.Field(i)
			value := fmt.Sprintf("%v", refValue.Field(i))
			if fieldType.Name == "CPU" {
				value = strconv.Itoa(len(c))
			}
			if fieldType.Name == "Cores" {
				perCores, _ := strconv.ParseInt(value, 10, 32)
				value = strconv.Itoa(len(c) * int(perCores))
			}
			t.AppendRow(table.Row{fieldType.Name, value})
		}
	case "disk":

	default:
		t.AppendHeader(table.Row{"Name", "Description"})
		platform, family, version, _ := host.PlatformInformation()
		t.AppendRow(table.Row{"Family", family})
		t.AppendRow(table.Row{"Version", version})
		t.AppendRow(table.Row{"Platform", platform})
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

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
