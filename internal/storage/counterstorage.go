package storage

import "sync"

type CounterStorage struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounterStorage() *CounterStorage {
	counters := make(map[string]int)

	return &CounterStorage{
		counters: counters,
	}
}

func (s *CounterStorage) Increment(key string) {
	s.mu.Lock()

	s.counters[key]++
	s.mu.Unlock()
}

func (s *CounterStorage) GetCounter(key string) int {
	s.mu.Lock()

	defer s.mu.Unlock()
	return s.counters[key]
}

func (s *CounterStorage) GetCounters() map[string]int {
	return s.counters
}
