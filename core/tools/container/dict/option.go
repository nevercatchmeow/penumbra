package dict

type Option[K comparable, V any] func(*Dict[K, V])

func WithSource[K comparable, V any](source map[K]V) Option[K, V] {
	return func(m *Dict[K, V]) {
		m.data = source
	}
}
