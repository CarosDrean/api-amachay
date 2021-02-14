package controllers

import "github.com/CarosDrean/api-amachay/models"

type CategoryMockOK struct {}

func (cm CategoryMockOK) Create(item *models.Category) (int64, error) {
	return 1, nil
}

func (cm CategoryMockOK) Update(ID string, item *models.Category) (int64, error) {
	return 1, nil
}

func (cm CategoryMockOK) Delete(ID string) (int64, error) {
	return 1, nil
}

func (cm CategoryMockOK) GetByID(ID string) (models.Category, error) {
	return models.Category{}, nil
}

func (cm CategoryMockOK) GetAll() ([]models.Category, error){
	return []models.Category{}, nil
}
