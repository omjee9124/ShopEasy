# Makefile for E-Commerce Shopping Cart

.PHONY: help build run test clean docker-build docker-run setup-frontend install-deps

# Default target
help:
	@echo "E-Commerce Shopping Cart - Available Commands:"
	@echo ""
	@echo "Backend Commands:"
	@echo "  make run-backend     - Run the Go backend server"
	@echo "  make build-backend   - Build the Go backend binary"
	@echo "  make test-backend    - Run Go tests"
	@echo "  make clean-backend   - Clean backend build artifacts"
	@echo ""
	@echo "Frontend Commands:"
	@echo "  make setup-frontend  - Initialize React frontend"
	@echo "  make run-frontend    - Run the React development server"
	@echo "  make build-frontend  - Build React for production"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make docker-build    - Build Docker images"
	@echo "  make docker-run      - Run with Docker Compose"
	@echo "  make docker-stop     - Stop Docker containers"
	@echo "  make docker-clean    - Clean Docker containers and images"
	@echo ""
	@echo "Development Commands:"
	@echo "  make install-deps    - Install all dependencies"
	@echo "  make dev            - Run both backend and frontend in development"
	@echo "  make clean          - Clean all build artifacts"
	@echo "  make test           - Run all tests"

# Backend Commands
run-backend:
	@echo "Starting Go backend server..."
	cd backend && go run main.go

build-backend:
	@echo "Building Go backend..."
	cd backend && go build -o bin/ecommerce-backend main.go

test-backend:
	@echo "Running Go tests..."
	cd backend && go test -v ./...

clean-backend:
	@echo "Cleaning backend artifacts..."
	cd backend && rm -rf bin/ ecommerce.db

install-backend-deps:
	@echo "Installing Go dependencies..."
	cd backend && go mod tidy

# Frontend Commands
setup-frontend:
	@echo "Setting up React frontend..."
	npx create-react-app frontend
	cd frontend && npm install lucide-react
	cd frontend && npm install -D tailwindcss postcss autoprefixer
	cd frontend && npx tailwindcss init -p
	@echo "Frontend setup complete. Please copy the React component and configure Tailwind CSS."

run-frontend:
	@echo "Starting React development server..."
	cd frontend && npm start

build-frontend:
	@echo "Building React for production..."
	cd frontend && npm run build

install-frontend-deps:
	@echo "Installing React dependencies..."
	cd frontend && npm install

# Docker Commands
docker-build:
	@echo "Building Docker images..."
	docker-compose build

docker-run:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-stop:
	@echo "Stopping Docker containers..."
	docker-compose down

docker-clean:
	@echo "Cleaning Docker containers and images..."
	docker-compose down --rmi all --volumes --remove-orphans

# Development Commands
install-deps: install-backend-deps install-frontend-deps
	@echo "All dependencies installed!"

dev:
	@echo "Starting development environment..."
	@echo "Note: Run 'make run-backend' in one terminal and 'make run-frontend' in another"

clean: clean-backend
	@echo "Cleaning frontend artifacts..."
	cd frontend && rm -rf build/ node_modules/.cache/

test: test-backend
	@echo "Running frontend tests..."
	cd frontend && npm test -- --coverage --watchAll=false

# Database Commands
reset-db:
	@echo "Resetting database..."
	cd backend && rm -f ecommerce.db
	@echo "Database reset. It will be recreated on next server start."

# Quick Start Commands
quick-start-backend:
	@echo "Quick starting backend..."
	mkdir -p backend
	cd backend && go mod init ecommerce-backend
	@echo "Copy main.go and go.mod contents, then run 'make install-backend-deps && make run-backend'"

quick-start-frontend:
	@echo "Quick starting frontend..."
	@echo "Run 'make setup-frontend' first, then copy the React component"

# Production Commands
build-all: build-backend build-frontend
	@echo "Production builds complete!"

deploy-prep: clean build-all test
	@echo "Deployment preparation complete!"
	@echo "Backend binary: backend/bin/ecommerce-backend"
	@echo "Frontend build: frontend/build/"

# API Testing
test-api:
	@echo "Testing API endpoints..."
	@echo "Make sure the backend is running on port 8080"
	curl -X GET http://localhost:8080/items || echo "Backend not running"
	@echo "\nImport the Postman collection for comprehensive API testing"

# Development Server Status
status:
	@echo "Checking service status..."
	@echo "Backend (port 8080):"
	@curl -s http://localhost:8080/items > /dev/null && echo "✅ Backend is running" || echo "❌ Backend is not running"
	@echo "Frontend (port 3000):"
	@curl -s http://localhost:3000 > /dev/null && echo "✅ Frontend is running" || echo "❌ Frontend is not running"

# Help for specific components
help-backend:
	@echo "Backend Help:"
	@echo "  Framework: Gin (Go)"
	@echo "  Database: SQLite with GORM"
	@echo "  Authentication: JWT tokens"
	@echo "  Port: 8080"
	@echo "  Test user: testuser/password"

help-frontend:
	@echo "Frontend Help:"
	@echo "  Framework: React"
	@echo "  Styling: Tailwind CSS"
	@echo "  Icons: Lucide React"
	@echo "  Port: 3000"
	@echo "  Login: testuser/password"

help-api:
	@echo "API Endpoints:"
	@echo "  POST /users - Create user"
	@echo "  POST /users/login - Login"
	@echo "  GET /items - List items"
	@echo "  POST /carts - Add to cart (auth required)"
	@echo "  POST /orders - Create order (auth required)"
	@echo "  Use the Postman collection for testing"