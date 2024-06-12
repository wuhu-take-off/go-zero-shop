package db_init

func NewRedisConn(addr, password string) RedisConn {
	return &defaultRedisConn{RDB: connRedis(addr, password)}
}
