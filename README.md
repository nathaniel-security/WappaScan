<h1 align="center">🔎 WappaScan: Wappalyzer CLI Tool</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="MIT License">
</p>

<p align="center">
  <b>Identify website technologies from your terminal, powered by Wappalyzer.<br>
  Output results in clean, developer-friendly JSON.</b>
</p>

---

## 🚀 Features

- 🕵️ Detects technologies used on any website
- 💡 Simple CLI interface
- 📦 Outputs results in JSON format
- ⚡ Fast and lightweight

---

## 📦 Installation

1. **Install Go:**  
   [Download Go](https://golang.org/dl/) and follow the instructions for your OS.

2. **Clone this repository:**
   ```sh
   git clone https://github.com/nathaniel-security/WappaScan.git
   cd WappaScan
   ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

---

## 🛠️ Usage

Analyze a website by running:

```sh
go run src/main.go <url>
```

Replace `<url>` with the website you want to scan.

### Example

```sh
go run src/main.go https://example.com
```

**Sample Output:**
```json
{
  "url": "https://example.com",
  "technologies": [
    "Nginx",
    "Bootstrap",
    "Google Analytics"
  ]
}
```

---
