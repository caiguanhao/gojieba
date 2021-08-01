# gojieba

Fork of [github.com/yanyiwu/gojieba](https://github.com/yanyiwu/gojieba).

- Embed dictionaries into executable.
- gojieba.NewJieba(userDictData...) accepts only user dictionary content.

```go
package main

import (
        "fmt"
        "strings"

        "github.com/caiguanhao/gojieba"
)

func main() {
        text := strings.TrimSpace(`
2003年，一群希望证明电动车比燃油车更好、更快、并拥有更多驾驶乐趣的工程师创立了 Tesla。
今天，Tesla 不仅制造纯电动汽车，还可以生产能够无限扩容清洁能源收集及储存产品。
Tesla 相信，让世界越早摆脱对化石燃料的依赖，向零排放迈进，人类的前景就会更美好。
`)
        jb := gojieba.NewJieba()
        defer jb.Free()
        fmt.Printf("%#v\n", jb.Tag(text))
}
// []string{"2003/m", "年/m", "，/x", "一群/m", "希望/v", "证明/n", "电动车/n",
// "比/p", "燃油/n", "车/zg", "更好/d", "、/x", "更/d", "快/a", "、/x", "并/c",
// "拥有/v", "更/d", "多/m", "驾驶/v", "乐趣/a", "的/uj", "工程师/n", "创立/v",
// "了/ul", " /x", "Tesla/eng", "。/x", "\n/x", "今天/t", "，/x", "Tesla/eng", "
// /x", "不仅/c", "制造/v", "纯/a", "电动汽车/n", "，/x", "还/d", "可以/c",
// "生产/vn", "能够/v", "无限/v", "扩容/v", "清洁/a", "能源/n", "收集/v", "及/c",
// "储存/n", "产品/n", "。/x", "\n/x", "Tesla/eng", " /x", "相信/v", "，/x",
// "让/v", "世界/n", "越早/d", "摆脱/v", "对/p", "化石/n", "燃料/n", "的/uj",
// "依赖/v", "，/x", "向/p", "零排放/l", "迈进/v", "，/x", "人类/n", "的/uj",
// "前景/n", "就/d", "会/v", "更/d", "美好/a", "。/x"}
```
