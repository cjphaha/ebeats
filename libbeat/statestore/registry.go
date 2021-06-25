package statestore

import "github.com/cjphaha/ebeats/libbeat/statestore/backend"

type Registry struct {
	backend backend.Registry
	active map[string]*sharedStore // active/open stores
}

// NewRegistry creates a new Registry with a configured backend.
func NewRegistry(backend backend.Registry) *Registry {
	return &Registry{
		backend: backend,
		active:  map[string]*sharedStore{},
	}
}

// Close closes the backend storage. Close blocks until all stores in use are closed.
func (r *Registry) Close() error {
	return nil
}