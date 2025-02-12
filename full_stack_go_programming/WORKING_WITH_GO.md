# Go Modules: Dependency Management  

Ensure that other developers can build your code using a similar version.  

Go uses SIV (Semantic Import Versioning), based on [Semantic Versioning](https://semver.org/).  

### Highlights:
- Composed of three numbers, e.g., `v1.12.4`
- Format: **major**, **minor**, **patch**
  - Major bumps indicate breaking changes, requiring a manual upgrade.
  - Other bumps can be automated by tooling if needed.

### Examples:
- Installing `v1.12.4`: Other developers may build with `v1.13.0` if it releases.
- If releasing a new version from `v1.13.6` with breaking changes, the next version will probably be `v2.0.0`.
- If `v2.0.0` releases, tools **will NOT** upgrade to it automatically.

### Special Case:
- **v0** allows any breaking changes since it's a special version.
- Running `go get github.com/joncalhoun/somelib` **will NOT** get the most recent version.
- Use `go get github.com/joncalhoun/somelib/v4` or similar.  
- Tooling isn't always perfect but will hopefully improve over time. 

## 2. Working Outside of `GOPATH`  

Previously, all Go code had to reside within a single directory on your computer, known as the `GOPATH`.  

With Go Modules, we can now run code from anywhere as long as a module is initialized.

## 3. Setting Up Our Module

```bash
go mod init github.com/joncalhoun/lenslocked
# Use your own GitHub handle!
```

This command creates a `go.mod` file.  

From this point forward, there’s no need to think too much about modules. When using `go get` to fetch a library, the tooling will automatically update the `go.mod` file.

If necessary, I'll highlight cases where we need to consider modules.

**Gotcha:** In some editors, you must open the code from the directory containing the `go.mod` file for Go tooling to work correctly. Example:  

```bash
cd <directory with go.mod>
code .
```