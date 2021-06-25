package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type eventDataACKer struct {
	beat.ACKer
	fn func(acked int, data []interface{})
}

func (a *eventDataACKer) AddEvent(event beat.Event, published bool) {

}

func (a *eventDataACKer) onACK(acked, total int) {

}