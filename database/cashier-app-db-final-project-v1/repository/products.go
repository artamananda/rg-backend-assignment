package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	result := p.db.Create(&product)
	return result.Error // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	results := []model.Product{}
	result := p.db.Raw("SELECT * FROM products WHERE deleted_at is NULL").Take(&results)
	if result.Error != nil {
		return []model.Product{}, nil
	}
	return results, result.Error // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	result := p.db.Delete(&model.Product{}, id)
	return result.Error // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	result := p.db.Model(&model.Product{}).Where("id = ?", id).Updates(product)
	return result.Error // TODO: replace this
}
