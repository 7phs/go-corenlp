package server

import (
	"strconv"
	"strings"

	"bitbucket.org/7phs/tools/monitoring"
)

const (
	defaultTimeout   = 15000
	commandName      = "java"
	classPath        = "*"
	className        = "edu.stanford.nlp.pipeline.StanfordCoreNLPServer"
	startedLogPrefix = "[main] INFO CoreNLP - StanfordCoreNLPServer listening at"
)

type ServerParameter interface {
	monitoring.MonitoringParameter

	Scheme() string
	Host() string
	Port() int
}

type CoreNlpParameter struct {
	command       string
	runningMode   int32
	parallelCount int32

	workDir string

	classPath string
	className string
	port      int
}

func NewCoreNlpParameter(workDir string, port int) *CoreNlpParameter {
	return &CoreNlpParameter{
		command:       commandName,
		runningMode:   monitoring.RepeatInfinity,
		parallelCount: 1,

		workDir: workDir,

		classPath: classPath,
		className: className,
		port:      port,
	}
}

func (o *CoreNlpParameter) Scheme() string {
	return "http"
}

func (o *CoreNlpParameter) Host() string {
	return "localhost"
}

func (o *CoreNlpParameter) Port() int {
	return o.port
}

func (o *CoreNlpParameter) WorkDir() string {
	return o.workDir
}

func (o *CoreNlpParameter) Command() string {
	return o.command
}

func (o *CoreNlpParameter) ToArgs() []string {
	return []string{
		"-mx4g",
		"-cp", o.classPath,
		o.className,
		"-port", strconv.Itoa(o.port),
		"-timeout", strconv.Itoa(defaultTimeout),
	}
}

func (o *CoreNlpParameter) RunningMode() int32 {
	return o.runningMode
}

func (o *CoreNlpParameter) ParallelCount() int32 {
	return o.parallelCount
}

func (o *CoreNlpParameter) StdErrIsOk() bool {
	return true
}

func (o *CoreNlpParameter) checkStartLine(line string) bool {
	return strings.HasPrefix(line, startedLogPrefix)
}

func (o *CoreNlpParameter) CheckStartLine() func(string) bool {
	return o.checkStartLine
}
