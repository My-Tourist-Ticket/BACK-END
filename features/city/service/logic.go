package service

import (
	"fmt"
	"mime/multipart"
	"my-tourist-ticket/features/city"
)

type cityService struct {
	cityData city.CityDataInterface
}

func NewCity(repo city.CityDataInterface) city.CityServiceInterface {
	return &cityService{
		cityData: repo,
	}
}

// GetUserRoleById implements admin.AdminServiceInterface.
func (service *cityService) GetUserRoleById(userId int) (string, error) {
	return service.cityData.GetUserRoleById(userId)
}

// Insert implements city.CityDataInterface.
func (service *cityService) Create(input city.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	err := service.cityData.Insert(input, image, thumbnail)
	if err != nil {
		return fmt.Errorf("error creating city: %w", err)
	}

	return nil
}

// Update implements city.CityServiceInterface.
func (service *cityService) Update(cityId int, input city.Core, image *multipart.FileHeader, thumbnail *multipart.FileHeader) error {
	err := service.cityData.Update(cityId, input, image, thumbnail)
	if err != nil {
		return fmt.Errorf("error update city: %w", err)
	}

	return nil
}

// Delete implements city.CityDataInterface.
func (*cityService) Delete(cityId int) error {
	panic("unimplemented")
}

// SelectCityById implements city.CityServiceInterface.
func (service *cityService) SelectCityById(cityId int) (city.Core, error) {
	data, err := service.cityData.SelectCityById(cityId)
	if err != nil {
		return city.Core{}, err
	}

	return data, nil
}
