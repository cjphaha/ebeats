package beater

import (
	"github.com/cjphaha/ebeats/filebeat/input/file"
	"github.com/cjphaha/ebeats/libbeat/beat"
	"github.com/cjphaha/ebeats/libbeat/common/acker"
)

type statefulLogger interface {
	Published(states []file.State)
}

type statelessLogger interface {
	Published(c int) bool
}


// 这里statelessLogger
// statefulLogger是 finishedLogger
// statefulOut是
func eventACKer(statelessOut statelessLogger, statefulOut statefulLogger) beat.ACKer {
	return acker.EventPrivateReporter(func(_ int, data []interface{}) {
		//	这里是函数的内容
		stateless := 0
		states := make([]file.State, len(data))

		for _, datum := range data {
			if datum == nil {
				stateless++
				continue
			}
			// 判断一下是不是file.State类型
			// 如果传过来的数据是file.State类型的，就states ++ ，如果不是的就 stateless ++
			st, ok := datum.(file.State)
			if !ok {
				stateless ++
				continue
			}

			states = append(states, st)
		}

		if len(states) > 0 {
			statefulOut.Published(states)
		}

		if stateless > 0 {
			statelessOut.Published(stateless)
		}
	})
}