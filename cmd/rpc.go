package main

import (
	"os"

	"github.com/eonianmonk/go-blog/rpc"
	"github.com/sirupsen/logrus"
)

func runRpc(port string, log *logrus.Logger) {
	log.WithFields(logrus.Fields{"port": port}).Info("starting gateway")
	dbconnStr := os.Getenv(dbenv)
	if dbconnStr == "" {
		panic("no db connection string specified in rpc")
	}
	rpc.NewServer(port, dbconnStr, log)
}
