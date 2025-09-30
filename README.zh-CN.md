# go-sweets

[English](README.md) | 简体中文

> ⚠️ **仓库重组**: 本仓库已重组。所有代码已迁移至主仓库。

**新仓库**: [go-sweets/go-sweets](https://github.com/go-sweets/go-sweets)

---

## 概述

go-sweets 是一个用于构建云原生微服务的 Go 框架，采用现代工具和最佳实践。它提供：

- **CLI 工具**: 项目脚手架工具
- **服务模板**: 基于 CloudWeGo 框架的生产级微服务实现
- **共享包**: 常见任务的可重用工具

## 快速开始

### 1. 安装 CLI 工具

```bash
git clone https://github.com/go-sweets/go-sweets.git
cd go-sweets/cli
go build -o mpctl main.go
```

### 2. 生成新服务

```bash
./mpctl new <服务名称>
```

### 3. 运行服务

生成的服务包含您需要的一切：

```bash
cd <服务名称>
go mod tidy
make run
```

默认情况下，服务监听以下端口：

- **HTTP**: `http://localhost:8080` (Hertz)
- **RPC**: `localhost:9090` (Kitex)

### 4. 测试服务

```bash
curl 'http://localhost:8080/v1/hello?id=1'
```

预期响应：

```json
{"id":"1","message":"Hello 1 !"}
```

## 架构

### 服务模板 (sweets-layout)

完整的微服务实现，包含：

- **CloudWeGo Hertz**: 高性能 HTTP 框架
- **CloudWeGo Kitex**: 高性能 RPC 框架（支持 Protocol Buffers）
- **Wire**: 编译时依赖注入
- **GORM**: 数据库 ORM（支持 Goose 迁移）
- **Redis**: 缓存和会话管理
- **DDD 架构**: 领域驱动设计（有界上下文）

详细文档请参见 [sweets-layout/CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/sweets-layout/CLAUDE.md)。

### 共享包 (common/)

独立的工具包：

- `conf/`: 配置管理
- `di/`: 依赖注入工具
- `validator/`: 输入验证
- `hash/`: 哈希工具
- `lock/`: 分布式锁
- `migrate/`: 数据库迁移工具
- `resp/`: 响应格式化
- `contains/`: 容器工具
- `convert/`: 类型转换工具
- `errcode/`: 错误码管理
- `str/`: 字符串工具
- `plugins/gorm/filter/`: GORM 数据库过滤器

## 开发

### 构建 CLI 工具

```bash
cd cli
go build -o mpctl main.go
```

### 使用服务模板

```bash
cd sweets-layout
make init     # 安装依赖和工具
make proto    # 生成 protobuf 代码
make wire     # 运行 Wire 依赖注入
make run      # 运行服务
make test     # 运行测试
make lint     # 运行代码检查
```

### 使用共享包

在代码中导入包：

```go
import "github.com/go-sweets/go-sweets/common/<包名>"
```

或在开发时使用本地替换指令：

```go.mod
replace github.com/go-sweets/go-sweets/common/conf => ../common/conf
```

## 文档

- **主仓库**: [go-sweets/go-sweets](https://github.com/go-sweets/go-sweets)
- **服务架构**: [sweets-layout/CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/sweets-layout/CLAUDE.md)
- **项目指南**: [CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/CLAUDE.md)

## 许可证

Apache License Version 2.0 - 详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎贡献！请访问[主仓库](https://github.com/go-sweets/go-sweets)参与贡献。
