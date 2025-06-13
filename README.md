# 🧠 Go Auth App

Sebuah boilerplate project untuk aplikasi **RESTful API** berbasis **Golang**, menggunakan **Fiber**, **GORM**, dan **Goose** untuk migrasi database. Proyek ini dibuat sebagai fondasi untuk membangun backend modern yang efisien dan ringan, dengan konsep ala Laravel.

---

## 🚀 Fitur

- ✨ HTTP Framework: [Fiber v2](https://github.com/gofiber/fiber)
- 🔐 JWT Authentication
- 🧰 Validasi dengan [validator.v10](https://github.com/go-playground/validator)
- 🧱 ORM dengan [GORM](https://gorm.io/)
- 🐘 PostgreSQL driver dan migrasi via [Goose](https://github.com/pressly/goose)
- ⚙️ Environment config dengan `godotenv`
- 📦 Struktur modular dan scalable (DDD-style)

---

## 🗂️ Struktur Direktori

```
go-auth/
├── cmd/
│   └── main.go                    # Entry point utama: inisialisasi Fiber, DB, config, middleware, dan route
│
├── config/
│   └── config.go                  # Load konfigurasi dari .env atau file config
│
├── database/
│   ├── migration/
│   │   └── 001_create_users.sql   # File migrasi (bisa raw SQL atau tool)
│   ├── seed/
│   │   └── user_seeder.go         # Seeder dummy data awal
│   └── factory/
│       └── user_factory.go        # Factory data palsu untuk testing
│
├── internal/                      # Folder utama untuk modul-modul aplikasi (DDD-style)
│   ├── middleware/                # Middleware global & modular (auth, logger, recover)
│   │   ├── auth.go
│   │   ├── logger.go
│   │   └── recovery.go
│   │
│   └── auth/                      # Modul auth (login, register, dsb)
│       ├── delivery/              # Lapisan komunikasi keluar (HTTP, gRPC, CLI, dsb)
│       │   ├── grpc/              # (Opsional) Handler untuk gRPC jika diperlukan
│       │   │   ├── grpc_handler.go
│       │   └── http/              # Handler untuk HTTP (pakai Fiber)
│       │       ├── controller/    # Handler controller HTTP endpoint
│       │       ├── request/       # DTO/request validator (mirip Laravel FormRequest)
│       │       └── response/      # Optional: Format response spesifik (bisa pakai resource juga)
│       │
│       ├── domain/                # Berisi interface dan struct utama domain (User, dsb)
│       │   └── user.go
│       │
│       ├── repository/            # Implementasi interface data access
│       │   ├── mysql/             # Menggunakan MySQL dan GORM
│       │   |   └── user_repository.go
│       │   └── pgsql/             # Menggunakan Pgsql dan GORM
│       │       └── user_repository.go
│       │
│       ├── service/               # Logika bisnis utama
│       │   └── auth_service.go
│       │
│       ├── resource/              # Format standar untuk JSON response (mirip Laravel Resource)
│       │   └── user_resource.go
│       │
│       └── routes/                # Routing lokal modul
│           └── routes.go          # Fungsi RegisterRoutes(router fiber.Router)
│
├── pkg/                           # Utilities (JWT, Hash, Response, Time, dsb)
│   ├── jwt.go
│   ├── hash.go
│   ├── validator.go
│   └── response.go
│
├── shared/                        # Struct & helper yang bisa dipakai lintas modul
│   └── dto.go
│
├── routes/
│   └── routes.go                  # Global route registry: panggil semua RegisterRoutes() dari modul
```

---

## 📦 Dependencies Utama

- `github.com/gofiber/fiber/v2`
- `github.com/golang-jwt/jwt/v5`
- `gorm.io/gorm`
- `gorm.io/driver/postgres`
- `github.com/pressly/goose/v3`
- `github.com/joho/godotenv`
- `github.com/go-playground/validator/v10`

---

## 🛠️ Instalasi

1. Clone repository:

```bash
git clone https://github.com/ZAIDALMAGHFUR/go-auth.git
cd go-auth
```

2. Copy file `.env` contoh:

```bash
cp .env.example .env
```

3. Edit `.env` kamu:

```env
DB_URL=postgres://postgres:root@localhost:5432/go-auth?sslmode=disable
MIGRATION_DIR=database/migration
SEED_DIR=database/seed
```

4. Install dependency Go:

```bash
go mod tidy
```

5. Install `goose` (jika belum ada):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

---

## 🧬 Migrasi Database

### 🔨 Buat File Migrasi Baru

```bash
make Migration name=create_roles_table
```

### ▶️ Jalankan Migrasi

```bash
make Migrate
```

### 🔁 Rollback Migrasi

```bash
make rollback
```

---

## 🌱 Seeder Database

### 🔨 Buat Seeder Baru

```bash
make Seeder name=seed_roles
```

### ▶️ Jalankan Seeder

```bash
make seed
```

### 🔁 Rollback Seeder

```bash
make rollback-seed
```

---

## 🧪 Menjalankan Aplikasi

### 🔥 Run Production Mode

```bash
make serve
```

### 👨‍💻 Run Development Mode (Hot reload via Air)

```bash
make dev
```

> Pastikan kamu sudah menginstall [Air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
```

---

## 📜 Lisensi

Standart Arsitektur is proprietary software developed by [PT Mac Tech Inv]. All rights reserved.

This software cannot be used, modified, or distributed without explicit permission from [PT Mac Tech Inv]. For more details, see the [LICENSE](LICENSE) file.

![License](https://img.shields.io/badge/license-proprietary-red)

---

## 📫 Kontak

Jika kamu ingin berdiskusi atau kontribusi, hubungi:  
📧 [mectechinv@gmail.com](mailto:mectechinv@gmail.com)
