package gocache

import (
	"time"
)

const (
	CLOSE = "close"
)

func newDefaultDispatcher() dispatcher {
	return dispatcher{queue:make(chan *commons,DefaultCommonsChannelSize),
		stateCh:make(chan string, 1)}
}

func (d *dispatcher) start(li *Liquidator)  {
	go d.run(li)
}

func (d *dispatcher) doClose() {
	close(d.queue)
	for commons := range d.queue {
		(*commons.fn)(commons.data)
	}
}

func (d *dispatcher) run(li *Liquidator)  {
	ticker := time.NewTicker(DefaultCleatStep)
	for {
		select {
		case commons := <- d.queue: (*commons.fn)(commons.data)
		case state := <- d.stateCh:
			if state == CLOSE {
				d.doClose()
				ticker.Stop()
				return
			}
		case <- ticker.C:
			li.clearFunc()
		}
	}
}



