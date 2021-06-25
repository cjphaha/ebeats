package monitoring

import "sync/atomic"

func NewInt(r *Registry, name string) *atomic.Value{
	return &atomic.Value{}
}