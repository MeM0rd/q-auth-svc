package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MeM0rd/q-auth-svc/internal/entities"
	"github.com/MeM0rd/q-auth-svc/pkg/client/postgres"
	"github.com/MeM0rd/q-auth-svc/pkg/logger"
	authPbService "github.com/MeM0rd/q-auth-svc/pkg/pb/auth"
	"github.com/MeM0rd/q-auth-svc/pkg/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Server struct {
	authPbService.UnimplementedAuthPbServiceServer
	logger logger.Logger
}

func (s *Server) Register(ctx context.Context, in *authPbService.RegisterRequest) (*authPbService.RegisterResponse, error) {
	var user entities.User
	var exists bool

	rd := entities.RegisterData{
		Email:    in.GetEmail(),
		Surname:  in.GetSurname(),
		Name:     in.GetName(),
		Password: in.GetPassword(),
	}

	q := `SELECT EXISTS(SELECT 1 FROM users WHERE email LIKE $1);`

	err := postgres.DB.QueryRow(q, rd.Email).Scan(&exists)
	if err != nil {
		return &authPbService.RegisterResponse{
			Status: "error",
			Msg:    nil,
			Err:    fmt.Sprintf("%v", err),
		}, err
	}

	if exists == true {
		return &authPbService.RegisterResponse{
			Status: "error",
			Msg:    []byte("user already exists"),
			Err:    "",
		}, nil
	}

	password, err := bcrypt.GenerateFromPassword([]byte(rd.Password), 14)
	if err != nil {
		log.Printf("error bcrypt password %s: %v", user.Password, err)
		return &authPbService.RegisterResponse{
			Status: "error",
			Msg:    nil,
			Err:    fmt.Sprintf("error bcrypt password %s: %v", user.Password, err),
		}, err
	}

	rd.Password = string(password)

	q = `INSERT INTO users(email, surname, name, password) VALUES ($1, $2, $3, $4) RETURNING id, surname, name, email`

	err = postgres.DB.QueryRow(q, rd.Email, rd.Surname, rd.Name, rd.Password).Scan(&user.Id, &user.Surname, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return &authPbService.RegisterResponse{
			Status: "error",
			Msg:    nil,
			Err:    fmt.Sprintf("Error creating user: %v", err),
		}, err
	}

	j, err := json.Marshal(user)
	if err != nil {
		log.Printf("Marshaling users error: %v", err)
		return &authPbService.RegisterResponse{
			Status: "error",
			Msg:    nil,
			Err:    "",
		}, err
	}

	return &authPbService.RegisterResponse{
		Status: "success",
		Msg:    j,
		Err:    "",
	}, nil
}

func (s *Server) Login(ctx context.Context, in *authPbService.LoginRequest) (*authPbService.LoginResponse, error) {
	var user entities.User

	ld := entities.LoginData{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	q := `SELECT id, email, password FROM users WHERE email = $1`

	err := postgres.DB.QueryRow(q, ld.Email).Scan(&user.Id, &user.Email, &user.Password)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return &authPbService.LoginResponse{
			Msg: "",
			Err: fmt.Sprintf("Error queryRow creating user: %v", err),
		}, err
	case err != nil:
		return &authPbService.LoginResponse{
			Msg: "",
			Err: fmt.Sprintf("Error queryRow creating user: %v", err),
		}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ld.Password))
	if err != nil {
		return &authPbService.LoginResponse{
			Msg: "",
			Err: fmt.Sprintf("error compare hash an password: %v", err),
		}, err
	}

	session, err := sessions.CreateSession(&user.Id)
	if err != nil {
		return &authPbService.LoginResponse{
			Msg: "",
			Err: fmt.Sprintf("error generate session: %v", err),
		}, err
	}

	return &authPbService.LoginResponse{
		Token:    session.Token,
		Lifetime: session.ExpiredAt.Format(time.DateTime),
		Msg:      session.Name,
		Err:      "",
	}, nil
}

func (s *Server) Logout(ctx context.Context, in *authPbService.LogoutRequest) (*authPbService.LogoutResponse, error) {
	err := sessions.DeleteSession(in.GetToken())
	if err != nil {
		return &authPbService.LogoutResponse{
			Status: "error",
			Err:    fmt.Sprintf("error deleting session in psql: %v", err),
		}, err
	}

	return &authPbService.LogoutResponse{
		Status: "success",
		Err:    "",
	}, nil
}
