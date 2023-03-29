package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"section24/model"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllUsers(t *testing.T) {
	// setup
	db, err := gorm.Open(mysql.Open("db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to setup database: %s", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %s", err)
	}

	// create test data
	users := []model.User{
		{Name: "asmii", Email: "asmiii@gmail.com"},
		{Name: "ravitaa", Email: "rvtas@gmail.com"},
	}
	err = db.Create(&users).Error
	if err != nil {
		t.Fatalf("failed to create test data: %s", err)
	}

	// create request context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call the handler function
	err = GetAllUsers(db)(c)
	if err != nil {
		t.Fatalf("failed to get all users: %s", err)
	}

	// check the response
	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", rec.Code, http.StatusOK)
	}
	var respData struct {
		Data []model.User `json:"data"`
	}
	err = json.NewDecoder(rec.Body).Decode(&respData)
	if err != nil {
		t.Fatalf("failed to decode response body: %s", err)
	}
	if len(respData.Data) != len(users) {
		t.Errorf("unexpected number of users: got %d, want %d", len(respData.Data), len(users))
	}
	for i, user := range users {
		if respData.Data[i].Username != user.Name {
			t.Errorf("unexpected username: got %s, want %s", respData.Data[i].Name, user.Name)
		}
		if respData.Data[i].Email != user.Email {
			t.Errorf("unexpected email: got %s, want %s", respData.Data[i].Email, user.Email)
		}
	}
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := {"Name":"asmiii", "Email":"rvtas@gmail.com"}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Buat mock DB
db, err := gorm.Open(mysql.Open("db"), &gorm.Config{})
if err != nil {
	t.Fatal(err)
}
defer db.Close()

// Migrasi tabel User
err = db.AutoMigrate(&model.User{})
if err != nil {
	t.Fatal(err)
}

// Panggil CreateUser
err = CreateUser(db)(c)

// Assert
if err != nil {
	t.Fatal(err)
}
if rec.Code != http.StatusOK {
	t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
}

var response struct {
	Data model.User `json:"data"`
}
err = json.NewDecoder(rec.Body).Decode(&response)
if err != nil {
	t.Fatal(err)
}
if response.Data.Username != "johndoe" {
	t.Errorf("Expected username %s but got %s", "johndoe", response.Data.Name)
}
if response.Data.Email != "johndoe@example.com" {
	t.Errorf("Expected email %s but got %s", "johndoe@example.com", response.Data.Email)
}
}
