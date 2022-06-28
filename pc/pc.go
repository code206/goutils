package pc

import "sync"

type Pool struct {
	ConsumerNumber int
	ChannelLength  int
	Tasks          chan interface{}
	Consumer       func(interface{})
	Producer       func(chan<- interface{})
	wg             sync.WaitGroup
}

//
func NewPool(channelLength int, consumerNumber int, produceFunc func(chan<- interface{}), consumeFunc func(interface{})) *Pool {
	pool := &Pool{
		ChannelLength:  channelLength,
		Tasks:          make(chan interface{}, channelLength),
		ConsumerNumber: consumerNumber,
		Producer:       produceFunc,
		Consumer:       consumeFunc,
	}
	return pool
}

func (p *Pool) producerDispatch() {
	defer close(p.Tasks)
	p.Producer(p.Tasks)
}

func (p *Pool) consumerDispatch() {
	for i := 0; i < p.ConsumerNumber; i++ {
		p.wg.Add(1)
		go func(i int) {
			defer p.wg.Done()
			for task := range p.Tasks {
				p.Consumer(task)
			}
		}(i)
	}

	p.wg.Wait()
}

func (p *Pool) Run() {
	go p.producerDispatch()
	p.consumerDispatch()
}
