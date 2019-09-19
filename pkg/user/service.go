package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UserService struct {
	dao Dao
	l   *zap.Logger
}

func NewUserService(dao Dao, l *zap.Logger) *UserService {
	return &UserService{
		dao: dao,
		l:   l,
	}
}

func (us *UserService) GetRoutes(router *mux.Router) *mux.Router {
	r := router.PathPrefix("/users").Subrouter()

	r.HandleFunc("/", us.ListUsers).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", us.GetUser).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", us.DeleteUser).Methods("DELETE")
	r.HandleFunc("/", us.CreateUser).Methods("PUT")
	return r
}

func (us *UserService) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := us.dao.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (us *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(sid, 10, 64)

	if err != nil {
		fmt.Fprintln(w, "error: id is not a number")
		return
	}

	u, err := us.dao.GetUser(id)
	if err != nil {
		fmt.Fprintf(w, "error: %s\r\n", err.Error())
		return
	}
	json.NewEncoder(w).Encode(u)
	return
}

func (us *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		fmt.Fprintf(w, "error: pad payload: %s\r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := us.dao.CreateUser(&u)
	if err != nil {
		fmt.Fprintf(w, "error: unable to create user: %s\r\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "created used with id %d\r\n", id)
}

func (us *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
