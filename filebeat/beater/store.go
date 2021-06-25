package beater

import (
	"github.com/cjphaha/ebeats/libbeat/statestore"
	"time"
)

type filebeatStore struct {
	registry      *statestore.Registry
	storeName     string
	cleanInterval time.Duration
}

func (s *filebeatStore) Access() (*statestore.Store, error) {
	panic("implement me")
}

func openStateStore()  (*filebeatStore, error) {
	return &filebeatStore{
		registry:      statestore.NewRegistry(nil),
		//storeName:     info.Beat,
		//cleanInterval: cfg.CleanInterval,
	}, nil
}

func (s *filebeatStore) Close() {
	s.registry.Close()
}