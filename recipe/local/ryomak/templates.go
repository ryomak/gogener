package ryomak

var modTempl = `
module [[.ModName]]

go 1.12
`

var mainTmpl = `
package main

import (
	"[[.ModName]]/src/interface/router"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	if os.Getenv("APP_ENV") == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("start server port:", os.Getenv("APP_PORT"), "\nenv:", os.Getenv("APP_ENV"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), router.New())
}

`

var envTmpl = `APP_NAME=[[.AppName]]
APP_PORT=8081
APP_REQUEST_ID=app
APP_ENV=debug
`

var usecaseTmpl = `
package usecase

import (
	"context"
	"[[.ModName]]/src/domain/user"
)

type IUserUseCase interface {
	GetUserByID(context.Context, uint) (*user.User, error)
}

type userUseCase struct {
	userRepo user.IUserRepository
}

func NewIUserUseCase(
	user user.IUserRepository,
) IUserUseCase {
	return &userUseCase{user}
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uint) (*user.User, error) {
	fUser, err := u.userRepo.FindOneByID(ctx, id)
	if err != nil {
		return nil, err
	}
	fUser.Name = user.ToSnakeName(fUser)
	return fUser, nil
}
`

var domainEntityUserTmpl = `
package user

type User struct {
	ID   uint   ` + "`" + `json:"id"` + "`\n" +
	`Name string ` + "`" + `json:"name"` + "`\n" +
	`}`

var domainUserRepositoryTmpl = `
package user

import "context"

type IUserRepository interface {
	FindOneByID(context.Context, uint) (*User, error)
}
`

var domainUserServiceTmpl = `
package user

import (
	"github.com/iancoleman/strcase"
)

func ToSnakeName(user *User) string {
	return strcase.ToSnake(user.Name)
}
`

var infraUserRepositoryTmpl = `
package repository

import (
	"context"
	"errors"

	"[[.ModName]]/src/domain/user"
)

func NewIUserRepository() user.IUserRepository {
	// mock
	users := map[uint]*user.User{
		1: &user.User{
			ID:   1,
			Name: "ryomak",
		},
		2: &user.User{
			ID:   2,
			Name: "test user",
		},
		3: &user.User{
			ID:   3,
			Name: "test user3",
		},
	}
	return &userRepository{users}
}

type userRepository struct {
	users map[uint]*user.User
}

func (r *userRepository) FindOneByID(ctx context.Context, id uint) (*user.User, error) {
	if val, ok := r.users[id]; !ok {
		return nil, errors.New("not found")
	} else {
		return val, nil
	}
}
`

var handlerTmpl = `
package handler

import (
	"net/http"
	"os"
)

func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello " + os.Getenv("APP_NAME")))
	}
}
`

var userHandlerTmpl = `
package handler

import (
	"encoding/json"
	"strconv"

	"net/http"

	"[[.ModName]]/src/application/usecase"
	"[[.ModName]]/src/infrastructure/repository"
	"[[.ModName]]/src/internal/logger"

	"github.com/gorilla/mux"
)

func GetUserHandler() func(w http.ResponseWriter, r *http.Request) {
	uUsecase := usecase.NewIUserUseCase(
		repository.NewIUserRepository(),
	)
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerWithCtx(r.Context())

		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "id not number"})
			return
		}
		logger.Info("Start usecase:", userID)
		user, err := uUsecase.GetUserByID(r.Context(), uint(userID))
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "notfound"})
			return
		}
		logger.Info("End usecase:", userID)
		json.NewEncoder(w).Encode(user)
	}
}
`

var routerTmpl = `
package router

import (
	"[[.ModName]]/src/interface/handler"
	"[[.ModName]]/src/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	DeleteMethod = "DELETE"
)

type Route struct {
	Path    string
	Methods []string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func routes() []*Route {
	return []*Route{
		{
			Path:    "/",
			Methods: []string{GetMethod, PostMethod},
			Handler: handler.HelloHandler(),
		},
		{
			Path:    "/users/{id:[0-9]+}",
			Methods: []string{GetMethod},
			Handler: handler.GetUserHandler(),
		},
	}
}

func New() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ReqIdMiddleware)
	for _, v := range routes() {
		r.HandleFunc(v.Path, v.Handler).Methods(v.Methods...)
	}
	return r
}

`

var loggerTmpl = `
package logger

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func GetLoggerWithCtx(ctx context.Context) *log.Entry {
	return log.WithField(os.Getenv("APP_REQUEST_ID"), ctx.Value(os.Getenv("APP_REQUEST_ID")).(string))
}

`

var middlewareTmpl = `
package middleware

import (
	"context"
	"[[.ModName]]/src/internal/logger"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func ReqIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := ""
		if id = r.Header.Get("UUID"); id == "" {
			id = uuid.New().String()
		}
		ctx := r.Context()
		r = r.WithContext(context.WithValue(ctx, os.Getenv("APP_REQUEST_ID"), id))

		log := logger.GetLoggerWithCtx(r.Context())
		log.Debugf("Incomming request %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Debugf("Finished handling request")
	})
}

`

var makeFileTmpl = `
run:
	go run src/main.go

`
