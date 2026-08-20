# Penetration Testing Kit

![CI](https://github.com/Qyroxen/Penetration-Testing-Kit/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/Penetration-Testing-Kit?style=social)

> A powerful CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Penetration-Testing-Kit?style=social)](https://github.com/Qyroxen/Penetration-Testing-Kit/stargazers)

## What is it?

Penetration Testing Kit is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Penetration-Testing-Kit.git
cd Penetration-Testing-Kit
go build -o penetrationtestingkit .

# Run
./penetrationtestingkit --help
```

## CLI Usage

```bash
# Basic usage
./penetrationtestingkit

# With flags
./penetrationtestingkit --verbose --output json

# Get help
./penetrationtestingkit --help
```

## Examples

```bash
# Example 1
./penetrationtestingkit example1

# Example 2
./penetrationtestingkit example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o penetrationtestingkit .

# Lint
go vet ./...
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Penetration-Testing-Kit/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Penetration-Testing-Kit?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Penetration-Testing-Kit/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Penetration-Testing-Kit?style=social" alt="Fork this repo">
  </a>
</p>
