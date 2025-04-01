
# Managing Dependency Injection and Configuration in Go Microservices

This document outlines best practices for managing dependencies and configuration in Go microservices for clean, testable, and maintainable code.

---

## 🧱 1. Dependency Injection (DI) in Go

Go favors **explicit dependency passing** rather than using built-in DI frameworks.

---

### ✅ A. Constructor Injection

Inject dependencies using constructors.

```go
type UserService struct {
    repo UserRepository
    logger *log.Logger
}

func NewUserService(repo UserRepository, logger *log.Logger) *UserService {
    return &UserService{repo: repo, logger: logger}
}
```

---

### ✅ B. Interface-Based Injection

Define interfaces to allow mock implementations for testing.

```go
type UserRepository interface {
    GetUser(id string) (*User, error)
}
```

---

### ✅ C. DI Libraries (Optional)

| Library        | Description                          |
|----------------|--------------------------------------|
| `uber-go/fx`   | Uber’s full-featured DI framework     |
| `google/wire`  | Compile-time dependency injection     |
| `dig`          | Reflection-based container (used by fx)|

#### 🛠️ fx Example:
```go
app := fx.New(
    fx.Provide(NewUserService, NewUserRepository),
    fx.Invoke(StartServer),
)
```

---

## ⚙️ 2. Configuration Management

---

### ✅ A. Environment Variables (12-Factor Style)

Use `os.Getenv` or `envconfig`.

```go
type Config struct {
    Port     string `envconfig:"PORT" default:"8080"`
    LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

var cfg Config
envconfig.Process("", &cfg)
```

---

### ✅ B. Configuration Files (YAML, JSON)

Use libraries like `viper`, `go-yaml`.

```go
viper.SetConfigName("config")
viper.AddConfigPath(".")
viper.AutomaticEnv()

err := viper.ReadInConfig()
port := viper.GetString("server.port")
```

---

### ✅ C. Dynamic Config Reload

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config changed:", e.Name)
})
```

---

## 🧪 Testing with Injected Dependencies

Use mock implementations for easy testing.

```go
type MockRepo struct{}
func (m *MockRepo) GetUser(id string) (*User, error) {
    return &User{Name: "MockUser"}, nil
}

userService := NewUserService(&MockRepo{}, log.New(os.Stdout, "", 0))
```

---

## ✅ Summary Table

| Concern         | Approach                        | Tools/Libraries               |
|----------------|----------------------------------|-------------------------------|
| DI (manual)     | Constructor, interfaces          | N/A                           |
| DI (auto)       | Declarative or reflection-based  | `fx`, `wire`, `dig`           |
| Config (env)    | Read from env vars               | `envconfig`, `os.Getenv`      |
| Config (files)  | Load from YAML/JSON              | `viper`, `go-yaml`            |
| Config Reload   | Live reload                      | `viper`, `fsnotify`           |
