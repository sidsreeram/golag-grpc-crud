package Usecase

import (
	"errors"
	"fmt"

	"github.com/msproject2/internal/models"
	interfaces "github.com/msproject2/pkg"
	"gorm.io/gorm"
)
type Usecase struct {
	repo interfaces.UserRepository
}
func NewUsecase(repo interfaces.UserRepository) interfaces.UseUsecase{
	return &Usecase{repo}
}
func (u *Usecase) Create(user models.User) (string, error) {
	// check if valid email is supplied
	if _, err := u.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound){
	  return "", errors.New("the email is already associated with another user")
	}
  
	// email does not exist so, now proceed
	_, err := u.repo.Create(user)
	if err != nil {
		return "", err
	}
	return "User created successfully", nil
}

  
func (u *Usecase) Get(id string) (models.User, error){
	var user models.User
	var err error
  
	if user, err = u.repo.Get(id); err != nil{
	  if errors.Is(err, gorm.ErrRecordNotFound){
		return models.User{}, errors.New("no such user with the id supplied")
	  }
	  return models.User{}, err
	}
  
	return user, nil
  }
  
  
  func (u *Usecase) Update(updateUser models.User) (error){
	var user models.User
	var err error
	if user, err = u.Get(string(rune(updateUser.ID))); err != nil {
	  return fmt.Errorf("Can't convert userid into string :%w",err)
	}

	if user.Email != updateUser.Email {
	  return errors.New("email cannot be changed")
	}
	
	err = u.repo.Update(updateUser)
	if err != nil {
	  return fmt.Errorf("Can't update the email : %w",err)
	}
  
	return nil
  }
  
  func (u *Usecase) Delete(id string) error{
	var err error
	if _, err = u.Get(id); err != nil {
	  return fmt.Errorf("Can't find User with this Id :%w",err)
	}
  
	err = u.repo.Delete(id)
	if err != nil {
	  return fmt.Errorf("Can't Delete the user :%w",err)
	}
  
	return nil
  }