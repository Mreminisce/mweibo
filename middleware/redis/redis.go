package redis

// import (
// 	"encoding/json"
// 	"mweibo/conf"
// 	"time"

// 	"github.com/gomodule/redigo/redis"
// )

// var RedisConn *redis.Pool

// func init() {
// 	Setup()
// }

// func Setup() error {
// 	RedisConn = &redis.Pool{
// 		MaxIdle:     conf.Redisconfig.RedisMaxidle,
// 		MaxActive:   conf.Redisconfig.RedisMaxActive,
// 		IdleTimeout: conf.Redisconfig.RedisIdleTimeout,
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial("tcp", conf.Redisconfig.RedisHost, redis.DialDatabase(0))
// 			if err != nil {
// 				return nil, err
// 			}
// 			if conf.Redisconfig.RedisPassword != "" {
// 				if _, err := c.Do("AUTH", conf.Redisconfig.RedisPassword); err != nil {
// 					c.Close()
// 					return nil, err
// 				}
// 			}
// 			return c, err
// 		},
// 		TestOnBorrow: func(c redis.Conn, t time.Time) error {
// 			_, err := c.Do("PING")
// 			return err
// 		},
// 	}
// 	return nil
// }

// func Set(key string, data interface{}, time int) error {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	value, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = conn.Do("SET", key, value)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = conn.Do("EXPIRE", key, time)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func Exists(key string) bool {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	exists, err := redis.Bool(conn.Do("EXISTS", key))
// 	if err != nil {
// 		return false
// 	}
// 	return exists
// }

// func Get(key string) ([]byte, error) {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	reply, err := redis.Bytes(conn.Do("GET", key))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return reply, nil
// }

// func Delete(key string) (bool, error) {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	return redis.Bool(conn.Do("DEL", key))
// }

// func LikeDeletes(key string) error {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
// 	if err != nil {
// 		return err
// 	}
// 	for _, key := range keys {
// 		_, err = Delete(key)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func Lpush(key string, value interface{}, time int) error {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	_, err := conn.Do("LPUSH", key, value)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func Brpop(key string) (string, error) {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	reply, err := redis.String(conn.Do("RPOP", key))
// 	if err != nil {
// 		return reply, err
// 	}
// 	return reply, nil
// }
