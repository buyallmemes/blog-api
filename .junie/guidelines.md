# Blog API Developer Guidelines

## Project Overview

This is a serverless Go application that serves blog posts via AWS Lambda. The application fetches Markdown files from
GitHub, converts them to HTML, and serves them through an API Gateway endpoint.

## Tech Stack

- **Language**: Go 1.24
- **Framework**: AWS Serverless Application Model (SAM)
- **Runtime**: AWS Lambda (ARM64)
- **Dependencies**:
    - AWS Lambda Go SDK
    - GitHub API client
    - Goldmark (Markdown processing)
    - Konfig (configuration management)

## Project Structure

```
blog-api/
├── .github/workflows/    # CI/CD workflows
├── posts/                # Blog post content in Markdown
├── resources/            # Configuration files
├── src/                  # Application source code
│   └── blog/             # Blog-related functionality
│       ├── fetcher/      # Post fetching logic
│       └── md/           # Markdown processing
├── main.go               # Application entry point
├── main_test.go          # Main tests
├── Makefile              # Build and run commands
└── template.yaml         # AWS SAM template
```

## Getting Started

### Prerequisites

- Go 1.24 or later
- AWS SAM CLI
- GitHub token (for fetching posts)

### Setup

1. Clone the repository
2. Run `make deps` to install dependencies

## Development Workflow

### Running Locally

```bash
make run
```

This builds the application and starts a local API Gateway emulator.

### Testing

```bash
make go-test
```

This runs all tests with race detection and Go's vet tool.

### Building

```bash
make build
```

This cleans previous builds, runs tests, and builds the application using SAM.

### Deployment

The project uses GitHub Actions for CI/CD. See the workflows in `.github/workflows/`.

## Best Practices

### Code Organization

- Keep business logic in the `src` directory
- Maintain separation of concerns (fetching, parsing, etc.)
- Use interfaces for dependency injection and testability

### Testing

- Write unit tests for all new functionality
- Ensure tests are independent and don't rely on external services
- Run tests before submitting changes

### Configuration

- Use environment variables for configuration
- Store configuration in `resources/application.yaml`
- Never commit sensitive information (use parameters in SAM template)

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
