package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

type trackingACKer struct {
	fn func(acked, total int)
}

func (a *trackingACKer) AddEvent(event beat.Event, published bool) {

}

func (a *trackingACKer) ACKEvents(n int) {

}

func (a *trackingACKer) Close() {

}

func TrackingCounter(fn func(acked, total int)) beat.ACKer {
	a := &trackingACKer{fn: fn}
	return a
}