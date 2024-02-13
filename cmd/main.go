package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	runCmd(os.Args[1:])
}

func runCmd(args []string) {
	cmd, err := app.Parse(args)
	if err != nil {
		panic(fmt.Errorf("failed to parse cmd: %s", err.Error()))
	}
	log := logrus.New()
	switch cmd {
	case migrateCmd.FullCommand():
		runMigrate()
	case rpcCmd.FullCommand():
		runRpc(*port, log)
		return
	case gatewayCmd.FullCommand():
		runGateway(*port, *rpcAddr, log)
		return
	default:
		log.Fatal("unknown command")
	}

}
