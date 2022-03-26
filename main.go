package main

import (
	"fmt"
	"os"

	"github.com/diiyw/pp/builtin"
)

func main() {
	for _, cmd := range builtin.Commands {
		if cmd.Valid(os.Args...) {
			result := cmd.Run()
			fmt.Println(result)
		}
	}
}
