package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
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
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	cartItems = append(cartItems, entity.CartItem{
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
	})

	err = s.database.SaveCartItems(cartItems)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	// database := database.NewDatabase()
	productCheck := false
	_, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	newCartItems := []entity.CartItem{}

	for i, p := range cartItems {
		if p.ProductName != productName {
			newCartItems = append(newCartItems, cartItems[i])
		} else if p.ProductName == productName {
			productCheck = true
		}
	}

	if productCheck == false {
		return errors.New("product not found")
	}

	err = s.database.SaveCartItems(newCartItems)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	newCart := []entity.CartItem{}
	err := s.database.SaveCartItems(newCart)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	prodcts := s.database.GetProductData()

	return prodcts, nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cartItems, _ := s.database.GetCartItems()
	totalPrice := 0

	for _, item := range cartItems {
		totalPrice += item.Price * item.Quantity
	}
	change := money - totalPrice
	// fmt.Printf("kembalian = %d", change)
	details := entity.PaymentInformation{
		ProductList: cartItems,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      change,
	}

	if change < 0 {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	return details, s.ResetCart() // TODO: replace this
}
