package utils

import "sync"

type APIKeyManager struct {
	keys       []string
	keyUsage   map[string]int
	currentKey string
	mu         sync.Mutex
}

func NewAPIKeyManager(keys []string) *APIKeyManager {
	keyUsage := make(map[string]int)
	for _, key := range keys {
		keyUsage[key] = 0
	}

	return &APIKeyManager{
		keys:       keys,
		keyUsage:   keyUsage,
		currentKey: keys[0], // Start with the first key
		mu:         sync.Mutex{},
	}
}

func (m *APIKeyManager) GetNextKey() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.keyUsage[m.currentKey]++
	if m.keyUsage[m.currentKey] >= 50 {
		// Move to the next key
		currentIndex := -1
		for i, key := range m.keys {
			if key == m.currentKey {
				currentIndex = i
				break
			}
		}
		nextIndex := (currentIndex + 1) % len(m.keys)
		m.currentKey = m.keys[nextIndex]
		m.keyUsage[m.currentKey] = 0
	}

	return m.currentKey
}
