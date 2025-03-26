package usecases

import (
	"dairy-management-backend/entities"
	repositories "dairy-management-backend/repositories"
)

type CowUseCase struct {
	CowRepo *repositories.CowRepository
}

func NewCowUseCase(repo *repositories.CowRepository) *CowUseCase {
	return &CowUseCase{CowRepo: repo}
}

func (uc *CowUseCase) AddCow(cow *entities.Cow) error {
	return uc.CowRepo.CreateCow(cow)
}

func (uc *CowUseCase) GetCows() ([]entities.Cow, error) {
	return uc.CowRepo.GetAllCows()
}

func (uc *CowUseCase) GetCowByTagID(tagID string) (*entities.Cow, error) {
	return uc.CowRepo.GetCowByTagID(tagID) // Call repository method with tag_id
}

func (uc *CowUseCase) UpdateCow(cow *entities.Cow) error {
	return uc.CowRepo.UpdateCow(cow)
}

func (uc *CowUseCase) RemoveCow(id uint) error {
	return uc.CowRepo.DeleteCow(id)
}
