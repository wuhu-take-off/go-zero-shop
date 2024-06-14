package db_init

import (
	"TongChi_shop/tool/auth"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/threading"
	"time"
)

type (
	RedisConn interface {
		AddEmailAuth(ctx context.Context, email string) error
		CheckEmailAuth(ctx context.Context, email, auth string) bool
	}
	defaultRedisConn struct {
		RDB *redis.Client
	}
)

func connRedis(addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功...")
	return client
}

// AddAuth 记录验证码到redis中
const emailAuthHead = "emailAuth:"
const emailAuthExpiration = time.Minute * 5
const reEmailAuthExpiration = time.Minute * 4

func (m *defaultRedisConn) AddEmailAuth(ctx context.Context, email string) error {
	key := emailAuthHead + email

	//该邮箱从未发送过验证码，或者距离上次发送验证码的时间超过了1分钟就重新发送验证码
	ttl := m.RDB.TTL(ctx, key)
	exp, _ := ttl.Result()
	if ttl.Err() == redis.Nil || exp < reEmailAuthExpiration {
		code := auth.CreateAuth(auth.EmailAuthLen)
		m.RDB.Set(ctx, key, code, emailAuthExpiration)
		//异步的go协程发送邮件
		threading.GoSafe(func() {
			auth.Send(auth.SendEmail, auth.Password, email, auth.Subject, code)
		})
		return nil

	}
	return errors.New("验证码重复发送")
}

func (m *defaultRedisConn) CheckEmailAuth(ctx context.Context, email, auth string) (ok bool) {
	key := emailAuthHead + email
	result, _ := m.RDB.Get(ctx, key).Result()
	ok = result == auth
	if ok {
		m.RDB.Del(ctx, key)
	}
	return
}
