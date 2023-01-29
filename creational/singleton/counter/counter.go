package counter

import "sync"

type Counter struct {
	value int
}

var (
	lock = &sync.Mutex{}
	// once      sync.Once
	singleton *Counter
)

func GetInstance() *Counter {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			singleton = &Counter{}
		}
	}

	return singleton
}

func (c *Counter) Incr() {
	lock.Lock()
	defer lock.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	return c.value
}
