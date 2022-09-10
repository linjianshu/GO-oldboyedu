package session

import "fmt"

var sessionMgr SessionMgr

// Init 指出中间件让用户选择使用哪个版本
func Init(provider string, addr string, options ...string) error {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持")
	}
	err := sessionMgr.Init(addr, options...)
	return err
}
