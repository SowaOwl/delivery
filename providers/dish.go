package providers

import (
	"delivery/models"
	"errors"
	"net/http"

	"encoding/json"
	"io/ioutil"
	"os"
)

type DishProvider interface {
	GetDishes() ([]models.Dish, error)
}

type FileDishProvider struct {
	FilePath string
}

type HttpDishProvider struct {
	Url string
}

func NewFileDishProvider(filePath string) *FileDishProvider {
	return &FileDishProvider{FilePath: filePath}
}

func NewHttpDishProvider(url string) *HttpDishProvider {
	return &HttpDishProvider{Url: url}
}

func (p *FileDishProvider) GetDishes() ([]models.Dish, error) {
	file, err := os.Open(p.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var dishes []models.Dish
	err = json.Unmarshal(data, &dishes)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (p *HttpDishProvider) GetDishes() ([]models.Dish, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var dishes []models.Dish
	err = json.NewDecoder(resp.Body).Decode(&dishes)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
