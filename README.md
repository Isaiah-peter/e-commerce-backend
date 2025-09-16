# 🛒 E-Commerce Backend (Go)

A scalable **e-commerce backend API** built with **Golang**, featuring authentication, product management, shopping cart, orders, and planned payment integration.  

---

## 🚀 Features

- **Authentication & Authorization**
  - User registration & login (JWT-based)
  - Role-based access control (Admin, Customer)

- **Product Management**
  - CRUD operations for products & categories
  - Inventory tracking

- **Shopping Cart**
  - Add / update / remove items
  - Cart persistence per user

- **Orders**
  - Place orders from cart
  - Order history & status updates

- **Utilities**
  - Secure password hashing (bcrypt)
  - JWT-based authentication
  - Structured JSON responses

- **Planned**
  - Payment gateway integration (Stripe, PayPal, Flutterwave)
  - Swagger API documentation
  - CI/CD pipeline

---

## 📂 Project Structure

```
.
├── cmd/                # Entry point (main.go)
├── pkg/
│   ├── config/         # App & DB configuration
│   ├── routes/         # API routes
│   ├── controllers/    # HTTP request handlers
│   ├── services/       # Business logic
│   ├── repository/     # Database queries
│   ├── models/         # Database schemas
│   ├── middleware/     # Auth, logging, etc.
│   └── util/           # JWT, hashing, helpers
├── tests/              # Unit & integration tests
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
└── README.md
```

---

## ⚙️ Installation & Setup

### Prerequisites
- [Go 1.21+](https://go.dev/dl/)  
- [PostgreSQL](https://www.postgresql.org/) or MySQL  
- [Docker](https://www.docker.com/) (optional)

### Clone the Repository
```bash
git clone https://github.com/Isaiah-peter/e-commerce-backend.git
cd e-commerce-backend
```

### Install Dependencies
```bash
go mod tidy
```

### Environment Variables
Create a `.env` file in the root directory:

```env
PORT=8080
DB_URL=postgres://username:password@localhost:5432/ecommerce?sslmode=disable
JWT_SECRET=your_jwt_secret_key
```

You can also copy from the example file:
```bash
cp .env.example .env
```

---

## ▶️ Running the App

### Run Locally
```bash
go run cmd/main.go
```

### Run with Docker
```bash
docker-compose up --build
```

The server will start at **http://localhost:8080**

---

## 🧪 Testing

Run tests with:
```bash
go test ./... -v
```

---

## 📖 API Endpoints (Sample)

| Method | Endpoint              | Description                  |
|--------|-----------------------|------------------------------|
| POST   | `/auth/register`      | Register a new user          |
| POST   | `/auth/login`         | Login and get JWT token      |
| GET    | `/products`           | Get all products             |
| POST   | `/products` (Admin)   | Add a new product            |
| GET    | `/cart`               | Get current user’s cart      |
| POST   | `/orders`             | Place a new order            |

---

## 🛡️ Security

- Passwords hashed using **bcrypt**  
- JWT-based authentication with expiry  
- Middleware for role-based access  

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)  
- **Web Framework**: Gin / Fiber (depending on implementation)  
- **Database**: PostgreSQL (via GORM)  
- **Auth**: JWT  
- **Containerization**: Docker  

---

## 🤝 Contributing

1. Fork the repo  
2. Create a new branch (`feature/awesome-feature`)  
3. Commit changes  
4. Push branch and open a PR  

---

## 📜 License

MIT License © 2025 [Isaiah Peter](https://github.com/Isaiah-peter)
