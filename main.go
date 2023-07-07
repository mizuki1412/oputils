package main

import (
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"waster/cmd"
)

func main() {
	initkit.LoadConfig()
	cmd.Execute()
}
