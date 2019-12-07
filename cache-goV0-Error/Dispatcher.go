package gocache

import (
	"time"
)

const (
	CLOSE = "close"
)

func newDefaultDispatcher() dispatcher {
	return dispatcher{queue:make(chan commons, DefaultCommonsChannelSize),
		stateCh:make(chan string, 1),
		}
}

func (d *dispatcher) start(li *Liquidator)  {
	go d.run(li)
}

func (d *dispatcher) run(li *Liquidator)  {
	ticker := time.NewTicker(DefaultCleatStep)
	for {
		select {
		case commons := <- d.queue: commons.fn()
			//fmt.Printf("run    %T , %p\n",commons.fn, &commons.fn)
		case state := <- d.stateCh:
			if state == CLOSE {
				close(d.queue)
			}
		case <- ticker.C:
			li.clearFunc()
		}
	}

}