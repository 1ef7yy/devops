package domain

import (
	"github.com/1ef7yy/devops-course/internal/models"
)

func (d *domain) GetUserByName(name string) (models.User, error) {

	data, err := d.pg.GetUserByName(name)
	if err != nil {
		d.Logger.Error("Unable to get user: " + err.Error())
		return models.User{}, err
	}

	return data, nil

}

func (d *domain) InsertUser(user models.User) error {

	err := d.pg.InsertUser(user)
	if err != nil {
		d.Logger.Error("Unable to insert user: " + err.Error())
		return err
	}

	return nil
}
