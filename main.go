package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"

	"github.com/diiyw/pp/builtin"
)

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	ppDir := u.HomeDir + "/.pp/"
	_, err = os.Stat(ppDir)
	if err != nil {
		if err = os.MkdirAll(ppDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	handlerShell := ppDir + ".pp_profile"
	_, err = os.Stat(handlerShell)
	if err == nil {
		return
	}
	err = os.WriteFile(handlerShell, []byte("#!/bin/zsh\n\ncommand_not_found_handler() {\n    pp $@\n    echo \"zsh: command not found: $1\"\n    return 127\n}"), 0755)
	if err != nil {
		log.Fatal("Initialize pp env failed: ", err)
	}
	cmd := exec.Command("source", handlerShell)
	cmd.Stdout = nil
	cmd.Stdin = nil
	cmd.Stderr = nil
	if err = cmd.Run(); err != nil {
		log.Fatal("source failed: ", err)
	}
}

func main() {
	for _, cmd := range builtin.Commands {
		if cmd.Valid(os.Args...) {
			result := cmd.Run()
			fmt.Println(result)
		}
	}
}
