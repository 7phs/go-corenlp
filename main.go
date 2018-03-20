package main

import (
	"fmt"
	"time"

	"bitbucket.org/7phs/go-corenlp/corenlp/server"
	"bitbucket.org/7phs/tools/common"
)

const (
	defaultPort = 9000
)

func param() *server.CoreNlpParameter {
	return server.NewCoreNlpParameter("./local-server", common.CheckLocalPort(defaultPort))
}

func checkServer() error {
	param := param()
	fmt.Println(param.ToArgs())

	fmt.Println("Start LocalCoreNLP")
	if err := server.StartLocalServer(param); err != nil {
		fmt.Println("failed to start with", err)
		return err
	}

	duration := 20 * time.Second

	fmt.Println("Test server: Wait ", duration)
	time.Sleep(duration)

	fmt.Println("Test server: normal shutdown")

	if err := server.ShutdownLocalServer(); err != nil {
		fmt.Println("failed to shutdown with", err)
	}

	return nil
}

func main() {
	if err := checkServer(); err != nil {
		fmt.Println("failed to run a LocalCoreNLP server with", err)
	}
}
