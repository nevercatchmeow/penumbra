package dictionary

import (
	"encoding/json"
	"sync"
)

func NewDictionary[K comparable, V any](options ...Option[K, V]) *Dictionary[K, V] {
	dict := &Dictionary[K, V]{
		data: make(map[K]V),
	}
	for _, option := range options {
		option(dict)
	}
	return dict
}

type Dictionary[K comparable, V any] struct {
	mutex sync.RWMutex
	data  map[K]V
}

func (slf *Dictionary[K, V]) Set(k K, v V) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	slf.data[k] = v
}

func (slf *Dictionary[K, V]) Get(k K) (V, bool) {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	val, ok := slf.data[k]
	return val, ok
}

func (slf *Dictionary[K, V]) Delete(k K) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	delete(slf.data, k)
}

func (slf *Dictionary[K, V]) Exist(key K) bool {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	_, exist := slf.data[key]
	return exist
}

func (slf *Dictionary[K, V]) Len() int {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return len(slf.data)
}

func (slf *Dictionary[K, V]) ForEach(f func(K, V)) {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	for key, val := range slf.data {
		f(key, val)
	}
}

func (slf *Dictionary[K, V]) Keys() []K {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	keys := make([]K, 0, len(slf.data))
	for key := range slf.data {
		keys = append(keys, key)
	}
	return keys
}

func (slf *Dictionary[K, V]) Values() []V {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	values := make([]V, 0, len(slf.data))
	for _, value := range slf.data {
		values = append(values, value)
	}
	return values
}

func (slf *Dictionary[K, V]) Map() map[K]V {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	source := make(map[K]V)
	for k, v := range slf.data {
		source[k] = v
	}
	return source
}

func (slf *Dictionary[K, V]) MarshalJSON() ([]byte, error) {
	m := slf.Map()
	return json.Marshal(m)
}

func (slf *Dictionary[K, V]) UnmarshalJSON(bytes []byte) error {
	source := make(map[K]V)
	slf.mutex.Lock()
	slf.mutex.Unlock()
	if err := json.Unmarshal(bytes, &source); err != nil {
		return err
	}
	slf.data = source
	return nil
}
