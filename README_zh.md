Period
============


[![Author](http://img.shields.io/badge/author-@maogou-blue.svg?style=flat-square)](https://github.com/maogou)
[![Latest Version](https://img.shields.io/github/release/maogou/period.svg?style=flat-square)](https://github.com/maogou/period/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/maogou/period.svg)](https://pkg.go.dev/github.com/maogou/period)
[![License](https://img.shields.io/github/license/maogou/period.svg)](https://github.com/maogou/period/blob/master/LICENSE)

项目介绍
-------

Period 是一个使用 Go 语言编写并实现 [thephpleague/period](https://github.com/thephpleague/period) 的项目，主要用于处理和操作时间段。它提供了一系列的方法，如并集、交集、差集等，以便于用户对时间段进行各种操作。

> ⚠️ **警告**：当前,Go(1.21.x) 语言还不支持枚举类型。如果 Go 语言在将来版本中支持枚举类型，我们将重构 `bounds.go` 文件中的方法。

[英文介绍](https://github.com/maogou/period/blob/main/README.md)

主要功能
-------

- **创建时间段**：用户可以创建一个时间段，包括开始时间和结束时间。
- **时间段并集**：用户可以计算两个或多个时间段的并集。
- **时间段交集**：用户可以计算两个或多个时间段的交集。
- **时间段差集**：用户可以计算两个时间段的差集。


如何使用
-------

首先，你需要在你的 Go 项目中引入 Period 项目。你可以在你的 `go.mod` 文件中添加以下依赖：

```go
require (
    github.com/maogou/period v1.0.0
)
```

然后，你可以在你的代码中引入 Period 项目，并使用它提供的方法。以下是一个简单的示例：

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/maogou/period"
)

func main() {
	
	ts1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	te1 := time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local)
	
	p1 := period.NewDefaultPeriod(ts1, te1)
	
	ts2 := time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)
	te2 := time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local)
	p2 := period.NewDefaultPeriod(ts2, te2)
	
	unions := p1.Union(p2)
	
	for index, item := range unions.GetInterval() {
		fmt.Println(index, ":", item.Format(time.DateTime))
	}

}

```

方法
-------

以下是 `Period` 结构体的主要方法：

- `GetTimestampInterval()`: 返回时间段的时间戳间隔。
- `Subtract(Period)`: 从当前时间段中减去另一个时间段，返回一个新的时间段。
- `Merge(Period)`: 合并当前时间段和另一个时间段，返回一个新的时间段。
- `Intersect(Period)`: 返回当前时间段和另一个时间段的交集。
- `Overlaps(Period)`: 判断当前时间段是否与另一个时间段重叠。
- `Abuts(Period)`: 判断当前时间段是否与另一个时间段相邻。
- `Contains(Period)`: 判断当前时间段是否包含另一个时间段。
- `Equals(Period)`: 判断当前时间段是否等于另一个时间段。
- `Gap(Period)`: 返回当前时间段和另一个时间段之间的间隙。
- `Diff(Period)`: 返回当前时间段和另一个时间段的差异。
- `Union(Period...)`: 获取当时时间段集合的并集。
- `IsZero()`: 判断当前时间段是否为零。

以下是 `Sequence` 结构体的主要方法：

- `NewSequence(Period...)`: 创建一个新的时间段序列。
- `totalTimeDuration()`: 返回时间段序列的总时间长度。
- `GetTotalTimestampInterval()`: 返回时间段序列的总时间戳间隔。
- `Sort(func(Period, Period) bool)`: 根据给定的比较函数对时间段序列进行排序。
- `Contains(Period...)`: 判断时间段序列是否包含给定的时间段。
- `Subtract(Sequence)`: 从当前时间段序列中减去另一个时间段序列，返回一个新的时间段序列。
- `Equals(Sequence)`: 判断当前时间段序列是否等于另一个时间段序列。
- `Unions()`: 返回时间段序列的并集。
- `Gaps()`: 返回时间段序列的间隙。
- `IsEmpty()`: 判断时间段序列是否为空。
- `IndexOf(Period)`: 返回给定时间段在时间段序列中的索引。
- `Count()`: 返回时间段序列的数量。
- `Get(int)`: 返回时间段序列中给定索引的时间段。
- `Set(int, Period)`: 在时间段序列的给定索引处设置时间段。
- `Push(Period...)`: 在时间段序列的末尾添加时间段。
- `Intersections()`: 返回时间段序列的交集。
- `Remove(int)`: 移除时间段序列中给定索引的时间段。
- `Filter(func(Period) bool)`: 根据给定的过滤函数过滤时间段序列。
- `Map(func(Period) Period)`: 根据给定的映射函数映射时间段序列。
- `Every(func(Period, int) bool)`: 判断时间段序列是否每个元素都满足给定的条件。
- `Some(func(Period, int) bool)`: 判断时间段序列是否有元素满足给定的条件。
- `Clear()`: 清空时间段序列。

测试
-------

Period 项目包含一系列的单元测试，以确保代码的质量和功能的正确性。你可以通过运行以下命令来执行这些测试：

```bash
go test ./... -v -count=1
```

代码覆盖率(100%)
-------

我们使用 Go 的内置测试工具来跟踪代码覆盖率。你可以通过以下步骤生成和查看覆盖率报告：

```bash
 go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
``` 

贡献
-------

我们欢迎任何形式的贡献，包括但不限于提交问题、提供反馈、提出新的功能建议、改进代码等。如果你有任何问题或建议，欢迎在 GitHub 上提交 issue 或 pull request。

许可证
-------

Period 项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。