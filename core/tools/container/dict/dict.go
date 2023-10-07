package dict

import (
	"encoding/json"
	"sync"
)

func New[K comparable, V any](options ...Option[K, V]) *Dict[K, V] {
	dict := &Dict[K, V]{
		data: make(map[K]V),
	}
	for _, option := range options {
		option(dict)
	}
	return dict
}

type Dict[K comparable, V any] struct {
	mutex sync.RWMutex
	data  map[K]V
}

func (slf *Dict[K, V]) Set(k K, v V) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	slf.data[k] = v
}

func (slf *Dict[K, V]) Get(k K) (V, bool) {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	val, ok := slf.data[k]
	return val, ok
}

func (slf *Dict[K, V]) Delete(k K) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	delete(slf.data, k)
}

func (slf *Dict[K, V]) Exist(key K) bool {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	_, exist := slf.data[key]
	return exist
}

func (slf *Dict[K, V]) Len() int {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return len(slf.data)
}

func (slf *Dict[K, V]) ForEach(f func(K, V)) {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	for key, val := range slf.data {
		f(key, val)
	}
}

func (slf *Dict[K, V]) Keys() []K {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	keys := make([]K, 0, len(slf.data))
	for key := range slf.data {
		keys = append(keys, key)
	}
	return keys
}

func (slf *Dict[K, V]) Values() []V {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	values := make([]V, 0, len(slf.data))
	for _, value := range slf.data {
		values = append(values, value)
	}
	return values
}

func (slf *Dict[K, V]) Map() map[K]V {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	source := make(map[K]V)
	for k, v := range slf.data {
		source[k] = v
	}
	return source
}

func (slf *Dict[K, V]) MarshalJSON() ([]byte, error) {
	m := slf.Map()
	return json.Marshal(m)
}

func (slf *Dict[K, V]) UnmarshalJSON(bytes []byte) error {
	source := make(map[K]V)
	slf.mutex.Lock()
	slf.mutex.Unlock()
	if err := json.Unmarshal(bytes, &source); err != nil {
		return err
	}
	slf.data = source
	return nil
}
