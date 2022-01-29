package sync_tutorial

type Counter struct {
	count int
}

func (counter *Counter) Inc() {
	counter.count++
}

func (counter *Counter) Val() int {
	return counter.count
}
