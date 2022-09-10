package session

type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (value interface{}, err error)
	Del(key string) error
	Save() error
}
