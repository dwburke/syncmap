package syncmap

import (
	"sync"
)

type IntSyncMap struct {
	sync.RWMutex
	internal map[int]interface{}
}

func NewIntSyncMap() *IntSyncMap {
	return &IntSyncMap{
		internal: make(map[int]interface{}),
	}
}

func (rm *IntSyncMap) Load(key int) (value interface{}, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *IntSyncMap) Delete(key int) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *IntSyncMap) Store(key int, value interface{}) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}
