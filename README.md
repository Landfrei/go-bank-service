# Go Learning & Practice Repository ("Gold Standard")

This repository contains a step-by-step practical course for learning Go (Golang): from basic syntax and error handling to concurrency, REST API design, and Clean Architecture.

---

## 📂 Course Roadmap (Modules 01–30)

### 🔹 Language Basics & Typing (01–07)
- `01_bank_account` — Variables, basic types, struct fundamentals, and control flow.
- `02_validation` — Conditional logic, string manipulation, and input validation.
- `03_structs_methods` — Structs, value receivers, and pointer receivers.
- `04_interfaces` — Basic interfaces (`Payer`, `Kreditkarte`).
- `05_errors` — Basic error handling and custom error types (`GuthabenUeberzogenError`).
- `06_packages` — Go module organization, visibility rules (`export/unexport`), and `go.mod`.
- `07_pointers_receivers` — Pointers, memory mutation, and mutable methods.

### 🔹 Advanced OOP & Error Handling (08–09)
- `08_interfaces` — Interface composition, type assertions, and empty interface `any`.
- `09_errors` — Error wrapping (`fmt.Errorf` with `%w`), `errors.Is`, and `errors.As`.

### 🔹 Concurrency (10–15)
- `10_goroutines_sync` — Goroutines, non-determinism, and `sync.WaitGroup`.
- `11_channels` — Unbuffered channels and synchronization via message passing.
- `12_buffered_channels` — Buffered channels, throughput management, and non-blocking operations.
- `13_select` — Channel multiplexing with `select`, timeouts (`time.After`), and `default` cases.
- `14_context` — `context.WithTimeout`, `context.WithCancel`, and long-running task cancellation.
- `15_sync_mutex` — Shared state management, race conditions, `sync.Mutex`, `sync.RWMutex`, and `-race` detector.

### 🔹 Advanced Concurrency Patterns (16–20)
- `16_worker_pool` — Worker Pool pattern (distributing tasks across multiple goroutines).
- `17_pipeline` — Pipeline pattern (streaming data processing through channels).
- `18_fan_in_fan_out` — Parallel execution (Fan-Out) and result aggregation (Fan-In).
- `19_atomic_operations` — Low-level atomic synchronization (`sync/atomic`).
- `20_rate_limiting` — Rate limiting with `time.Ticker` and channels.

### 🔹 Standard Library & Web Services (21–25)
- `21_json_marshaling` — Working with JSON: encoding, decoding, and struct tags (`json:"key"`).
- `22_http_server` — Basic HTTP server with `net/http`, routing, and `http.HandlerFunc`.
- `23_http_middleware` — Creating middleware (logging, authentication, panic recovery).
- `24_http_client` — Building HTTP clients, response parsing, and request timeouts.
- `25_rest_api_crud` — Fully functional In-Memory REST API (CRUD operations).

### 🔹 Architecture, Databases & Testing (26–30)
- `26_unit_testing` — Unit testing with `testing` package and Table-Driven Tests.
- `27_mocking` — Interfaces for dependency mocking in tests.
- `28_database_sql` — SQL databases (PostgreSQL/SQLite) via `database/sql` and transactions.
- `29_clean_architecture` — Layered architecture: Handler -> Service -> Repository.
- `30_final_project` — Production-grade banking microservice with REST API, database persistence, context support, and full test coverage.

---

## 🛠 Tech Stack
- **Language**: Go (Golang)
- **VCS**: Git & GitHub

  ### Статистика репозиторію

![GitHub Stats](https://vercel.app)

