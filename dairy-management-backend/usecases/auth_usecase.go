package usecases

import (
	"dairy-management-backend/entities"
	repositories "dairy-management-backend/repositories"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepo *repositories.UserRepository
}

func NewAuthUseCase(repo *repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepo: repo}
}

// AuthenticateUser validates user credentials and returns a JWT token
func (uc *AuthUseCase) AuthenticateUser(email, password string) (string, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	fmt.Println("Stored Hash:", user.Password)
	fmt.Println("Entered Password:", password)

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password mismatch!")
		return "", errors.New("invalid email or password")
	}

	fmt.Println("Password match successful!")

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("default-secret")
	}

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (uc *AuthUseCase) CreateUser(user *entities.User) error {
	// Hash password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword
	fmt.Println("Saving Hashed Password:", user.Password)
	// Save user to repository
	return uc.UserRepo.CreateUser(user)
}
