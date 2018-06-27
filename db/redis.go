package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

type RedisPool interface {
	Get() redis.Conn
	Close() error
}

func NewRedisPool(addr, password string, db ...int) RedisPool {
	opts := make([]redis.DialOption, 0)
	opts = append(opts, redis.DialConnectTimeout(5*time.Second))
	if password != "" {
		opts = append(opts, redis.DialPassword(password))
	}
	if len(db) != 0 {
		opts = append(opts, redis.DialDatabase(db[0]))
	}
	// FIXME: NewPool 将会失效
	pool = redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp4", addr, opts...)
		}, 20)
	conn := pool.Get()
	_, err := conn.Do("PING")
	if err != nil {
		panic(err)
	}
	return pool
}
