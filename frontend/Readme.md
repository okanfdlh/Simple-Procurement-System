Frontend – Simple Procurement App

Frontend sederhana untuk sistem procurement menggunakan HTML, jQuery, dan Tailwind CSS.

🚀 Tech Stack

HTML5

Tailwind CSS (CDN)

jQuery 3.7

Docker + Nginx

📁 Struktur Folder
frontend/
├── index.html        # Login page
├── dashboard.html    # Items list
├── purchase.html     # Create purchase
├── js/
│   ├── api.js        # AJAX helper
│   ├── auth.js       # Login logic
│   ├── auth-guard.js # JWT guard
│   ├── dashboard.js
│   └── purchase.js
├── Dockerfile

▶️ Menjalankan Frontend (Docker)
docker compose up -d --build frontend


Akses:

http://localhost:8080

🔐 Authentication Flow

Login → /api/login

JWT disimpan di localStorage

Semua request pakai Bearer Token

Guard otomatis redirect ke login jika token invalid

🛒 Fitur Frontend

✅ Login
✅ Dashboard Item
✅ Create Purchase
✅ Cart Management
✅ JWT Guard
✅ Tailwind UI

🔁 API Endpoint yang Digunakan
Method	Endpoint
POST	/api/login
GET	/api/items
GET	/api/suppliers
POST	/api/purchasing
🧠 Notes Penting

Frontend tidak build step

Cocok untuk demo / admin internal

API base URL ada di js/api.js

🧪 Testing Login
{
  "username": "admin",
  "password": "password123"
}

🐳 Full Stack Run (Recommended)
docker compose up -d --build


Akses:

Frontend → http://localhost:8080

Backend → http://localhost:3000

Swagger → http://localhost:3000/swagger/index.html