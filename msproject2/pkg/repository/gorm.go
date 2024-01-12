package repository

import (
	"github.com/msproject2/internal/models"
	interfaces "github.com/msproject2/pkg"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepository {
	return &Repo{db}
}

func (repo *Repo) Create(user models.User) (string, error){
	err := repo.db.Create(&user).Error
	if err != nil {
		return "", err
	}
	return "User created successfully", nil
}

func (r *Repo) Get(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id=?", id).First(&user).Error
	return user, err
}
func (r *Repo) Update(user models.User) error {
	var dbuser models.User
	if err := r.db.Where("id=?", user.ID).First(&dbuser).Error; err != nil {
		return err
	}
	dbuser.Name = user.Name
	err := r.db.Save(dbuser).Error
	return err

}
func (r *Repo) Delete(id string) error {
	err := r.db.Where("id=?", id).Delete(&models.User{}).Error
	return err
}
func (r *Repo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email=?", email).First(&user).Error
	return user, err
}
