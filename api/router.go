package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	IdParam = "id"
)

type Route struct {
	Name    string
	Path    string
	Action  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var Routes []Route = []Route{
	{Name: "create task", Path: "/createTask", Action: "POST", Handler: CreateTask},
	{Name: "read task", Path: "/readTask", Action: "GET", Handler: ReadTask},
	{Name: "delete task", Path: "/deleteTask", Action: "DELETE", Handler: DeleteTask},
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Action)
	}
	return r
}
