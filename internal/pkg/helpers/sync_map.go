package helpers

import "sync"

// SyncMapUint64Bool ...
type SyncMapUint64Bool struct {
	data map[uint64]bool
	mu   sync.RWMutex
}

// NewSyncMapUint64Bool ...
func NewSyncMapUint64Bool(capacity int) *SyncMapUint64Bool {
	return &SyncMapUint64Bool{
		data: make(map[uint64]bool, capacity),
		mu:   sync.RWMutex{},
	}
}

// Load ...
func (m *SyncMapUint64Bool) Load(key uint64) (bool, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[key]
	return val, ok
}

// Store ...
func (m *SyncMapUint64Bool) Store(key uint64, value bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// GetData ...
func (m *SyncMapUint64Bool) GetData() map[uint64]bool {
	return m.data
}
