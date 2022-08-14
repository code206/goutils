// 生产者，消费者模型
// 参考源码：https://github.com/qmhball/pc/blob/master/pc.go
// 实例说明：https://blog.csdn.net/qmhball/article/details/107234965
// 这个模型中，生产者一次性产生多份数据，由指定数量的消费者逐个完成，完成后关闭
// 这个模型中，无法不间断的产生数据给多个消费者完成，需要一次性产生
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
