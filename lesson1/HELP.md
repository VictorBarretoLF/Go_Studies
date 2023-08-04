# 🚀 Go Unleashed: A Guide to Building & Running Go Applications

### ▶️ Quick Start: Run Your Go File

Execute the following command on bash to run a Go file:

```bash
go run main.go
```

### 🏗️ Building Your Go Application

Compile your Go code into an executable with:

```bash
go build main.go
```

### 🌍 Building for Different Operating Systems

You can build your Go application for different OS platforms using specific flags. Here's how:

- **Windows:**

  ```bash
  GOOS=windows go build main.go # for windows
  ```

- **Mac (64-bit):**

  ```bash
  GOOS=darwin GOARCH=amd64 go build main.go
  ```

- **Linux (64-bit):**

  ```bash
  GOOS=linux GOARCH=amd64 go build main.go
  ```

### ❌ Troubleshooting Errors in VS Code
