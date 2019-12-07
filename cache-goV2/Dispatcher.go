package gocache

import (
	"time"
)

const (
	CLOSE = "close"
)

func newDefaultDispatcher() dispatcher {
	return dispatcher{queue:make([]chan *commons, 4,DefaultCommonsChannelSize),
		stateCh:make(chan string, 1)}
}

func (d *dispatcher) start(li *Liquidator)  {
	for i := 0; i < len(d.queue); i++ {
		go d.run(li, d.queue[i])
	}
}

func (d *dispatcher)  run(li *Liquidator,queue chan *commons)  {
	ticker := time.NewTicker(DefaultCleatStep)
	for {
		select {
		case commons := <- queue: (*commons.fn)(commons.data)
			//fmt.Println(commons)
		case state := <- d.stateCh:
			if state == CLOSE {
				close(queue)
				return
			}
		case <- ticker.C:
			li.clearFunc()
		}
	}

}
