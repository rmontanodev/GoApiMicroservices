package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/handler"
	"user-service/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userHandler := handler.NewUserHandler(mockRepo)

	users := []model.User{
		{ID: 1, Name: "User 1", Email: "user1@example.com"},
		{ID: 2, Name: "User 2", Email: "user2@example.com"},
	}

	mockRepo.On("GetAllUsers").Return(users, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(userHandler.GetAllUsers)
	httpHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var returnedUsers []model.User
	err = json.Unmarshal(rr.Body.Bytes(), &returnedUsers)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, users, returnedUsers)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userHandler := handler.NewUserHandler(mockRepo)

	user := model.User{Name: "User 1", Email: "user1@example.com"}

	mockRepo.On("GetUserByID", 1).Return(user, nil)

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(userHandler.GetUserByID)
	httpHandler.ServeHTTP(rr, req)
	t.Logf("Response body: %s", rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code)
	var returnedUser model.User
	err = json.Unmarshal(rr.Body.Bytes(), &returnedUser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user, returnedUser)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userHandler := handler.NewUserHandler(mockRepo)

	newUser := model.User{Name: "User 1", Email: "user1@example.com"}
	createdUser := model.User{ID: 1, Name: "User 1", Email: "user1@example.com"}

	mockRepo.On("CreateUser", newUser).Return(createdUser, nil)

	body, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(userHandler.CreateUser)
	httpHandler.ServeHTTP(rr, req)
	t.Logf("Response body: %s", rr.Body.String())

	assert.Equal(t, http.StatusCreated, rr.Code)
	var returnedUser model.User
	err = json.Unmarshal(rr.Body.Bytes(), &returnedUser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, createdUser, returnedUser)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userHandler := handler.NewUserHandler(mockRepo)

	updatedUser := model.User{ID: 1, Name: "User 1 Updated", Email: "user1updated@example.com"}

	mockRepo.On("UpdateUser", updatedUser).Return(nil)

	body, err := json.Marshal(updatedUser)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users/update/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(userHandler.UpdateUser)
	httpHandler.ServeHTTP(rr, req)
	t.Logf("Response body: %s", rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code)
	var returnedUser model.User
	err = json.Unmarshal(rr.Body.Bytes(), &returnedUser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, updatedUser, returnedUser)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userHandler := handler.NewUserHandler(mockRepo)

	mockRepo.On("DeleteUser", 1).Return(nil)

	req, err := http.NewRequest("DELETE", "/users/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(userHandler.DeleteUser)
	httpHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockRepo.AssertExpectations(t)
}
