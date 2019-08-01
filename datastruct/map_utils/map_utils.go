package map_utils

import "sync"

func EraseSyncMap(m *sync.Map) {
	m.Range(func(key interface{}, value interface{}) bool {
		m.Delete(key)
		return true
	})
}
