Backend – Simple Procurement API

Backend API untuk sistem procurement sederhana menggunakan Go (Fiber), MySQL, JWT Authentication, dan Swagger.

🚀 Tech Stack

Go 1.22+

Fiber v2

GORM

MySQL 8

JWT Authentication

Swagger (swaggo)

Docker & Docker Compose

📁 Struktur Folder
backend/
├── config/         # Database & env config
├── controllers/    # Logic API
├── middlewares/    # JWT middleware
├── models/         # Database models
├── routes/         # API routes
├── seeders/        # Database seeders
├── utils/          # Helper & response
├── docs/           # Swagger docs
├── main.go
├── go.mod
├── Dockerfile

⚙️ Environment Variables

Backend menggunakan ENV, contoh:

DB_HOST=mysql
DB_PORT=3306
DB_USER=app
DB_PASSWORD=app
DB_NAME=inventory_db
JWT_SECRET=supersecret


Saat menggunakan Docker, ENV sudah di-handle oleh docker-compose.yml.

▶️ Menjalankan Backend (Docker)
docker compose up -d --build backend


Akses API:

http://localhost:3000/api

📖 Swagger Documentation

Swagger otomatis tersedia di:

http://localhost:3000/swagger/index.html


Fitur:

Auth (Register, Login)

Items

Suppliers

Purchasing

JWT Authorization

🔐 Authentication

Gunakan Bearer Token pada header:

Authorization: Bearer <JWT_TOKEN>

🧪 Seeder Otomatis

Saat backend pertama kali dijalankan:

Admin user otomatis dibuat

Data supplier & item sample dibuat

Seeder tidak akan duplikat

🛠️ Command Lokal (Tanpa Docker)
go mod tidy
go run main.go

📌 API Response Format

Semua response menggunakan format konsisten:

{
  "success": true,
  "message": "Success message",
  "data": {}
}

🧠 Author

Indirokan Fadhilah
Fullstack Developer Intern Candidate