package blockpool

import (
	"tinychain/common"
)

type Config struct {
	MaxBlockSize uint64 // Maximum number of blocks
}

func newConfig(config *common.Config) *Config {
	return &Config{
		MaxBlockSize: uint64(config.GetInt64("blockpool.max_block_size")),
	}
}