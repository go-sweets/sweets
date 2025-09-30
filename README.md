# go-sweets

English | [简体中文](README.zh-CN.md)

> ⚠️ **Repository Reorganized**: This repository has been reorganized. All code has been moved to the main repository.

**New Repository**: [go-sweets/go-sweets](https://github.com/go-sweets/go-sweets)

---

## Overview

go-sweets is a Go framework for building cloud-native microservices with modern tools and best practices. It provides:

- **CLI Tool**: Project scaffolding tool
- **Service Template**: Production-ready microservice implementation using CloudWeGo framework
- **Shared Packages**: Reusable utilities for common tasks

## Quick Start

### 1. Install CLI Tool

```bash
git clone https://github.com/go-sweets/go-sweets.git
cd go-sweets/cli
go build -o mpctl main.go
```

### 2. Generate a New Service

```bash
./mpctl new <service-name>
```

### 3. Run Your Service

The generated service includes everything you need:

```bash
cd <service-name>
go mod tidy
make run
```

By default, the service listens on:

- **HTTP**: `http://localhost:8080` (Hertz)
- **RPC**: `localhost:9090` (Kitex)

### 4. Test Your Service

```bash
curl 'http://localhost:8080/v1/hello?id=1'
```

Expected response:

```json
{"id":"1","message":"Hello 1 !"}
```

## Architecture

### Service Template (sweets-layout)

Complete microservice implementation featuring:

- **CloudWeGo Hertz**: High-performance HTTP framework
- **CloudWeGo Kitex**: High-performance RPC framework with Protocol Buffers
- **Wire**: Compile-time dependency injection
- **GORM**: Database ORM with Goose migrations
- **Redis**: Caching and session management
- **DDD Architecture**: Domain-Driven Design with bounded contexts

See [sweets-layout/CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/sweets-layout/CLAUDE.md) for detailed documentation.

### Shared Packages (common/)

Independent utility packages:

- `conf/`: Configuration management
- `di/`: Dependency injection utilities
- `validator/`: Input validation
- `hash/`: Hashing utilities
- `lock/`: Distributed locking
- `migrate/`: Database migration tools
- `resp/`: Response formatting
- `contains/`: Container utilities
- `convert/`: Type conversion utilities
- `errcode/`: Error code management
- `str/`: String utilities
- `plugins/gorm/filter/`: GORM database filters

## Development

### Building the CLI

```bash
cd cli
go build -o mpctl main.go
```

### Working with Service Template

```bash
cd sweets-layout
make init     # Install dependencies and tools
make proto    # Generate protobuf code
make wire     # Run Wire dependency injection
make run      # Run the service
make test     # Run tests
make lint     # Run linter
```

### Using Shared Packages

Import packages in your code:

```go
import "github.com/go-sweets/go-sweets/common/<package-name>"
```

Or use local replace directives during development:

```go.mod
replace github.com/go-sweets/go-sweets/common/conf => ../common/conf
```

## Documentation

- **Main Repository**: [go-sweets/go-sweets](https://github.com/go-sweets/go-sweets)
- **Service Architecture**: [sweets-layout/CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/sweets-layout/CLAUDE.md)
- **Project Guide**: [CLAUDE.md](https://github.com/go-sweets/go-sweets/blob/main/CLAUDE.md)

## License

Apache License Version 2.0 - See [LICENSE](LICENSE) for details.

## Contributing

Contributions are welcome! Please visit the [main repository](https://github.com/go-sweets/go-sweets) to contribute.
