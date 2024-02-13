package main

import "github.com/alecthomas/kingpin"

var (
	app  = kingpin.New("blog", "service to run micro blog")
	port = app.Flag("port", "svc port").Default(":8080").String()

	gatewayCmd = app.Command("run-gateway", "run http gateway")
	rpcAddr    = gatewayCmd.Flag("rpc", "rpc addr").Default(":8081").Required().String()

	rpcCmd = app.Command("run-rpc", "run rpc")

	migrateCmd = app.Command("migrate", "apply migrations to db")
)
