package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Models
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Token     string    `json:"token,omitempty"`
	CartID    *uint     `json:"cart_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Item struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:active"`
	CreatedAt time.Time `json:"created_at"`
}

type Cart struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	UserID    uint        `json:"user_id" gorm:"not null"`
	Name      string      `json:"name"`
	Status    string      `json:"status" gorm:"default:active"`
	Items     []CartItem  `json:"items,omitempty" gorm:"foreignKey:CartID"`
	CreatedAt time.Time   `json:"created_at"`
}

type CartItem struct {
	CartID uint `json:"cart_id" gorm:"primaryKey"`
	ItemID uint `json:"item_id" gorm:"primaryKey"`
	Item   Item `json:"item,omitempty" gorm:"foreignKey:ItemID"`
}

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CartID    uint      `json:"cart_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Cart      Cart      `json:"cart,omitempty" gorm:"foreignKey:CartID"`
	CreatedAt time.Time `json:"created_at"`
}

// Request/Response structs
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ItemRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

type CartRequest struct {
	ItemIDs []uint `json:"item_ids" binding:"required"`
}

type OrderRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
}

// Global variables
var db *gorm.DB
var jwtSecret = []byte("your-secret-key")

// JWT Claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func main() {
	// Initialize database
	initDB()

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	setupRoutes(r)

	// Seed some initial data
	seedData()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate schemas
	err = db.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// User routes
	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.POST("/users/login", loginUser)

	// Item routes
	r.POST("/items", createItem)
	r.GET("/items", getItems)

	// Cart routes (require authentication)
	r.POST("/carts", authMiddleware(), createCart)
	r.GET("/carts", authMiddleware(), getCarts)

	// Order routes (require authentication)
	r.POST("/orders", authMiddleware(), createOrder)
	r.GET("/orders", authMiddleware(), getOrders)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// User handlers
func createUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func loginUser(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/password"})
		return
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Update user token in database
	user.Token = tokenString
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"user_id":  user.ID,
		"username": user.Username,
	})
}

// Item handlers
func createItem(c *gin.Context) {
	var req ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Status == "" {
		req.Status = "active"
	}

	item := Item{
		Name:   req.Name,
		Status: req.Status,
	}

	if err := db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func getItems(c *gin.Context) {
	var items []Item
	db.Where("status = ?", "active").Find(&items)
	c.JSON(http.StatusOK, items)
}

// Cart handlers
func createCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req CartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already has an active cart
	var existingCart Cart
	if err := db.Where("user_id = ? AND status = ?", userID, "active").First(&existingCart).Error; err == nil {
		// User already has a cart, add items to existing cart
		for _, itemID := range req.ItemIDs {
			cartItem := CartItem{
				CartID: existingCart.ID,
				ItemID: itemID,
			}
			// Use FirstOrCreate to avoid duplicates
			db.Where(CartItem{CartID: existingCart.ID, ItemID: itemID}).FirstOrCreate(&cartItem)
		}

		// Return updated cart with items
		db.Preload("Items.Item").First(&existingCart, existingCart.ID)
		c.JSON(http.StatusOK, existingCart)
		return
	}

	// Create new cart
	cart := Cart{
		UserID: userID,
		Name:   "Cart for User " + strconv.Itoa(int(userID)),
		Status: "active",
	}

	if err := db.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	// Add items to cart
	for _, itemID := range req.ItemIDs {
		cartItem := CartItem{
			CartID: cart.ID,
			ItemID: itemID,
		}
		db.Create(&cartItem)
	}

	// Update user's cart_id
	db.Model(&User{}).Where("id = ?", userID).Update("cart_id", cart.ID)

	// Return cart with items
	db.Preload("Items.Item").First(&cart, cart.ID)
	c.JSON(http.StatusCreated, cart)
}

func getCarts(c *gin.Context) {
	userID := c.GetUint("user_id")
	var carts []Cart
	db.Where("user_id = ?", userID).Preload("Items.Item").Find(&carts)
	c.JSON(http.StatusOK, carts)
}

// Order handlers
func createOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify cart belongs to user
	var cart Cart
	if err := db.Where("id = ? AND user_id = ? AND status = ?", req.CartID, userID, "active").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found or already ordered"})
		return
	}

	// Create order
	order := Order{
		CartID: cart.ID,
		UserID: userID,
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Mark cart as ordered
	db.Model(&cart).Update("status", "ordered")

	// Clear user's cart_id
	db.Model(&User{}).Where("id = ?", userID).Update("cart_id", nil)

	// Return order with cart details
	db.Preload("Cart.Items.Item").First(&order, order.ID)
	c.JSON(http.StatusCreated, order)
}

func getOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	var orders []Order
	db.Where("user_id = ?", userID).Preload("Cart.Items.Item").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func seedData() {
	// Check if items already exist
	var count int64
	db.Model(&Item{}).Count(&count)
	if count > 0 {
		return // Data already seeded
	}

	// Create sample items
	items := []Item{
		{Name: "Laptop", Status: "active"},
		{Name: "Mouse", Status: "active"},
		{Name: "Keyboard", Status: "active"},
		{Name: "Monitor", Status: "active"},
		{Name: "Headphones", Status: "active"},
		{Name: "Webcam", Status: "active"},
		{Name: "Smartphone", Status: "active"},
		{Name: "Tablet", Status: "active"},
	}

	for _, item := range items {
		db.Create(&item)
	}

	// Create a test user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	testUser := User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	db.Create(&testUser)

	log.Println("Sample data seeded successfully")
}