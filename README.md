# Blog API

A serverless Go application that serves blog posts through an API. It fetches Markdown files from a remote source (e.g.,
GitHub), converts them to HTML, and serves them via an AWS API Gateway endpoint.

## Features

- Serverless architecture using AWS Lambda and API Gateway
- Fetches blog posts from GitHub repositories
- Parses Markdown content with frontmatter support
- Converts Markdown to HTML
- Clean architecture with separation of concerns
- Configurable through environment variables

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

- **Domain Layer**: Core entities and interfaces
- **Use Case Layer**: Application business logic
- **Infrastructure Layer**: External systems integration (GitHub, Markdown parsing)

### Project Structure

```
├── .github/workflows/  # CI/CD workflows
├── posts/              # Blog post content in Markdown
├── resources/          # Configuration files
├── src/                # Application source code
│   ├── domain/         # Domain layer – entities and interfaces
│   ├── usecase/        # Use case layer – business logic
│   └── infrastructure/ # Infrastructure – external interfaces
├── main.go             # Application entry point
├── template.yaml       # AWS SAM template
└── Makefile            # Common build/test/run commands
```

## Tech Stack

- **Language**: Go 1.24
- **Framework**: AWS Serverless Application Model (SAM)
- **Runtime**: AWS Lambda (ARM64)
- **Key Dependencies**:
    - AWS Lambda Go SDK
    - Goldmark (Markdown parser)
    - Konfig (Configuration manager)
    - GitHub API client
    - slog (Structured logging)

## Getting Started

### Prerequisites

- Go 1.24+
- AWS SAM CLI
- AWS CLI configured with appropriate credentials
- GitHub token (if accessing private repositories)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/buyallmemes/blog-api.git
   cd blog-api
   ```

2. Install dependencies:
   ```bash
   make deps
   ```

3. Configure environment variables:
    - `GITHUB_TOKEN`: Your GitHub personal access token
    - `GITHUB_OWNER`: GitHub repository owner (default: "buyallmemes")
    - `GITHUB_REPO`: GitHub repository name (default: "blog-api")
    - `GITHUB_PATH`: Path to blog posts in the repository (default: "posts")

### Running Locally

```bash
make run
```

This will start a local API Gateway emulator. You can access the API at: http://localhost:3000/

### Testing

```bash
make go-test
```

This runs all unit tests with race detection and linting.

### Building

```bash
make build
```

This cleans, tests, and compiles the project using AWS SAM.

### Deployment

The project uses AWS SAM for deployment. You can deploy manually with:

```bash
sam deploy --guided
```

Or use the CI/CD pipeline defined in `.github/workflows/` for automated deployments.

## API Endpoints

- `GET /`: Returns all blog posts as JSON

## Development Guidelines

### Code Organization

- Follow **Clean Architecture** principles
- Place business logic in the appropriate layer
- Use interfaces for clear contracts and testability
- Avoid tight coupling between layers

### Error Handling

- Always wrap errors with context
- Log errors clearly and consistently
- Return meaningful HTTP status codes in handlers

### Testing

- Write unit tests for all new functionality
- Avoid relying on real external services; use mocks or fakes
- Keep tests fast, isolated, and deterministic

### Logging

- Use the structured logging package based on Go's `slog` library
- Log at appropriate levels (DEBUG, INFO, WARN, ERROR)
- Include relevant context in log entries using structured fields
- In development, set DEBUG level for more verbose logging
- In production, use INFO level for normal operations

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
