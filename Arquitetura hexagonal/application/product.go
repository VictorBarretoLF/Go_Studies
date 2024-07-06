package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

// IsValid checks if the product is valid based on its price and status
func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled of disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater of equal zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Enable sets the product status to enabled
func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("the price must be greater than zero to enable the product")
}

// Disable sets the product status to disabled
func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("the price must be zero in order to have the product disabled")
}

// GetId returns the ID of the product
func (p *Product) GetId() string {
	return p.ID
}

// GetName returns the name of the product
func (p *Product) GetName() string {
	return p.Name
}

// GetStatus returns the status of the product
func (p *Product) GetStatus() string {
	return p.Status
}

// GetPrice returns the price of the product
func (p *Product) GetPrice() float64 {
	return p.Price
}
