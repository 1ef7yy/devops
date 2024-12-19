package routes

import (
	"net/http"

	"github.com/1ef7yy/devops-course/internal/view"

	v1 "github.com/1ef7yy/devops-course/internal/routes/v1"
)

func InitRouter(view view.View) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := v1.NewRouter(view)

	mux.Handle("/api/", v1.Endpoints())

	return mux
}
