# 🤝 Contributing to Ruwet

Thank you for considering contributing to Ruwet! It's people like you that make Ruwet such a great tool.

## 🎉 Ways to Contribute

- 🐛 Report bugs
- 💡 Suggest new features
- 📖 Improve documentation
- 🧑‍💻 Submit pull requests

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+ installed
- Node.js (for testing JS/TS analysis)
- Git

### Setup

```bash
# Fork the repo on GitHub, then:
git clone https://github.com/YOUR_USERNAME/ruwet.git
cd ruwet
go mod download
```

### Build & Test

```bash
# Build
go build -o ruwet .

# Run tests (once we have them 😉)
go test ./...

# Run locally
./ruwet scan .
```

---

## 📝 Pull Request Process

1. **Fork** the repository
2. **Create** a branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### PR Guidelines

- ✅ Keep changes focused and atomic
- ✅ Write clear commit messages
- ✅ Update documentation if needed
- ✅ Add tests for new features
- ✅ Make sure `go build` and `go vet` pass

---

## 🐛 Reporting Bugs

When reporting bugs, please include:

- Go version (`go version`)
- OS (macOS, Linux, Windows)
- Ruwet version (`ruwet --version` or commit hash)
- Steps to reproduce
- Expected vs actual behavior

Use the [bug report template](https://github.com/alimasyhur/ruwet/issues/new?template=bug_report.md) for structured reports.

---

## 💡 Suggesting Features

Have an idea? We'd love to hear it!

- 🗣️ Start a [Discussion](https://github.com/alimasyhur/ruwet/discussions)
- 🎫 Open a [Feature Request](https://github.com/alimasyhur/ruwet/issues/new?template=feature_request.md)

---

## 🧑‍💻 Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Keep functions small and focused
- Comment complex logic (ironic, right? 😄)

---

## 🏆 Contributors

Thanks to all the people who have contributed!

<!-- ALL-CONTRIBUTORS-LIST:START -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

Want to see your name here? Start contributing! 🚀

---

## 📧 Questions?

Feel free to reach out via:
- 💬 [GitHub Discussions](https://github.com/alimasyhur/ruwet/discussions)
- 🐛 [Issue Tracker](https://github.com/alimasyhur/ruwet/issues)

---

**Thank you for making Ruwet better! 🙏**
