package main

import (
	_ "github.com/zfinn/t4cobra/init" // init must be imported first

	"github.com/zfinn/t4cobra/cmd"
)

func main() {
	cmd.Execute()
}
