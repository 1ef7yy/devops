package domain

import (
	"context"
	"os"

	"github.com/1ef7yy/devops-course/internal/models"
	"github.com/1ef7yy/devops-course/internal/storage/db"
	"github.com/1ef7yy/devops-course/pkg/logger"
)

type domain struct {
	Logger logger.Logger
	pg     db.Postgres
}

type Domain interface {
	GetUserByName(id string) (models.User, error)
	InsertUser(data models.User) error
}

func NewDomain(logger logger.Logger) Domain {
	pg, err := db.NewPostgres(context.Background(), os.Getenv("POSTGRES_CONN"), logger)
	if err != nil {
		logger.Error("Unable to create connection to database: " + err.Error())
		return nil
	}
	return &domain{
		Logger: logger,
		pg:     *pg,
	}

}
