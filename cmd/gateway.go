package main

import (
	"github.com/eonianmonk/go-blog/http"
	"github.com/sirupsen/logrus"
)

func runGateway(port string, rpcEp string, log *logrus.Logger) {
	log.WithFields(logrus.Fields{"port": port}).Info("starting gateway")
	http.Run(port, rpcEp, log)
}
