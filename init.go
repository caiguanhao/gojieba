package gojieba

import (
	"github.com/caiguanhao/gojieba/deps/cppjieba"
	"github.com/caiguanhao/gojieba/deps/limonp"
	"github.com/caiguanhao/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
