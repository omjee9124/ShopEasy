# 🛒 E-Commerce Shopping Cart System

A modern, full-stack e-commerce shopping cart application built with **Go (Gin)** backend and **React** frontend, featuring JWT authentication, real-time cart management, and order processing.

![Project Status](https://img.shields.io/badge/Status-Complete-brightgreen)
![Backend](https://img.shields.io/badge/Backend-Go%20%2B%20Gin-blue)
![Frontend](https://img.shields.io/badge/Frontend-React%20%2B%20Tailwind-61dafb)
![Database](https://img.shields.io/badge/Database-SQLite-003b57)

## 🌟 Features

### 🔐 Authentication & Security
- JWT-based user authentication
- Secure password hashing with bcrypt
- Protected API routes
- Single-device login sessions

### 🛍️ Shopping Experience
- Modern, responsive user interface
- Product catalog with beautiful card layouts
- Real-time cart updates with item counters
- One-click add to cart functionality
- Seamless checkout process

### 📋 Order Management
- Convert cart to order with single click
- Complete order history tracking
- Order details with item information
- User profile management

### 🔧 Technical Features
- RESTful API design
- CORS-enabled for seamless frontend integration
- Automatic database migration
- Sample data seeding
- Comprehensive error handling

## 🏗️ Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   React Client  │────▶│   Go API Server │────▶│  SQLite Database│
│   (Port 3000)   │     │   (Port 8080)   │     │   (ecommerce.db)│
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

## 📊 Database Schema

```sql
Users                    Items                   Carts
├── id (PK)             ├── id (PK)             ├── id (PK)
├── username (Unique)   ├── name                ├── user_id (FK)
├── password (Hashed)   ├── status              ├── name
├── token               └── created_at          ├── status
├── cart_id (FK)                                └── created_at
└── created_at          

CartItems               Orders
├── cart_id (PK, FK)   ├── id (PK)
├── item_id (PK, FK)   ├── cart_id (FK)
└── item (Relation)    ├── user_id (FK)
                       └── created_at
```

## 🚀 Quick Start

### Prerequisites
- **Go 1.19+** - [Download](https://golang.org/dl/)
- **Node.js 16+** - [Download](https://nodejs.org/)
- **Git** - [Download](https://git-scm.com/)

### 1️⃣ Clone Repository
```bash
git clone (https://github.com/omjee9124/ShopEasy.git)
cd ecommerce-shopping-cart
```

### 2️⃣ Backend Setup
```bash
# Navigate to backend directory
cd backend

# Initialize Go modules
go mod tidy

# Run the server
go run main.go
```
✅ **Backend running on:** `http://localhost:8080`

### 3️⃣ Frontend Setup
```bash
# Navigate to frontend directory (in new terminal)
cd frontend

# Install dependencies
npm install lucide-react

# Install Tailwind CSS
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

**Configure Tailwind CSS:**

Update `tailwind.config.js`:
```javascript
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: { extend: {} },
  plugins: [],
}
```

Update `src/index.css`:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

```bash
# Start the React app
npm start
```
✅ **Frontend running on:** `http://localhost:3000`

## 🎯 Usage Guide

### 🔑 Login Credentials
```
Username: testuser
Password: password
```

### 🛒 Shopping Flow
1. **Login** → Use test credentials or create new account
2. **Browse** → View products on home page
3. **Add to Cart** → Click "Add to Cart" on desired items
4. **Review Cart** → Click cart icon to view items
5. **Checkout** → Click "Place Order" to complete purchase
6. **Order History** → View all past orders

## 📡 API Documentation

### 🔓 Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/users` | Create new user |
| `GET` | `/users` | List all users |
| `POST` | `/users/login` | User authentication |
| `POST` | `/items` | Create new item |
| `GET` | `/items` | List all items |

### 🔒 Protected Endpoints (Require JWT Token)

| Method | Endpoint | Description | Headers |
|--------|----------|-------------|---------|
| `POST` | `/carts` | Add items to cart | `Authorization: Bearer <token>` |
| `GET` | `/carts` | Get user's carts | `Authorization: Bearer <token>` |
| `POST` | `/orders` | Create order from cart | `Authorization: Bearer <token>` |
| `GET` | `/orders` | Get user's orders | `Authorization: Bearer <token>` |

### 📝 Request Examples

**User Login:**
```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password"}'
```

**Add Items to Cart:**
```bash
curl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"item_ids":[1,2,3]}'
```

**Create Order:**
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"cart_id":1}'
```

## 🗂️ Project Structure

```
ecommerce-shopping-cart/
├── backend/
│   ├── main.go              # Main Go application
│   ├── go.mod              # Go dependencies
│   └── ecommerce.db        # SQLite database (auto-generated)
│
├── frontend/
│   ├── src/
│   │   ├── App.js          # Main React component
│   │   ├── index.css       # Tailwind CSS imports
│   │   └── index.js        # React entry point
│   ├── package.json        # Node dependencies
│   └── tailwind.config.js  # Tailwind configuration
│
└── README.md               # This file
```

## ✅ Testing Checklist

### Manual Testing
- [ ] User registration works
- [ ] Login with valid credentials succeeds
- [ ] Login with invalid credentials fails
- [ ] Items display correctly with images
- [ ] Add to cart functionality works
- [ ] Cart counter updates in real-time
- [ ] Cart page shows added items
- [ ] Checkout creates order successfully
- [ ] Order history displays correctly
- [ ] Logout clears session

### API Testing
Use the provided curl commands or import into Postman for comprehensive API testing.

## 🛠️ Technical Decisions

| Component | Technology | Rationale |
|-----------|------------|-----------|
| **Backend** | Go + Gin | High performance, excellent concurrency |
| **Database** | SQLite + GORM | Simple setup, perfect for development |
| **Frontend** | React | Component-based, excellent ecosystem |
| **Styling** | Tailwind CSS | Utility-first, rapid development |
| **Authentication** | JWT | Stateless, scalable |
| **State Management** | React Hooks | Built-in, no additional dependencies |

## 🚨 Troubleshooting

### Common Issues & Solutions

**🔸 CORS Errors**
```bash
# Ensure backend allows frontend origin
# Check CORS configuration in main.go
```

**🔸 Token Issues**
```javascript
// Clear localStorage and login again
localStorage.clear();
```

**🔸 Database Issues**
```bash
# Delete database file and restart
rm ecommerce.db
go run main.go
```

**🔸 Port Conflicts**
```bash
# Kill processes on required ports
lsof -ti:8080 | xargs kill -9  # Backend
lsof -ti:3000 | xargs kill -9  # Frontend
```

## 🔮 Future Enhancements

- [ ] **Product Management** - Categories, search, filters
- [ ] **Inventory System** - Stock tracking, availability
- [ ] **Payment Integration** - Stripe, PayPal support
- [ ] **User Profiles** - Address management, preferences
- [ ] **Admin Dashboard** - Product/order management
- [ ] **Email Notifications** - Order confirmations
- [ ] **Mobile App** - React Native implementation
- [ ] **Analytics** - Sales reports, user behavior

## 🔒 Security Features

- ✅ Password hashing with bcrypt
- ✅ JWT token expiration (24 hours)
- ✅ User-specific data isolation
- ✅ Input validation and sanitization
- ✅ CORS configuration for security
- ✅ Protected route authentication

## 📱 Screenshots

### Login Screen
![Login](https://img.shields.io/badge/UI-Modern%20Login-blue)

### Product Catalog
![Products](https://img.shields.io/badge/UI-Product%20Grid-green)

### Shopping Cart
![Cart](https://img.shields.io/badge/UI-Interactive%20Cart-orange)

### Order History
![Orders](https://img.shields.io/badge/UI-Order%20Tracking-purple)

## 👨‍💻 Development

### Running in Development Mode
```bash
# Backend with auto-reload
go run main.go

# Frontend with hot reload
npm start
```

### Building for Production
```bash
# Backend binary
go build -o ecommerce-server main.go

# Frontend build
npm run build
```

## 📄 License

This project is developed for educational and assessment purposes.

---

**🎉 Project Status: Complete and Fully Functional**

*Built with ❤️ (by Om Ji Dubey)using Go, React, and modern web technologies*
