package repository

import (
	"dairy-management-backend/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
func (repo *UserRepository) DeleteUser(id string) error {
	// Example: Deleting a user by ID in the database
	result := repo.DB.Delete(&entities.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) UpdateUser(id string, user *entities.User) error {
	return r.DB.Model(&entities.User{}).Where("id = ?", id).Updates(user).Error
}
