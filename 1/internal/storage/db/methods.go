package db

import (
	"context"

	"github.com/1ef7yy/devops-course/internal/models"
)

func (pg *Postgres) GetUserByName(name string) (models.User, error) {

	data, err := pg.Database.Query(context.Background(), "SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		pg.Logger.Error("Unable to get user: " + err.Error())
		return models.User{}, err
	}

	if data.Next() {
		var user models.User
		err = data.Scan(&user.Name, &user.Age)
		if err != nil {
			pg.Logger.Error("Unable to scan user: " + err.Error())
			return models.User{}, err
		}
		return user, nil
	}
	return models.User{}, nil
}

func (pg *Postgres) InsertUser(user models.User) error {

	_, err := pg.Database.Exec(context.Background(), "INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		pg.Logger.Error("Unable to insert user: " + err.Error())
		return err
	}

	return nil
}
