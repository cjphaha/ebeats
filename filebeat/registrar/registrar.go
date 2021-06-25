package registrar

import (
	"github.com/cjphaha/ebeats/filebeat/input/file"
	"github.com/cjphaha/ebeats/libbeat/statestore"
	"time"
)

type Registrar struct {
	// registrar event input and output
	Channel              chan []file.State
}

type successLogger interface {
	Published(n int) bool
}

type StateStore interface {
	Access() (*statestore.Store, error)
}

func New(stateStore StateStore, out successLogger, flushTimeout time.Duration) (*Registrar, error) {
	return nil, nil
}
