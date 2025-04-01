
# Go Routine with Periodic Task and Graceful Shutdown

This Go program demonstrates how to run a periodic task every 10 seconds and gracefully shut it down when receiving an interrupt signal.

---

## 📦 Features

- Uses `time.Ticker` for periodic execution
- Handles `SIGINT`/`SIGTERM` with `os/signal`
- Uses `context.WithCancel` for coordinated shutdown

---

## 📝 Code Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    // Handle interrupt (Ctrl+C) or SIGTERM
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-sigs
        fmt.Println("\n🔴 Interrupt received. Shutting down gracefully...")
        cancel()
    }()

    go periodicTask(ctx)

    // Wait until context is canceled
    <-ctx.Done()
    fmt.Println("✅ Cleanup done. Exiting.")
}

func periodicTask(ctx context.Context) {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("⏹️ Stopping periodic task...")
            return
        case t := <-ticker.C:
            fmt.Printf("⏰ Task executed at %s\n", t.Format(time.RFC1123))
        }
    }
}
```

---

## ▶️ Example Output

```
⏰ Task executed at Mon, 01 Apr 2025 10:00:00 IST
⏰ Task executed at Mon, 01 Apr 2025 10:00:10 IST
^C
🔴 Interrupt received. Shutting down gracefully...
⏹️ Stopping periodic task...
✅ Cleanup done. Exiting.
```

---

## ✅ How to Run

```bash
go run main.go
```

Then press `Ctrl+C` to simulate a graceful shutdown.

