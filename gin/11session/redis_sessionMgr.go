package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	//redis地址
	addr string
	//密码
	password string
	//连接池
	pool *redis.Pool
	//锁
	rwLock sync.RWMutex
	//大map
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) error {
	//若有其他参数
	if len(options) > 0 {
		r.password = options[0]
	}
	//创建连接池
	r.pool = myPool(addr, r.password)
	r.addr = addr
	return nil
}

func myPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			//若有密码 , 判断
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		//链接测试 , 开发时写
		//上线测试
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
		MaxIdle:         64,
		MaxActive:       1000,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	session = NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists ")
		return
	}
	return
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}
