package server

import (
	"fmt"
	"os"
	"testing"

	"bitbucket.org/7phs/tools/common"
)

func TestMain(m *testing.M) {
	err := StartLocalServer(NewCoreNlpParameter("../../local-server", common.CheckLocalPort(9000)))
	if err != nil {
		fmt.Println("failed to start local server with", err)
	}

	result := m.Run()

	ShutdownLocalServer()

	os.Exit(result)
}
