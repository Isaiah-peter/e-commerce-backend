# 🛒 E-Commerce Backend (Go)

A scalable **e-commerce backend API** built with **Golang**, featuring authentication, product management, shopping cart, orders, and planned payment integration.  

---

## 🚀 Features


# 🛒 E-Commerce Backend (Go)

This project is a scalable backend API for an e-commerce platform, built with Go. It supports user authentication, product management (with multiple images via Cloudinary), cart and order management, and is ready for payment integration.

- **Orders**
  - Place orders from cart

  # E-Commerce Backend (Go)

  This is a backend API for an e-commerce platform built with Go, GORM, Gorilla Mux, and Cloudinary for image uploads. It supports user authentication, product management, cart, orders, and more.

  ## Features
  - User authentication (OAuth, JWT)
  - Product CRUD with support for multiple images (Cloudinary integration)
  - Category, Color, and Size management
  - Cart and Order management
  - Stripe payment integration
  - Secure password hashing
  - RESTful API design

  ## Tech Stack
  - Go (Golang)
  - GORM (ORM)
  - Gorilla Mux (Router)
  - Cloudinary (Image hosting)
  - Stripe (Payments)
  - JWT (Authentication)

  ## Getting Started

  ### Prerequisites
  - Go 1.24+
  - Cloudinary account (for image uploads)
  - Stripe account (for payments)
  - Database (e.g., PostgreSQL, MySQL)

  ### Installation
  1. Clone the repository:
    ```bash
    git clone https://github.com/Isaiah-peter/e-commerce-backend.git
    cd e-commerce-backend
    ```
  2. Install dependencies:
    ```bash
    go mod tidy
    ```
  3. Set environment variables:
    - `CLOUDINARY_URL` (Cloudinary credentials)
    - Database credentials (see `pkg/config/app.go`)
    - Stripe keys

  ### Running the Server
  ```bash
  go run cmd/main.go
  ```

  ## API Endpoints

  ### Product
  - `POST /products` - Create product (supports 1-6 images)
  - `PUT /products/{id}` - Update product (supports image update)
  - `GET /products` - List products
  - `GET /products/{id}` - Get product by ID
  - `DELETE /products/{id}` - Delete product

  ### Category, Color, Size
  - CRUD endpoints for each

  ### Cart & Order
  - Add to cart, checkout, order history

  ### Auth
  - Register, login, OAuth

  ## Image Uploads
  - Products support 1-6 images per product
  - Images are uploaded to Cloudinary via multipart form-data (`images` key)

  ## Stripe Payments
  - Integrated for order checkout

  ## Contributing
  Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

  ## License
  MIT

  ## Author
  Isaiah Peter

---

## 🧪 Testing

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
| POST   | `/products` (Admin)   | Add a new product (1-6 images)|
| PUT    | `/products/{id}`      | Update product (images)      |
| GET    | `/cart`               | Get current user’s cart      |
| POST   | `/orders`             | Place a new order            |

---

## 🛡️ Security

- Passwords hashed with bcrypt
- JWT authentication
- Role-based access middleware

---

## 🤝 Contributing

1. Fork the repo
2. Create a new branch
3. Commit changes
4. Open a PR

---

## 📜 License

MIT License © 2025 Isaiah Peter
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
