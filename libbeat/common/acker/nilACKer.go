package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type nilACKer struct{}

func (nilACKer) AddEvent(event beat.Event, published bool) {}
func (nilACKer) ACKEvents(n int)                           {}
func (nilACKer) Close()                                    {}