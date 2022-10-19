# ID Generator

## 安装

```shell
go get github.com/istepheny/idgenerator@latest
```

## 生成 int53 ID

```go
package main

import (
	"fmt"
	"github.com/istepheny/idgenerator"
)

func main() {
	// 默认使用 int53 ID 生成器，用于要和前端页面交互的场景
	id := idgenerator.ID()
	fmt.Println(id)
}
```

## 生成 int64 ID

```go
package main

import (
	"fmt"
	"github.com/istepheny/idgenerator"
	"github.com/istepheny/idgenerator/drivers/int64IDGenerator"
)

func main() {
	id := idgenerator.Driver(int64IDGenerator.DriverName).ID()
	fmt.Println(id)
}
```

## 自定义 ID 生成器

```go
// drivers/fakeIDGenerator/fakeIDGenerator.go
package fakeIDGenerator

import (
	"github.com/istepheny/idgenerator/contract"
	"github.com/istepheny/idgenerator/drivers"
)

const DriverName = "FakeIDGenerator"

func init() {
	drivers.Register(DriverName, func() contract.IDGenerator { return NewFakeIDGenerator() })
}

type FakeIDGenerator struct{}

func NewFakeIDGenerator() *FakeIDGenerator {
	return &FakeIDGenerator{}
}

func (FakeIDGenerator) ID() int64 {
	// 在这里写 ID 生成逻辑
	return 0
}
```

## 使用自定义 ID 生成器

```go
package main

import (
	"fmt"
	"github.com/istepheny/idgenerator"
	"github.com/istepheny/idgenerator/drivers/fakeIDGenerator"
)

func main() {
	id := idgenerator.Driver(fakeIDGenerator.DriverName).ID()
	fmt.Println(id)
}

```

## 参考

- https://github.com/bwmarrin/snowflake
- https://github.com/sony/sonyflake
- https://xr1s.me/2021/05/02/json-etc
- https://juejin.cn/post/6844903981886472206
