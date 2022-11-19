package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	results := []model.JoinCart{}
	result := c.db.Model(&model.Cart{}).Select("carts.id, carts.product_id, products.name, carts.quantity, carts.total_price").Joins("left join products on products.id = carts.product_id").Scan(&results)
	if result.Error != nil {
		return []model.JoinCart{}, nil
	}
	return results, result.Error // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	results := model.Cart{}
	result := c.db.First(&results, "product_id = ?", product.ID)
	if result.Error != nil {
		results.ProductID = product.ID
		results.Quantity = 1
		results.TotalPrice = product.Price - (product.Discount / 100 * product.Price)
		c.db.Create(&results)
		product.Stock -= 1
		c.db.Model(&product).Where("id = ?", product.ID).Updates(&product)
		return nil
	}
	results.Quantity += 1
	results.TotalPrice += product.Price - (product.Discount / 100 * product.Price)
	c.db.Where("product_id = ?", results.ProductID).Updates(&results)
	product.Stock -= 1
	c.db.Model(&product).Where("id = ?", product.ID).Updates(&product)
	return result.Error // TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	resultCart := model.Cart{}
	resultProduct := model.Product{}
	result := c.db.First(&resultCart, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	result = c.db.Model(&model.Product{}).First(&resultProduct, "id = ?", productID)
	if result.Error != nil {
		return nil
	}
	c.db.Where("id = ?", id).Delete(&model.Cart{})
	resultProduct.Stock += 1
	c.db.Model(&model.Product{}).Where("id = ?", productID).Updates(&resultProduct)
	return result.Error // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	result := c.db.Model(&model.Cart{}).Where("id = ?", id).Updates(cart)
	return result.Error // TODO: replace this
}
