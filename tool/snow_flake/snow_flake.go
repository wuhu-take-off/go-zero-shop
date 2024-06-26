package snow_flake

import "github.com/bwmarrin/snowflake"

// GetSnowFlake 通过雪花算法生成唯一ID
func GetSnowFlake(id int64) snowflake.ID {
	node, _ := snowflake.NewNode(id)
	return node.Generate()
}
