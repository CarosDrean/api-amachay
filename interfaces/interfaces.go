package interfaces

import "github.com/CarosDrean/api-amachay/models"

type CategoryStorage interface {
	Create(item *models.Category) (int64, error)
	Update(ID string, item *models.Category) (int64, error)
	Delete(ID string) (int64, error)
	GetByID(ID string) (models.Category, error)
	GetAll() ([]models.Category, error)
}
