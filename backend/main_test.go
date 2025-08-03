package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Auto migrate schemas
	db.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	setupRoutes(r)
	return r
}

func TestCreateUser(t *testing.T) {
	setupTestDB()
	router := setupTestRouter()

	userReq := UserRequest{
		Username: "testuser",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(userReq)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData