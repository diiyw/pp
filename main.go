package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/creack/pty"
	"github.com/diiyw/pp/commands"
	"golang.org/x/term"
)

func main() {
	// Create arbitrary command.
	c := exec.Command("zsh")

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		log.Fatalln(err)
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.
	go func() {
		io.Copy(ptmx, os.Stdin)
	}()
	output := make([]byte, 1024)
	for {
		n, err := ptmx.Read(output)
		if err != nil {
			break
		}
		rawString := string(output[:n])
		raw := strings.Split(rawString, ":")
		badCmd := ""
		if len(raw) == 3 {
			if raw[2] == " command not found" {
				badCmd = strings.TrimSpace(raw[1])
			}
			if raw[1] == " command not found" {
				badCmd = strings.TrimSpace(raw[2])
			}
			if hook(badCmd) {
				continue
			}
		}
		os.Stdout.Write(output[:n])
	}
}

func hook(str string) bool {
	for _, cmd := range commands.Commands {
		if cmd.Valid(str) {
			result := cmd.Run()
			os.Stdout.WriteString(result)
			return true
		}
	}
	return false
}
