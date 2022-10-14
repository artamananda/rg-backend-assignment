package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"fmt"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	ent, _ := s.database.Load()
	e, err := s.database.GetProductByname(productName)
	data := entity.CartItem{}

	if quantity <= 0 {
		err = fmt.Errorf("invalid quantity")
		return err
	}
	data.ProductName = productName
	data.Price = e.Price
	data.Quantity = quantity

	ent = append(ent, data)

	_ = s.database.Save(ent)

	return err // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	res := []entity.CartItem{}
	ent, _ := s.database.Load()
	var err error
	b := false

	for i := 0; i < len(ent); i++ {
		if productName == ent[i].ProductName {
			b = true
			break
		}
	}

	if !b {
		err = fmt.Errorf("product not found")
		return err
	}

	for i := 0; i < len(ent); i++ {
		if ent[i].ProductName != productName {
			res = append(res, ent[i])
		}
	}
	_ = s.database.Save(res)

	return err // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	_, err := s.database.Load()
	newEnt := []entity.CartItem{}
	_ = s.database.Save(newEnt)
	return err // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	res := s.database.GetProductData()

	return res, nil // TODO: replace this
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {
	cart, err := s.database.Load()
	res := entity.PaymentInformation{}
	res.ListProduct = cart
	res.MoneyPaid = money
	for i := 0; i < len(cart); i++ {
		res.TotalPrice += cart[i].Price * cart[i].Quantity
	}
	fmt.Println(res.MoneyPaid)
	fmt.Println(res.TotalPrice)
	if res.MoneyPaid < res.TotalPrice {
		err = fmt.Errorf("money is not enough")
		res = entity.PaymentInformation{}
		return res, err
	}

	res.Change = res.MoneyPaid - res.TotalPrice
	s.ResetCart()
	return res, err // TODO: replace this
}
