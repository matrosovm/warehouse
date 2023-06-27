package helpers

import "sync"

// SyncMapUint64Bool ...
type SyncMapUint64Bool struct {
	sync.RWMutex
	data map[uint64]bool
}

// NewSyncMapUint64Bool ...
func NewSyncMapUint64Bool(capacity int) *SyncMapUint64Bool {
	return &SyncMapUint64Bool{
		data: make(map[uint64]bool, capacity),
	}
}

// Load ...
func (m *SyncMapUint64Bool) Load(key uint64) (bool, bool) {
	m.RLock()
	defer m.RUnlock()
	val, ok := m.data[key]
	return val, ok
}

// Store ...
func (m *SyncMapUint64Bool) Store(key uint64, value bool) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

// GetData ...
func (m *SyncMapUint64Bool) GetData() map[uint64]bool {
	return m.data
}
