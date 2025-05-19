# Developer Guidelines

## Project Overview

This project is a **serverless Go application** that serves blog posts through an API. It fetches Markdown files from a
remote source (e.g., GitHub), converts them to HTML, and serves them via an API Gateway endpoint. The application is
designed for scalability, clean structure, and testability.

## Tech Stack

- **Language**: Go 1.2x
- **Framework**: AWS Serverless Application Model (SAM)
- **Runtime**: AWS Lambda (ARM64)
- **Key Dependencies**:
    - AWS Lambda Go SDK
  - Markdown parser (e.g., Goldmark)
  - Configuration manager (e.g., Konfig)
  - GitHub or remote content fetcher

## Project Structure

- `.github/workflows/`: CI/CD workflows
- `posts/`: Blog post content in Markdown
- `resources/`: Configuration files (YAML, etc.)
- `src/`: Application source code
    - `domain/`: Domain layer – entities and interfaces
        - `blog/`: Domain models and contracts
    - `usecase/`: Use case layer – business logic
        - `blog/`: Blog service logic
    - `infrastructure/`: Infrastructure – external interfaces
        - `markdown/`: Markdown parser implementation
        - `repository/`: Content repository implementations
            - `github/`: GitHub-based repository
            - `local/`: Local file-based repository
    - `legacy/`: Legacy or deprecated code
- `main.go`: Application entry point
- `main_test.go`: Basic integration tests
- `Makefile`: Common build/test/run commands
- `template.yaml`: AWS SAM template

## Getting Started

### Prerequisites

- Go 1.2x
- AWS SAM CLI
- Valid credentials (e.g., GitHub token if using private content)

### Setup

```bash
git clone <your-repo>
cd <project>
make deps
```

## Development Workflow

### Run Locally

```bash
make run
```

Starts a local Lambda/API Gateway emulator.

### Run Tests

```bash
make go-test
```

Runs all unit tests with race detection and linting.

### Build

```bash
make build
```

Cleans, tests, and compiles the project using AWS SAM.

### Deploy

This project uses GitHub Actions for CI/CD. See `.github/workflows/` for pipeline definitions.

## Best Practices

### Code Organization

- Follow **Clean Architecture**: separate domain, use case, and infrastructure layers
- Place business logic in the `src` directory
- Use interfaces for clear contracts and testability
- Avoid tight coupling between layers

### Configuration

- Use environment variables for sensitive configuration
- Store default values in `resources/application.yaml`
- Never commit credentials — use SAM `Parameters` or AWS Secrets Manager

### Testing

- Write **unit tests** for all new functionality
- Avoid relying on real external services; use mocks or fakes
- Keep tests fast, isolated, and deterministic

### Error Handling

- Use proper error wrapping with context
- Log errors with appropriate detail
- Return meaningful HTTP status codes

## Architectural Principles

### Clean Architecture

This project follows Clean Architecture principles, organizing code into layers:

- **Domain Layer** (`src/domain`): Contains business entities and core business rules
- **Use Case Layer** (`src/usecase`): Application-specific business rules
- **Infrastructure Layer** (`src/infrastructure`): Frameworks, drivers, and external interfaces

Benefits of this architecture include:

- Independence from frameworks
- Testability
- Independence from UI
- Independence from database
- Independence from external agencies

### Clean Code Practices

The project adheres to Clean Code principles:

- Write self-explanatory code with meaningful names
- Follow the SOLID principles.
- Keep functions small and focused
- Minimize side effects
- Write comprehensive tests
- Refactor regularly
- Use consistent formatting and style