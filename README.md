# Go Learning & Practice Repository ("Gold Standard")

Репозиторій містить поетапний практичний курс із вивчення мови Go (Golang): від базового синтаксису до обробки помилок, конкурентності, REST API та Clean Architecture.

---

## 📂 Структура навчального плану (Modules 01–30)

### 🔹 Основи мови та типізація (01–07)
- `01_bank_account` — Змінні, базові типи, структурні основи та керування потоком.
- `02_validation` — Умовні оператори, стрінги та валідація вхідних даних.
- `03_structs_methods` — Структури, методи зі значеннями та вказівниками.
- `04_interfaces` — Базові інтерфейси (`Payer`, `Kreditkarte`).
- `05_errors` — Базова обробка помилок та кастомні помилки (`GuthabenUeberzogenError`).
- `06_packages` — Організація Go-модулів, видимість (`export/unexport`), `go.mod`.
- `07_pointers_receivers` — Вказівники, мутації, мутабельні методи.

### 🔹 Поглиблений ООП та Error Handling (08–09)
- `08_interfaces` — Композиція інтерфейсів, `type assertion`, порожній інтерфейс `any`.
- `09_errors` — Обгортання помилок (`fmt.Errorf` з `%w`), `errors.Is`, `errors.As`.

### 🔹 Конкурентність / Concurrency (10–15)
- `10_goroutines_sync` — Конкурентність: Goroutines, недетермінованість та `sync.WaitGroup`.
- `11_channels` — Незабуферизовані канали, синхронізація через передачу повідомлень.
- `12_buffered_channels` — Буферизовані канали, пропускна здатність, неблокуюче читання/запис.
- `13_select` — Мультиплексування каналів за допомогою `select`, таймаути (`time.After`), `default`.
- `14_context` — `context.WithTimeout`, `context.WithCancel`, скасування довготривалих операцій.
- `15_sync_mutex` — Спільний стан, race conditions, `sync.Mutex`, `sync.RWMutex`, перевірка `-race`.

### 🔹 Продвинута конкурентність та паттерни (16–20)
- `16_worker_pool` — Паттерн Worker Pool (розподіл задач між кількома горутинами).
- `17_pipeline` — Паттерн Pipeline (конвеєрна обробка даних через канали).
- `18_fan_in_fan_out` — Паралельна обробка даних (Fan-Out) та їхнє об'єднання (Fan-In).
- `19_atomic_operations` — Низькорівнева атомарна синхронізація (`sync/atomic`).
- `20_rate_limiting` — Обмеження частоти запитів (Rate Limiting) за допомогою `time.Ticker` та каналів.

### 🔹 Стандартна бібліотека та веб-сервіси (21–25)
- `21_json_marshaling` — Робота з JSON: енкодинг, декодинг, теги структур (`json:"key"`).
- `22_http_server` — Базовий HTTP-сервер на `net/http`, роутинг, `http.HandlerFunc`.
- `23_http_middleware` — Створення Middleware (логи, автентифікація, обробка панік).
- `24_http_client` — Написання HTTP-клієнта, обробка відповідей, таймаути запитів.
- `25_rest_api_crud` — Повноцінне REST API зі збереженням даних в пам'яті (In-Memory CRUD).

### 🔹 Архітектура, БД та тестування (26–30)
- `26_unit_testing` — Юніт-тестування (`testing` пакет), Table-Driven Tests.
- `27_mocking` — Інтерфейси для мокінгу залежностей у тестах.
- `28_database_sql` — Робота з SQL (PostgreSQL/SQLite) через `database/sql` та транзакції.
- `29_clean_architecture` — Шари додатку: Handler -> Service -> Repository.
- `30_final_project` — Фінальний банківський microservice з REST API, базами даних, контекстами та повним тестовим покриттям.

---

## 🛠 Технології
- **Language**: Go (Golang)
- **VCS**: Git & GitHub