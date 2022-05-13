package repository

import (
	"github.com/jinzhu/gorm"
	"myproject/Model"
	"myproject/common"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*Model.Category, error) {
	cate := Model.Category{Name: name}
	if err := c.DB.Create(&cate).Error; err != nil {
		return nil, err
	}
	return &cate, nil
}

func (c CategoryRepository) Update(cate Model.Category, name string) (*Model.Category, error) {
	if err := c.DB.Model(&cate).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &cate, nil
}
func (c CategoryRepository) SelectByID(id int) (*Model.Category, error) {
	var cate Model.Category
	if err := c.DB.Where("ID=?", id).First(&cate).Error; err != nil {
		return nil, err
	}
	return &cate, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(Model.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
