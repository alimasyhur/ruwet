# 🔥 Ruwet - Detect Messy Code Before It Ruins Your Day

<p align="center">
  <img src="https://github.com/user-attachments/assets/9c1590cf-317e-4305-bd2c-d6d797def15a" alt="Ruwet Logo" width="200"/>  
</p>

<p align="center">
  <a href="https://github.com/alimasyhur/ruwet/actions"><img src="https://img.shields.io/github/actions/workflow/status/alimasyhur/ruwet/go.yml?style=flat-square" alt="Build Status"></a>
  <a href="https://pkg.go.dev/github.com/alimasyhur/ruwet"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="Go Reference"></a>
  <a href="https://github.com/alimasyhur/ruwet/releases"><img src="https://img.shields.io/github/v/release/alimasyhur/ruwet?style=flat-square" alt="Release"></a>
  <a href="https://github.com/alimasyhur/ruwet/blob/main/LICENSE"><img src="https://img.shields.io/github/license/alimasyhur/ruwet?style=flat-square" alt="License"></a>
  <a href="https://github.com/alimasyhur/ruwet/stargazers"><img src="https://img.shields.io/github/stars/alimasyhur/ruwet?style=flat-square&logo=github" alt="Stars"></a>
</p>

<p align="center">
  <i>"Ruwet" means "messy" or "tangled" in Indonesian 🇮🇩</i>
</p>

---

## 🎯 What is Ruwet?

**Ruwet** is a high-performance CLI tool that analyzes source code complexity using **Cyclomatic Complexity**, a well-established metric introduced by **Thomas J. McCabe**.

It helps developers identify high-risk, hard-to-maintain functions before they turn into technical debt.

> Stop shipping spaghetti code. Find the tangled functions before they tangle your sanity Start shipping maintainable systems. 🍝

### ✨ Why Ruwet?

| Feature | Ruwet | Others |
|---------|-------|--------|
| ⚡ **Fast** - Written in Go | ✅ | ⚠️ |
| 🎨 **Beautiful CLI** with progress bar | ✅ | ❌ |
| 🔍 **Multi-language** (Go, JS, TS) | ✅ | ⚠️ |
| 🎯 **Top 10 messy functions** view | ✅ | ❌ |
| 📦 **Zero config** - just run it | ✅ | ❌ |
| 🌳 **Tree-sitter powered** (JS/TS) | ✅ | ❌ |

---

### 🧠 Scientific Foundation

Ruwet is grounded in **Cyclomatic Complexity**, which measures the number of **independent execution paths** within a function.

Higher complexity is strongly correlated with:

Increased defect probability
Reduced testability
Lower readability
Higher maintenance cost

In practical terms:

More branches → more paths → more potential failure points.

---

## 🚀 Quick Start

### Installation

```bash
# Go 1.18+
go install github.com/alimasyhur/ruwet@latest
```

Or download pre-built binaries from [Releases](https://github.com/alimasyhur/ruwet/releases).

### Usage

```bash
# Scan current directory
ruwet scan

# Scan specific path
ruwet scan /path/to/your/project
```

### 📸 Example Output

```
🔍 Scanning files...
[██████████████████████████████] 100% (1207/1207) Processing: main.go

✅ Scan complete
✔ Detected: [Go JavaScript/TypeScript]

Found 342 functions

🔥 Top messy functions:

1. handlePayment
   File: src/payment/processor.js
   Complexity: 23
   ⚠ Very high complexity
   💡 Break into smaller functions

2. validateUserInput
   File: internal/auth/handler.go
   Complexity: 18
   ⚠ Very high complexity
   💡 Break into smaller functions

3. renderDashboard
   File: components/Dashboard.tsx
   Complexity: 15
   ⚠ Very high complexity
   💡 Break into smaller functions
...
```

---

## 🧠 How It Works

Ruwet calculates **cyclomatic complexity** using these rules:

### Go Analyzer
- Base complexity: **1**
- `+1` for each: `if`, `for`, `range`, `case`
- `+1` for each: `&&`, `||` in binary expressions

### JavaScript/TypeScript Analyzer (Tree-sitter powered)
- Base complexity: **1**
- `+1` for each: `if`, `for`, `while`, `do`, `catch`, ternary `? :`
- `+1` for each: `switch_case`
- `+1` for each: `&&`, `||` in binary expressions

### Complexity Thresholds

| Complexity | Level | Action |
|------------|-------|--------|
| ≤ 5 | ✅ Low | Good to go! |
| 6-10 | ⚠️ Moderate | Consider refactoring |
| 11-15 | 🔥 High | Refactor soon |
| 16+ | 💥 Very High | Drop everything and refactor! |

---

## 📦 Supported Languages

- ✅ **Go** (`.go`) - uses `go/ast`
- ✅ **JavaScript** (`.js`, `.jsx`) - uses tree-sitter
- ✅ **TypeScript** (`.ts`, `.tsx`) - uses tree-sitter

---

## 🛠️ Advanced Usage

### Skip Directories

Ruwet automatically skips all inside `.gitignore` file.

### CI/CD Integration

```yaml
# .github/workflows/complexity-check.yml
name: Complexity Check
on: [push, pull_request]

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go install github.com/alimasyhur/ruwet@latest
      - run: ruwet scan .
```

---

## 🤝 Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📊 Stats

<p align="center">
  <img src="https://repobeats.axiom.co/api/embed/your-embed-id-here.svg" alt="Repo Stats">
</p>

---

## 🌟 Showcase

Built with ❤️ by developers who are tired of messy code.

If Ruwet helped you clean up your codebase, please consider giving it a ⭐!

---

## 📜 License

MIT License - see [LICENSE](LICENSE) for details.

---

## 🔗 Links

- 🏠 [GitHub Repository](https://github.com/alimasyhur/ruwet)
- 🐛 [Issue Tracker](https://github.com/alimasyhur/ruwet/issues)
- 💬 [Discussions](https://github.com/alimasyhur/ruwet/discussions)

---

<p align="center">
  <b>Stop writing ruwet code. Start writing clean code. 🚀</b>
</p>

<p align="center">
  <a href="https://github.com/alimasyhur/ruwet/stargazers">
    <img src="https://img.shields.io/github/stars/alimasyhur/ruwet?style=social" alt="Stars">
  </a>
</p>
