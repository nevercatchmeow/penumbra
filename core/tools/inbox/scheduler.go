package inbox

const (
	defaultThroughput = 300
)

type scheduler int

func (slf scheduler) Schedule(fn func()) {
	go fn()
}

func (slf scheduler) Throughput() int {
	return int(slf)
}
