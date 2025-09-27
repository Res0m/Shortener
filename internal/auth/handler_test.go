package auth_test

import (
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/auth"
	"GolangAdvanced/internal/user"
	"GolangAdvanced/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a2@a.ru",
		Password: "1",
		NickName: "Вася",
	})
	reader := bytes.NewReader(data)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Register()(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got %d, expected %d", wr.Code, 201)
	}
}
func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a2@a.ru", "$2a$10$VcU7JMXreFTEuulfDWju8eWrnvpUpjJ/T2Yc1ZbG8gTYWfp8pZlCK")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})
	reader := bytes.NewReader(data)

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", wr.Code, 200)
	}
}
