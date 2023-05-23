package main

import (
	"fmt"
	"github.com/v2fly/v2ray-core/v5/main/commands"
	"github.com/v2fly/v2ray-core/v5/main/commands/base"
	_ "github.com/v2fly/v2ray-core/v5/main/distro/all"
	"os"
)

func main() {
	for _, arg := range os.Args {
		if arg == "--version" {
			fmt.Println("iseki version")
			return
		}
		if arg == "--test" || arg == "-test" {
			return
		}
	}

	commands.ExecuteRun()
}

func runIsTheFirst(i, j *base.Command) bool {
	left := i.Name()
	right := j.Name()
	if left == "run" {
		return true
	}
	if right == "run" {
		return false
	}
	return left < right
}
