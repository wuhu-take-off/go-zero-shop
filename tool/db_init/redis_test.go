package db_init

import (
	"context"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	rdb := connRedis("8.130.145.246:6379", "112154ZhouM..")
	// 使用 Ping 命令检查连接是否正常
	// 确保在 main 函数中创建一个上下文
	ctx := context.Background()

	// 订阅 Redis 键空间通知
	pubsub := rdb.PSubscribe(ctx, "__keyevent@0__:expired")

	// 确保在退出时关闭订阅
	defer pubsub.Close()

	// 启动一个 Goroutine 处理消息
	go func() {
		for msg := range pubsub.Channel() {
			fmt.Printf("Key expired: %s\n", msg.Payload)
		}
	}()

}
