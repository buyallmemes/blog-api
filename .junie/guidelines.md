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

- Always wrap errors with context using `fmt.Errorf` or `errors.Join`
- Log errors clearly and consistently
- Return meaningful HTTP status codes in handlers

## Architectural Principles

### Clean Architecture

Organize your code into layers with clear responsibilities:

- **Domain** (`src/domain`) – Core entities and interfaces
- **Use Case** (`src/usecase`) – Application business logic
- **Infrastructure** (`src/infrastructure`) – External systems (e.g., GitHub, Markdown)

Benefits:

- Framework agnostic
- Testable components
- Easy to evolve
- Clear boundaries

### Clean Code Practices

- Use descriptive, intention-revealing names
- Keep functions small and focused
- Minimize side effects
- Follow the **SOLID principles**
- Keep formatting consistent
- Refactor regularly
- Avoid unnecessary abstractions

## References

- *Clean Code* – Robert C. Martin
- *Clean Architecture* – Robert C. Martin
- *Refactoring* – Martin Fowler
- *Tidy First?* – Kent Beck
- *Extreme Programming Explained* – Kent Beck  