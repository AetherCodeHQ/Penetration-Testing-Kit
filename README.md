# 🛡️ Penetration Testing Kit

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-00ADD8?style=for-the-badge)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-00ADD8?style=for-the-badge)
![Lint](https://img.shields.io/badge/GolangCI--Lint-Passing-00ADD8?style=for-the-badge)
![Dependabot](https://img.shields.io/badge/Dependabot-Enabled-025E9C?style=for-the-badge&logo=dependabot&logoColor=white)

> Penetration-Testing-Kit - Open source tool by AetherCodeHQ

`cli` `golang` `pentesting` `security`

## What is it?

**Penetration Testing Kit** is a security tool designed for developers who need fast, reliable, and offline-capable tools. Built with Go for maximum performance and minimal resource usage.

## Why should you care?

- 🚀 **Fast** — Compiled Go binary, no runtime dependencies
- 🔒 **Secure** — CodeQL security analysis + Dependabot
- 🌐 **Offline-first** — Works without internet connection
- 📦 **Lightweight** — Single binary, minimal footprint
- 🛠️ **Developer-friendly** — Clean CLI with helpful documentation

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
git clone https://github.com/AetherCodeHQ/Penetration-Testing-Kit.git
cd Penetration-Testing-Kit
go build -o Penetration-Testing-Kit .
```

### Run
```bash
./Penetration-Testing-Kit --help
```

## Usage

```bash
# Basic usage
./Penetration-Testing-Kit --path ./target

# With options
./Penetration-Testing-Kit --path ./target --format json --output report.json

# Verbose mode
./Penetration-Testing-Kit --path ./target --verbose
```

## Features

- ✅ High-performance Go implementation
- ✅ Cross-platform support (Windows, Linux, macOS)
- ✅ JSON export for CI/CD integration
- ✅ Colored terminal output
- ✅ Configurable via YAML/JSON
- ✅ Comprehensive documentation

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Target directory | `.` |
| `--format` | Output format (json, text, badge) | `text` |
| `--output` | Output filename | `stdout` |
| `--verbose` | Enable verbose output | `false` |
| `--config` | Config file path | - |

## Examples

### Basic scan
```bash
./Penetration-Testing-Kit --path ./my-project
```

### JSON report
```bash
./Penetration-Testing-Kit --path ./my-project --format json --output report.json
```

### CI/CD integration
```yaml
# .github/workflows/scan.yml
- name: Run Penetration Testing Kit
  run: ./Penetration-Testing-Kit --path . --format json --output report.json
```

## Development

```bash
# Clone the repo
git clone https://github.com/AetherCodeHQ/Penetration-Testing-Kit.git
cd Penetration-Testing-Kit

# Build
go build -o Penetration-Testing-Kit .

# Run tests
go test ./...

# Lint
golangci-lint run
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

If you discover a security vulnerability, please report it responsibly. See [SECURITY.md](SECURITY.md) for details.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  Built with ❤️ by <a href="https://github.com/AetherCodeHQ">AetherCode</a> • <a href="https://github.com/AetherCode-Core">AetherCode-Core</a>
</p>


---

## What's New in v1.1.0

- Professional documentation with badges
- CI/CD pipeline with GitHub Actions
- Security analysis with CodeQL
- Dependency management with Dependabot
- Code quality with GolangCI-Lint
- Issue and PR templates
- Contributing guidelines

## Categories

| Category | Description |
|----------|-------------|
| Security & Penetration Testing | This project is part of the AetherCode ecosystem |

## Related Projects

Part of [AetherCode Core](https://github.com/AetherCode-Core) ecosystem.
