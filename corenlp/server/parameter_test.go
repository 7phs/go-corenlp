package server

import (
	"fmt"
	"strings"
	"testing"

	"bitbucket.org/7phs/tools/monitoring"
)

func TestCoreNlpParameter_MonitoringParameter(t *testing.T) {
	var (
		monitoringParam monitoring.MonitoringParameter
		coreParam       = NewCoreNlpParameter("*", 9000)
	)

	if monitoringParam = coreParam; monitoringParam == nil {
		t.Error("failed to cast CoreNlpParameter to MonitoringParameter")
	}

}

func TestCoreNlpParameter(t *testing.T) {
	workDir := "./local-server"

	coreNlpParams := NewCoreNlpParameter(workDir, 9000)

	var (
		params monitoring.MonitoringParameter = coreNlpParams
	)

	if exist := params.WorkDir(); exist != workDir {
		t.Error("failed to check a work dir")
	}

	if exist := params.Command(); exist != "java" {
		t.Error("failed to check a command name")
	}

	if exist := params.RunningMode(); exist != monitoring.RepeatInfinity {
		t.Error("failed to check a running mode")
	}

	if exist := params.ParallelCount(); exist != 1 {
		t.Error("failed to check a parallel count")
	}
}

func TestCoreNlpParameter_ToArgs(t *testing.T) {
	params := NewCoreNlpParameter("./local-server", 9000)

	port := params.Port()
	if port <= 0 {
		t.Error("failed to init parameters. Port (#", port, ") is less than or equal 0, but shouldn'ts")
	}

	expected := fmt.Sprintf("-mx4g -cp * edu.stanford.nlp.pipeline.StanfordCoreNLPServer -port %d -timeout 15000", port)
	if exist := strings.Join(params.ToArgs(), " "); exist != expected {
		t.Error("failed to convert parameters to an arguments list. Got '", exist, ", but epected is '", expected, "'")
	}
}
