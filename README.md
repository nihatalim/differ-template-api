# Differ Template API

A Go-based API template project designed following Hexagonal Architecture principles.

## Project Structure

```
differ-template-api/
├── api/                    # API documentation
│   └── swagger/           # Swagger/OpenAPI specs
├── cmd/                   # Application entry points
│   └── api/              # API server
├── internal/              # Private application code
│   ├── health/           # Health check feature
│   └── common/           # Shared utilities
├── pkg/                   # Public shared code
│   ├── config/           # Configuration
│   └── middleware/       # HTTP middlewares
```

## Technologies

- Go 1.21
- [Fiber](https://github.com/gofiber/fiber) - Fast HTTP framework
- Swagger/OpenAPI - API documentation

## Getting Started

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run cmd/api/main.go
```

3. Test the API:
```bash
curl http://localhost:8080/
```

## Features

- Hexagonal Architecture
- Middleware support
- Swagger documentation
- Health check endpoint
- Configurable port
- Logging middleware

## Project Layout

- `api/`: OpenAPI/Swagger specs, JSON schema files, protocol definition files
- `cmd/`: Main applications for this project
- `internal/`: Private application and library code
- `pkg/`: Library code that's safe to use by external applications

## Development

### Prerequisites

- Go 1.21 or higher
- Git

### Code Style

This project follows the standard Go project layout and coding conventions:
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details