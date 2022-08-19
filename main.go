package main

import (
	"fmt"
	"log"
	"os"

	"github.com/diiyw/pp/builtin"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "install" {
			install()
		}
	}
	for _, cmd := range builtin.Commands {
		if cmd.Valid(os.Args[1:]...) {
			result := cmd.Run()
			fmt.Println(result)
		}
	}
}

func install() {
	ppDir := os.Getenv("HOME") + "/.pp/"
	_, err := os.Stat(ppDir)
	if err != nil {
		if err = os.MkdirAll(ppDir, 0755); err != nil {
			log.Fatalln(err)
		}
	}
	handlerShell := ppDir + ".pp_profile"
	_, err = os.Stat(handlerShell)
	if err != nil {
		err = os.WriteFile(handlerShell, []byte("#!/bin/zsh\n\ncommand_not_found_handler() {\n    pp $@\n    echo \"zsh: command not found: $1\"\n    return 127\n}"), 0755)
		if err != nil {
			log.Fatalln("Initialize pp env failed: ", err)
		}
	}
	b, err := os.ReadFile("./pp")
	if err != nil {
		b, err = os.ReadFile("./pp.exe")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = os.WriteFile("/usr/local/bin/pp", b, 0755); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Install success! \nRun `. ~/.pp/.pp_profile` manually ")
}
