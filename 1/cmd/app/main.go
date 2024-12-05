package main

import (
	"net/http"
	"os"

	"github.com/1ef7yy/devops-course/internal/domain"
	"github.com/1ef7yy/devops-course/internal/routes"
	"github.com/1ef7yy/devops-course/internal/view"
	"github.com/1ef7yy/devops-course/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	domain := domain.NewDomain(log)

	view := view.NewView(log, domain)

	log.Info("Initializing router...")

	mux := routes.InitRouter(view)

	log.Info("Server started on: " + os.Getenv("SERVER_ADDRESS"))
	if err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), mux); err != nil {
		log.Error("Error starting server: " + err.Error())
	}

}
