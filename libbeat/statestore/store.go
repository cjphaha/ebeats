package statestore

import (
	"github.com/cjphaha/ebeats/libbeat/statestore/backend"
	"sync/atomic"
)

type sharedStore struct {
	reg      *Registry
	refCount atomic.Value

	name    string
	backend backend.Store
}

type Store struct {

}
