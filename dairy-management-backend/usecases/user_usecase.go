package usecases

import (
	"dairy-management-backend/entities"
	repositories "dairy-management-backend/repositories"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo *repositories.UserRepository
}

func NewUserUseCase(repo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: repo}
}

// HashPassword securely hashes the user's password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Create User (Admin only)
func (uc *UserUseCase) CreateUser(user *entities.User) error {
	// Hash password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	// Save user to repository
	return uc.UserRepo.CreateUser(user)
}

// Get All Users (Admin only)
func (uc *UserUseCase) GetAllUsers() ([]entities.User, error) {
	return uc.UserRepo.GetAllUsers()
}

// Get User by Email
func (uc *UserUseCase) GetUserByEmail(email string) (*entities.User, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update User
func (uc *UserUseCase) UpdateUser(id string, user *entities.User) error {
	return uc.UserRepo.UpdateUser(id, user)
}

// Delete User
func (uc *UserUseCase) DeleteUser(id string) error {
	return uc.UserRepo.DeleteUser(id)
}

// Authenticate User (Login)
func (uc *UserUseCase) AuthenticateUser(email, password string) (string, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Retrieve the JWT secret from environment variables
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret" // Fallback if not set
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
