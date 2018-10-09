// gpoll is a go lang goroutine pool
// it depends on the maxSize and minSize configure
//
package gpool

import (
	"sync"
	"sync/atomic"
)

const (
	MinSize = 100
	MaxSize = 1000

	StateRunning = 1
	StateStoping = 2
	StateStoped  = 3
)

var (
	DefaultGpool *Gpool
)

type Gpool struct {
	maxSize int64
	minSize int64

	wg          sync.WaitGroup
	job         chan func()
	stop        chan struct{}
	currentSize int64
	state       int64
}

func (g *Gpool) start() {
	for i := int64(0); i < g.maxSize; i++ {
		go func() {
			atomic.AddInt64(&g.currentSize, 1)
			g.wg.Add(1)
			defer g.wg.Done()
			defer atomic.AddInt64(&g.currentSize, -1)
			for {
				select {
				case f := <-g.job:
					f()
					// did over
					if atomic.LoadInt64(&g.currentSize) < g.minSize {
						continue
					}
					if atomic.LoadInt64(&g.currentSize) > g.maxSize {
						return
					}
				case <-g.stop:
					return
				}
			}
		}()
	}
}

func New(min, max int64) *Gpool {
	g := new(Gpool)
	g.minSize = min
	g.maxSize = max
	if g.minSize <= 0 {
		g.minSize = 0
	}
	if g.minSize >= g.maxSize {
		g.maxSize = g.minSize + 10 // less 10 goroutine
	}

	g.stop = make(chan struct{}, 1)
	g.job = make(chan func(), g.maxSize)
	g.state = StateRunning

	go g.start()

	return g
}

func (g *Gpool) Stop() {
	atomic.StoreInt64(&g.state,StateStoping)
	size := atomic.LoadInt64(&g.currentSize)
	for i := int64(0) ; i < size ; i ++ {
		g.stop <- struct{}{}
	}
	g.wg.Wait()
	atomic.StoreInt64(&g.state,StateStoped)
}

func (g *Gpool) Go(f func()) {
	if atomic.LoadInt64(&g.state) == StateRunning {
		g.job <- f
	}
	// else
	f()
}

func Init() {
	DefaultGpool = New(MinSize, MaxSize)
}

func Go(f func()) {
	DefaultGpool.Go(f)
}

func Stop() {
	DefaultGpool.Stop()
}
