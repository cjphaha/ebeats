package beater

import (
	"github.com/cjphaha/ebeats/filebeat/input/file"
	"github.com/cjphaha/ebeats/filebeat/registrar"
	"sync"
	"sync/atomic"
)

type finishedLogger struct {
	wg *eventCounter
}

func (f finishedLogger) Published(n int) bool {
	panic("implement me")
}

type registrarLogger struct {
	done chan struct{}
	ch   chan<- []file.State
}

type eventCounter struct {
	added *atomic.Value
	done  *atomic.Value
	count *atomic.Value
	wg    sync.WaitGroup
}

func newFinishedLogger(wg *eventCounter) *finishedLogger {
	return &finishedLogger{wg}
}

func newRegistrarLogger(reg *registrar.Registrar) *registrarLogger {
	return &registrarLogger{
		done: make(chan struct{}),
		ch:   reg.Channel,
	}
}