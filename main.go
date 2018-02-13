package main

import (
	"log"
	"os"

	logs "github.com/appscode/go/log/golog"
	"github.com/sanjid133/crd-controller/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewCmdRun().Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
