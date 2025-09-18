
# 🛒 E-Commerce Backend (Go)

A scalable e-commerce backend API built with Golang, featuring authentication, product management, shopping cart, orders, and planned payment integration.

## 🚀 Features

- Authentication & Authorization (JWT, role-based)
- Product CRUD (supports multiple images via Cloudinary)
- Category, Color, Size management
- Shopping cart & order management
- Secure password hashing
- RESTful API design
- Stripe payment integration (planned)
- Docker support

## 📂 Project Structure

```
cmd/                # Entry point (main.go)
pkg/
  config/           # App & DB configuration
  routes/           # API routes
  controllers/      # HTTP request handlers
  models/           # Database schemas
  util/             # JWT, hashing, helpers
README.md
go.mod / go.sum
```

## ⚙️ Installation & Setup

### Prerequisites
- Go 1.21+
- PostgreSQL or MySQL
- Cloudinary account (for image uploads)
- Stripe account (for payments, optional)
- Docker (optional)

### Clone & Install
```bash
git clone https://github.com/Isaiah-peter/e-commerce-backend.git
cd e-commerce-backend
go mod tidy
```

### Environment Variables
Create a `.env` file:
```
PORT=8080
DB_URL=postgres://username:password@localhost:5432/ecommerce?sslmode=disable
JWT_SECRET=your_jwt_secret_key
CLOUDINARY_URL=cloudinary://<api_key>:<api_secret>@<cloud_name>
STRIPE_KEY=your_stripe_key
```

## ▶️ Running the App

```bash
go run cmd/main.go
```
Or with Docker:
```bash
docker-compose up --build
```

## 🧪 Testing

```bash
go test ./... -v
```

## 📖 API Endpoints (Sample)

| Method | Endpoint              | Description                  |
|--------|-----------------------|------------------------------|
| POST   | `/auth/register`      | Register a new user          |
| POST   | `/auth/login`         | Login and get JWT token      |
| GET    | `/products`           | Get all products             |
| POST   | `/products` (Admin)   | Add a new product (1-6 images)|
| PUT    | `/products/{id}`      | Update product (images)      |
| GET    | `/cart`               | Get current user’s cart      |
| POST   | `/orders`             | Place a new order            |

## 🛡️ Security

- Passwords hashed with bcrypt
- JWT authentication
- Role-based access middleware

## 🤝 Contributing

1. Fork the repo
2. Create a new branch
3. Commit changes
4. Open a PR

## 📜 License

MIT License © 2025 Isaiah Peter
