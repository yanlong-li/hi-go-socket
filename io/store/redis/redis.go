package redis

import (
	"fmt"
	"garyburd/redigo/redis"
)

var Conn redis.Conn

func init() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	Conn = conn

	username, err := redis.String(Conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func Set(name, value interface{}) (interface{}, error) {
	return Conn.Do("SET", name, value)
}
func Get(name interface{}) (interface{}, error) {
	return Conn.Do("GET", name)
}
func Del(name interface{}) (interface{}, error) {
	return Conn.Do("DEL", name)
}
