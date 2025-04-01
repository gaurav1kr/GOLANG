
# Golang Program: Watch a Directory for File Changes

This example uses the `fsnotify` package to watch a directory for file system changes like creation, modification, and deletion.

---

## 📦 Dependencies

You’ll need to install the `fsnotify` package:

```bash
go get github.com/fsnotify/fsnotify
```

---

## 📝 Code Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/fsnotify/fsnotify"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal("Failed to create watcher:", err)
    }
    defer watcher.Close()

    done := make(chan bool)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Printf("EVENT: %s %s\n", event.Op, event.Name)

                if event.Op&fsnotify.Create == fsnotify.Create {
                    fmt.Println("File created:", event.Name)
                }
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Println("File modified:", event.Name)
                }
                if event.Op&fsnotify.Remove == fsnotify.Remove {
                    fmt.Println("File removed:", event.Name)
                }

            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("ERROR:", err)
            }
        }
    }()

    path := "./watched" // Replace with the directory you want to watch
    err = watcher.Add(path)
    if err != nil {
        log.Fatal("Failed to watch directory:", err)
    }

    fmt.Println("Watching directory:", path)
    <-done
}
```

---

## ▶️ How to Run

1. Create a directory to watch:

```bash
mkdir watched
```

2. Save the code as `main.go`, and run:

```bash
go mod init watcher
go get github.com/fsnotify/fsnotify
go run main.go
```

3. Perform file operations inside the `watched` directory to see logs.

---

## ✅ Output Example

```bash
Watching directory: ./watched
EVENT: CREATE watched/test.txt
File created: watched/test.txt
```

