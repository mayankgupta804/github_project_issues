package redisclient

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	Pool *redis.Pool
}

func (rc *RedisClient) Init(redisHost string) *RedisClient {
	rc.Pool = newPool(redisHost)
	rc.cleanupHook()
	return rc
}

func (rc *RedisClient) cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		rc.Pool.Close()
		os.Exit(0)
	}()
}

func newPool(server string) *redis.Pool {

	u, err := url.Parse(server)

	if err != nil {
		log.Fatalf("Error encountered: %v", err)
	}

	if u.User == nil {
		log.Fatalf("Error encountered: %v", err)
	}

	pw, ok := u.User.Password()
	if !ok {
		log.Fatalf("Error encountered: %v", err)
	}

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", u.Host)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", pw)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (rc *RedisClient) Ping() error {

	conn := rc.Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func (rc *RedisClient) Get(key string) ([]byte, error) {

	conn := rc.Pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func (rc *RedisClient) Set(key string, value []byte) error {

	conn := rc.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func (rc *RedisClient) Exists(key string) (bool, error) {

	conn := rc.Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func (rc *RedisClient) Delete(key string) error {

	conn := rc.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}
