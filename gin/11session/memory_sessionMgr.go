package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
		rwLock:     sync.RWMutex{},
	}
	return sr
}

func (m *MemorySessionMgr) Init(addr string, options ...string) error {
	return nil
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//用uuid作为sessionid
	uuid := uuid.NewV4()
	//转成string
	sessionId := uuid.String()
	//创建一个session
	session = NewMemorySession(sessionId)
	//加入到大map
	m.sessionMap[sessionId] = session
	return
}

func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	s, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return s, nil
}
