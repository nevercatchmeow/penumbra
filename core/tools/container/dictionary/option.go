package dictionary

type Option[K comparable, V any] func(*Dictionary[K, V])

func WithSource[K comparable, V any](source map[K]V) Option[K, V] {
	return func(m *Dictionary[K, V]) {
		m.data = source
	}
}
