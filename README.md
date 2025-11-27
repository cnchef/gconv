[English](README_US.md) | 中文

# gconv – Go 通用类型转换工具包

[![Go Version](https://img.shields.io/badge/Go-%3E=1.20-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Coverage](https://img.shields.io/badge/Coverage-98.2%25-brightgreen)](gconv_test.go)

`gconv` 是一个零依赖、轻量级的 Go 工具包，为处理 **动态类型** 和 **接口返回数据类型不固定** 的场景提供一套类似 Python 的强力转换能力。

通过 `Cast[T]` 泛型函数，你可以将任意类型的值转换为指定的类型，极大提升处理外部 API、JSON、Map、动态字段时的开发体验。

## ✨ 功能特点

- ✔ 支持 Go 泛型 Cast[T]
- ✔ 任意类型 → string / int / float64 / bool
- ✔ 支持 map[string]any / []any
- ✔ 自动处理 json.Number
- ✔ 自动处理数值字符串
- ✔ 支持 struct：使用 JSON 进行自动转换
- ✔ 零第三方依赖，彻底轻量化
- ✔ 对不支持的类型自动返回零值，避免 panic

---

## 📦 安装

方式一：使用 go get 安装

```bash
go get github.com/cnchef/gconv
```

方式二：将 `gconv` 文件夹复制到你的项目中

或自定义 module 名：

```bash
go mod init gconv
```

---

## 🛠 使用示例

```go
package main

import (
    "fmt"
    "gconv"
)

func main() {
    var v any = "123"

    fmt.Println(gconv.Cast[int](v))        // 123
    fmt.Println(gconv.Cast[float64](v))    // 123.0
    fmt.Println(gconv.Cast[string](v))     // "123"
    fmt.Println(gconv.Cast[bool]("true"))  // true

    m := map[string]any{"a": 1}
    fmt.Println(gconv.Cast[map[string]any](m)) // map[a:1]
}
```

**更多示例：** 查看 [examples/](examples/) 目录，包含：
- `basic.go` - 基础使用示例
- `advanced.go` - 高级用法（struct 转换、API 响应处理、批量转换等）

---

## 🧩 Cast[T] 泛型转换

`Cast[T]` 是一个万能转换器，用法简单：

```go
value := Cast[T](v)
```

示例：

```go
age := gconv.Cast[int]("20")         // 20
price := gconv.Cast[float64]("99.9") // 99.9
flag := gconv.Cast[bool]("true")     // true
name := gconv.Cast[string](123)      // "123"
```

---

## 🔧 提供的转换函数

### ToString(v any) string

精准转换所有基础类型。

### ToInt(v any) int

自动处理 `"123"`、`123.0`、`json.Number`。

### ToFloat(v any) float64

自动支持字符串与数字混合格式。

### ToBool(v any) bool

支持 `"true"` `"1"` `"false"` `"0"`。

### ToMap(v any) map[string]any

自动强转不正确类型时返回空 map。

### ToSlice(v any) []any

自动强转不正确类型时返回空 slice。

---

## ⚙️ 文件结构

```
gconv/
 ├── gconv.go              # 核心转换函数
 ├── gconv_test.go         # 完整单元测试（覆盖率 98.2%）
 ├── go.mod                # Go 模块定义
 ├── README.md             # 中文文档
 ├── README_US.md          # 英文文档
 ├── LICENSE               # MIT 开源许可
 ├── CHANGELOG.md          # 版本更新记录
 ├── .gitignore            # Git 忽略配置
 ├── .github/
 │   └── workflows/
 │       └── test.yml      # GitHub Actions CI 配置
 └── examples/             # 使用示例
     ├── basic.go          # 基础使用示例
     ├── advanced.go       # 高级使用示例（struct 转换等）
     └── go.mod            # 示例模块配置
```

---

## 🧪 测试

运行单元测试：

```bash
go test -v -cover
```

运行性能测试：

```bash
go test -bench=. -benchmem
```

当前测试覆盖率：**98.2%**

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

在提交 PR 前，请确保：
- 所有测试通过 `go test ./...`
- 代码格式化 `go fmt ./...`
- 添加必要的测试用例

---

## 📄 License

[MIT License](LICENSE)

---

## 🔗 相关链接

- [GitHub 仓库](https://github.com/cnchef/gconv)
- [问题反馈](https://github.com/cnchef/gconv/issues)
- [变更日志](CHANGELOG.md)
