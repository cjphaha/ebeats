package acker

import "github.com/cjphaha/ebeats/libbeat/beat"

func Counting(fn func(n int)) beat.ACKer {
	return TrackingCounter(func(_ int, total int) {
		fn(total)
	})
}

// EventPrivateReporter 报告来自已发布或删除的所有事件的所有私有字段。
func EventPrivateReporter(fn func(acked int, data []interface{})) beat.ACKer {
	a := &eventDataACKer{fn: fn}
	a.ACKer = TrackingCounter(a.onACK)
	return a
}

func LastEventPrivateReporter(fn func(acked int, data interface{})) beat.ACKer {
	ignored := 0
	return EventPrivateReporter(func(acked int, data []interface{}) {
		for i := len(data) - 1; i >= 0; i-- {
			if d := data[i]; d != nil {
				fn(ignored+acked, d)
				ignored = 0
				return
			}
		}

		// complete batch has been ignored due to missing data -> add count
		ignored += acked
	})
}