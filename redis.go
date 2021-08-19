package goutils

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

func RedisNewPool(server, password string, db, maxIdle int) *redis.Pool {
	if maxIdle <= 0 {
		maxIdle = 1
	}
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		// TestOnBorrow: func(c redis.Conn, t time.Time) error {
		// 	_, err := c.Do("PING")
		// 	return err
		// },
	}
}

// expired 时效，单位：秒
func SetValueEX(pool *redis.Pool, key string, value interface{}, expired int) error {
	if pool == nil || len(key) <= 0 {
		return errors.New("invalid param")
	}
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", key, expired, value)
	return err
}

func Incr(pool *redis.Pool, key string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return IncrConn(conn, key)
}

func IncrConn(conn redis.Conn, key string) (int64, error) {
	return redis.Int64(conn.Do("INCR", key))
}

func IncrEX(pool *redis.Pool, key string, expired int) error {
	conn := pool.Get()
	defer conn.Close()
	return IncrEXConn(conn, key, expired)
}

func IncrEXConn(conn redis.Conn, key string, expired int) error {
	err := conn.Send("INCR", key)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, expired)
	return err
}

func DescEX(pool *redis.Pool, key interface{}, expired int) error {
	conn := pool.Get()
	defer conn.Close()
	return DescEXConn(conn,key,expired)
}

func DescEXConn(conn redis.Conn, key interface{}, expired int) error {
	err := conn.Send("DECR", key)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, expired)
	return err
}

func Exists(cache *redis.Pool, key string) (bool, error) {
	conn := cache.Get()
	defer conn.Close()
	return ExistsConn(conn, key)
}

func ExistsConn(conn redis.Conn, key string) (bool, error) {
	return redis.Bool(conn.Do("EXISTS", key))
}

// expired 时效，单位秒
func SendExpireConn(conn redis.Conn, key string, expired int) error {
	if conn == nil || len(key) <= 0 {
		return errors.New("invalid param")
	}
	return conn.Send("EXPIRE", key, expired) // 时效，单位秒
}

func Get(pool *redis.Pool, key string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return GetConn(conn,key)
}

func GetConn(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("GET", key)
}

func LPush(pool *redis.Pool, key string, value interface{}) error {
	conn := pool.Get()
	defer conn.Close()
	return conn.Send("LPUSH", key, value)
}

func LRange(pool *redis.Pool, key string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	return conn.Do("LRANGE", key, 0, -1)
}

func HIncrConn(conn redis.Conn, key, field interface{}) (int64, error) {
	if conn == nil {
		return 0, errors.New("invalid param")
	}
	return redis.Int64(conn.Do("HINCRBY", key, field, 1))
}

func HDescConn(conn redis.Conn, key, field interface{}) (int64, error) {
	if conn == nil {
		return 0, errors.New("invalid param")
	}
	return redis.Int64(conn.Do("HINCRBY", key, field, -1))
}

func HFieldExistsConn(conn redis.Conn, key, field interface{}) (bool, error) {
	if conn == nil {
		return false, errors.New("invalid param")
	}
	return redis.Bool(conn.Do("HEXISTS", key, field))
}

func HGetAllConn(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("HGETALL", key)
}

func HGetFieldsConn(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("HKEYS", key)
}
