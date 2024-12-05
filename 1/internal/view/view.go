package view

import (
	"net/http"

	"github.com/1ef7yy/devops-course/pkg/logger"

	"github.com/1ef7yy/devops-course/internal/domain"
)

type view struct {
	Logger logger.Logger
	domain domain.Domain
}

type View interface {
	GetUserByName(w http.ResponseWriter, r *http.Request)
	InsertUser(w http.ResponseWriter, r *http.Request)
}

func NewView(logger logger.Logger, domain domain.Domain) View {
	return &view{
		Logger: logger,
		domain: domain,
	}
}
