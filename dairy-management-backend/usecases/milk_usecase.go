// usecases/milk_usecase.go - Business logic for milk collection
package usecases

import (
	"dairy-management-backend/entities"
	repositories "dairy-management-backend/repositories"
)

type MilkUseCase struct {
	MilkRepo *repositories.MilkRepository
}

func NewMilkUseCase(repo *repositories.MilkRepository) *MilkUseCase {
	return &MilkUseCase{MilkRepo: repo}
}

func (uc *MilkUseCase) AddMilkRecord(milk *entities.MilkCollection) error {
	return uc.MilkRepo.CreateMilkRecord(milk)
}

func (uc *MilkUseCase) GetMilkRecords() ([]entities.MilkCollection, error) {
	return uc.MilkRepo.GetAllMilkRecords()
}

func (uc *MilkUseCase) GetMilkRecordByID(id uint) (*entities.MilkCollection, error) {
	return uc.MilkRepo.GetMilkRecordByID(id)
}

func (uc *MilkUseCase) UpdateMilkRecord(milk *entities.MilkCollection) error {
	return uc.MilkRepo.UpdateMilkRecord(milk)
}

func (uc *MilkUseCase) RemoveMilkRecord(id uint) error {
	return uc.MilkRepo.DeleteMilkRecord(id)
}
