E-Commerce Shopping Cart System
A complete e-commerce shopping cart system built with Go (Gin framework) backend and React frontend.
Features
Backend (Go + Gin + GORM)

User registration and authentication with JWT tokens
Item management (CRUD operations)
Shopping cart functionality (one cart per user)
Order management (convert cart to order)
SQLite database with GORM ORM
CORS enabled for frontend integration

Frontend (React)

User login interface
Items listing with add to cart functionality
Cart management (view cart items)
Order history
Checkout functionality
Responsive UI with Tailwind CSS

Database Schema
sqlUsers:
- id (Primary Key)
- username (Unique)
- password (Hashed)
- token
- cart_id (Foreign Key to Carts)
- created_at

Items:
- id (Primary Key)
- name
- status
- created_at

Carts:
- id (Primary Key)
- user_id (Foreign Key to Users)
- name
- status
- created_at

CartItems:
- cart_id (Primary Key, Foreign Key to Carts)
- item_id (Primary Key, Foreign Key to Items)

Orders:
- id (Primary Key)
- cart_id (Foreign Key to Carts)
- user_id (Foreign Key to Users)
- created_at
API Endpoints
User Management

POST /users - Create a new user
GET /users - List all users
POST /users/login - User login (returns JWT token)

Item Management

POST /items - Create a new item
GET /items - List all items

Cart Management (Requires Authentication)

POST /carts - Create/update cart with items
GET /carts - Get user's carts

Order Management (Requires Authentication)

POST /orders - Convert cart to order
GET /orders - Get user's orders

Setup Instructions
Backend Setup

Prerequisites

Go 1.19 or higher
Git


Clone and Setup
bashmkdir ecommerce-backend
cd ecommerce-backend

# Create main.go and go.mod files (copy from artifacts)

# Initialize Go modules
go mod tidy

# Run the server
go run main.go

The server will start on port 8080

Database will be automatically created (ecommerce.db)
Sample data will be seeded automatically
Test user: username=testuser, password=password



Frontend Setup

Prerequisites

Node.js 16+ and npm
Create React App


Setup React Application
bashnpx create-react-app ecommerce-frontend
cd ecommerce-frontend

# Install required dependencies
npm install lucide-react

# Install Tailwind CSS
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p

Configure Tailwind CSS
Update tailwind.config.js:
javascriptmodule.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
Update src/index.css:
css@tailwind base;
@tailwind components;
@tailwind utilities;

Replace src/App.js with the React component from artifacts
Run the frontend
bashnpm start
The app will open at http://localhost:3000

Usage Instructions
Testing the Application

Start Backend Server
bashcd ecommerce-backend
go run main.go

Start Frontend Application
bashcd ecommerce-frontend
npm start

Login

Use credentials: testuser / password
Or create a new user via API


Shopping Flow

Browse items on the main screen
Click "Add to Cart" on any item
Use "Cart" button to view cart contents
Use "Checkout" to convert cart to order
Use "Order History" to view past orders



API Testing with curl

Create User
bashcurl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"newuser","password":"password123"}'

Login
bashcurl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password"}'

Get Items
bashcurl http://localhost:8080/items

Add Items to Cart (requires token)
bashcurl -X POST http://localhost:8080/carts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"item_ids":[1,2,3]}'

Create Order (requires token)
bashcurl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"cart_id":1}'


Project Structure
ecommerce-backend/
├── main.go           # Main application file
├── go.mod           # Go modules file
└── ecommerce.db     # SQLite database (auto-generated)

ecommerce-frontend/
├── src/
│   ├── App.js       # Main React component
│   ├── index.css    # Tailwind CSS imports
│   └── index.js     # React entry point
├── package.json
└── tailwind.config.js
Technical Decisions

Database: SQLite for simplicity and ease of setup
Authentication: JWT tokens stored in localStorage
CORS: Enabled for frontend-backend communication
Frontend State: React hooks for state management
UI: Tailwind CSS for responsive design
Error Handling: User-friendly alerts and error messages

Testing
Manual Testing Checklist

 User can register new account
 User can login with valid credentials
 Invalid login shows error message
 Items are displayed correctly
 Add to cart functionality works
 Cart shows added items
 Checkout converts cart to order
 Order history shows past orders
 User can logout successfully

API Testing
Use the provided curl commands or import the Postman collection to test all endpoints.
Security Considerations

Passwords are hashed using bcrypt
JWT tokens expire after 24 hours
Protected routes require valid authentication
User can only access their own carts and orders
CORS is configured for local development

Future Enhancements

Inventory management
Product images and descriptions
Shopping cart persistence
Payment integration
User profiles and addresses
Order tracking
Product categories and search
Admin panel for item management

Troubleshooting
Common Issues

CORS Errors: Ensure backend CORS is configured for frontend URL
Database Issues: Delete ecommerce.db and restart server to reset
Token Issues: Clear localStorage and login again
Port Conflicts: Ensure ports 8080 (backend) and 3000 (frontend) are free

Development Tips

Use browser developer tools to monitor API calls
Check server logs for detailed error messages
Test API endpoints with curl before frontend integration
Use React Developer Tools for debugging state issues


License
This project is for educational purposes and assessment.