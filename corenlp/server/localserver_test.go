package server

import (
	"testing"

	"bitbucket.org/7phs/tools/common"
)

func TestLocalCoreNLP(t *testing.T) {
	if err := StartLocalServer(NewCoreNlpParameter("../../local-server", common.CheckLocalPort(9000))); err != nil {
		t.Error("failed to start coreNLP server with", err)
	}

	if err := ShutdownLocalServer(); err != nil {
		t.Error("failed to shutdown coreNLP server with", err)
	}
}
