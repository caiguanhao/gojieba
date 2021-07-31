package gojieba

import (
	"github.com/caiguanhao/gojieba/deps/cppjieba"
	"github.com/caiguanhao/gojieba/deps/limonp"
)

func init() {
	limonp.Init()
	cppjieba.Init()
}
