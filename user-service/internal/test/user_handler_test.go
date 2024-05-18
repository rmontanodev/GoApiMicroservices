package test

import (
	"auction/user-service/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Mock de UserService
type MockUserService struct{}

func (m *MockUserService) CreateUser(name, email string) error {
	return nil
}

func (m *MockUserService) GetUserByID(id uint) (*model.User, error) {
	return &model.User{Model: gorm.Model{ID: id}, Name: "Test User", Email: "test@example.com"}, nil
}

func TestCreateUser(t *testing.T) {
	r := gin.Default()
	userService := &MockUserService{}
	userHandler := handler.NewUserHandler(userService)
	r.POST("/users", userHandler.CreateUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUser(t *testing.T) {
	r := gin.Default()
	userService := &MockUserService{}
	userHandler := handler.NewUserHandler(userService)
	r.GET("/users/:id", userHandler.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test User")
}
