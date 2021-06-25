package beater

import (
	"github.com/cjphaha/ebeats/filebeat/registrar"
	"github.com/cjphaha/ebeats/libbeat/monitoring"
)

type Filebeat struct {

}

func (fb *Filebeat) Run() error {
	// count active events for waiting on shutdown
	wgEvents := &eventCounter{
		count: monitoring.NewInt(nil, "filebeat.events.active"),
		added: monitoring.NewInt(nil, "filebeat.events.added"),
		done:  monitoring.NewInt(nil, "filebeat.events.done"),
	}

	finishedLogger := newFinishedLogger(wgEvents)

	stateStore, err := openStateStore()
	if err != nil {
		return err
	}
	defer stateStore.Close()



	registrar, err := registrar.New(stateStore, finishedLogger, 0)

	registrarChannel := newRegistrarLogger(registrar)

	
}