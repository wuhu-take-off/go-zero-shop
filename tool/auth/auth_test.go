package auth

import (
	"fmt"
	"github.com/willf/bloom"
	"testing"
)

func TestCreateAuth(t *testing.T) {
	// 创建一个 Cuckoo 过滤器，设置桶的数量为10000
	filter := bloom.NewWithEstimates(1000000, 0.01)

	// 将元素添加到过滤器中
	filter.Add([]byte("hello"))

	// 检查元素是否存在于过滤器中
	exists := filter.Test([]byte("hello"))
	fmt.Println("Exists:", exists) // 输出: true
}
