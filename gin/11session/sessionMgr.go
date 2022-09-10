package session

// SessionMgr 定义管理者 管理所有的session
type SessionMgr interface {
	// Init 初始化
	Init(addr string, options ...string) error
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
