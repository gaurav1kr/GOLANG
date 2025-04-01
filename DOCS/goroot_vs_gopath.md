
# GOROOT vs GOPATH in Go

## GOROOT:
- **Definition**: GOROOT specifies the directory where your Go installation is located. This directory contains the Go standard library, compiler, and other Go tools.
- **Purpose**: It tells the Go toolchain where to find the Go installation, including the Go binaries and standard packages.
- **Set by Default**: In most cases, `GOROOT` is automatically set by the Go installation process, so you don’t typically need to modify or set this variable manually.

### Example:
- Let’s say you install Go at `/usr/local/go`, and you check the value of `GOROOT`:
  ```bash
  $ go env GOROOT
  /usr/local/go
  ```
  
- If you inspect `/usr/local/go/src`, you will see the Go standard library source code (e.g., `fmt`, `net`, etc.). These are built-in libraries that come with the Go installation.
- You don’t place your own code inside `GOROOT`; it's solely for Go's own tooling and libraries.

### Structure of GOROOT:
- `/bin`: The Go compiler and other tools.
- `/pkg`: Precompiled standard libraries.
- `/src`: The source code of the Go standard library.

## GOPATH:
- **Definition**: GOPATH defines the root directory for your Go workspace. This is where your Go project files (source code, binaries, and dependencies) are stored.
- **Purpose**: It tells Go where to look for and place your own Go projects, downloaded dependencies, and compiled binaries.
- **Customizable**: You can set `GOPATH` to any directory on your system, depending on where you want to store your Go projects and dependencies.

### Example:
- Let’s assume you set `GOPATH` to `/home/gaurav/go_projects`, where you store your own Go projects. You can set the `GOPATH` environment variable like this:
  ```bash
  export GOPATH=/home/gaurav/go_projects
  ```
  
  Then, your directory structure inside `GOPATH` would look like this:
  ```plaintext
  /home/gaurav/go_projects/
      bin/      # Compiled binaries (e.g., from 'go install')
      pkg/      # Compiled packages (e.g., installed dependencies)
      src/      # Your project source files
  ```

- If you create a new Go project, say `myapp`, your project structure might look like this:
  ```plaintext
  /home/gaurav/go_projects/
      src/
          github.com/gaurav/myapp/
              main.go
  ```

Now, when you build or run your project:
```bash
$ cd $GOPATH/src/github.com/gaurav/myapp
$ go build
$ ./myapp  # This would run your program
```

### Structure of GOPATH:
- `/src`: This is where your Go source code lives (organized by package).
- `/pkg`: Compiled package objects (this is where Go puts your dependency binaries).
- `/bin`: This is where Go puts the compiled binaries from your Go programs.

## Go Modules and GOPATH:
- Since Go 1.11+, **Go Modules** has reduced the dependency on `GOPATH`. With Go Modules, you can manage dependencies outside of `GOPATH` by using `go.mod` files. Go modules allow you to work on projects anywhere on your file system, without being restricted to the `GOPATH/src` directory.

