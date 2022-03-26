#!/bin/zsh

command_not_found_handler() {
    pp $@
    echo "zsh: command not found: $1"
    return 127
}