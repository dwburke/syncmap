package syncmap

import (
	"sync"
)

type StrSyncMap struct {
	sync.RWMutex
	internal map[string]interface{}
}

func NewStrSyncMap() *StrSyncMap {
	return &StrSyncMap{
		internal: make(map[string]interface{}),
	}
}

func (rm *StrSyncMap) Load(key string) (value interface{}, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *StrSyncMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *StrSyncMap) Store(key string, value interface{}) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}
