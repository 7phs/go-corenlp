package server

import (
	"context"
	"sync"
	"time"

	"bitbucket.org/7phs/tools/monitoring"
	"github.com/pkg/errors"
)

var (
	localInstance *localCoreNLPInstance
)

func init() {
	(&sync.Once{}).Do(func() {
		localInstance = &localCoreNLPInstance{}
	})
}

type localCoreNLPInstance struct {
	server *localCoreNLP
}

func StartLocalServer(parameter ServerParameter) error {
	(&sync.Once{}).Do(func() {
		localInstance.server = &localCoreNLP{
			ServerParameter: parameter,
			cmd:             monitoring.NewMonitoring(parameter),
		}

		localInstance.server.Start()
	})

	return localInstance.server.HasError()
}

func LocalServer() *localCoreNLP {
	return localInstance.server
}

func ShutdownLocalServer() error {
	(&sync.Once{}).Do(func() {
		localInstance.server.Shutdown()
		localInstance.server.Wait()
	})

	// TODO: make synchronized
	if localInstance.server != nil {
		return localInstance.server.HasError()
	} else {
		return errors.New("local server wasn't starting")
	}
}

type localCoreNLP struct {
	ServerParameter

	cmd *monitoring.Monitoring
}

func (o *localCoreNLP) Start() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	o.cmd.Start(ctx)

	return o.HasError()
}

func (o *localCoreNLP) Shutdown() error {
	o.cmd.Kill(context.Background())

	return o.HasError()
}

func (o *localCoreNLP) HasError() error {
	return o.cmd.HasError()
}

func (o *localCoreNLP) Wait() {
	o.cmd.Wait()
}
