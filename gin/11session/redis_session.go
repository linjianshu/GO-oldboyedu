package session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	//设置session , 可以先放在内存的map中
	//批量导入redis,提升性能
	sessionMap map[string]interface{}
	//读写锁
	rwLock sync.RWMutex
	//记录内存中的map是否被操作
	flag int
}

//用常量定义状态
const (
	// SessionFlagNone 内存数据没变化
	SessionFlagNone = iota
	// SessionFlagModify 有变化
	SessionFlagModify
)

func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
	return s
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	//加锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//设置值
	r.sessionMap[key] = value
	//标记记录
	r.flag = SessionFlagModify
	return
}

func (r *RedisSession) Save() (err error) {
	//加锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	//若数据没变, 不需要存
	if r.flag != SessionFlagModify {
		return
	}

	//内存中的session进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}

	//获取redis连接
	conn := r.pool.Get()
	//保存kv
	_, err = conn.Do("SET", r.sessionId, string(data))
	//改状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return

}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	//加锁
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	//先判断内存
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists")
	}
	return
}

func (r *RedisSession) Del(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}

//从redis里再次加载
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("Get", r.sessionId)
	if err != nil {
		return
	}

	//转字符串
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	//取到的东西 , 反序列化到内存的map中
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return err
	}
	return

}
