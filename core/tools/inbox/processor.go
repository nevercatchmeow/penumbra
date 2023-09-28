package inbox

type Processor[V any] interface {
	Invoke(envelopes []V)
}
