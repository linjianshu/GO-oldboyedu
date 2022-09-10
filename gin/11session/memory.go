package session

import (
	"errors"
	"sync"
)

// MemorySession 对象
type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	rwLock    sync.RWMutex
}

// NewMemorySession 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
		rwLock:    sync.RWMutex{},
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	value, ok := m.data[key]
	if !ok {
		return nil, errors.New("key not exists in session")
	}
	return value, nil
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() error {
	return nil
}
