package snow_flake

import (
	"fmt"
	"testing"
)

func TestGetSnowFlake(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetSnowFlake(int64(i)))
	}
}
