package structs

import (
	"github.com/beevik/guid"

	"github.com/go-playground/validator/v10"
)

type customerId struct {
	value guid.Guid `validate:"required"`
}

type customerName struct {
	value string `validate:"required,min=5"`
}

type customerAge struct {
	value int `validate:"gte=18"`
}

type Customer struct {
	id   customerId
	name customerName
	age  customerAge
}

func NewCustomer(name string, age int) (*Customer, error) {
	c := &Customer{
		id: customerId{
			value: *guid.New(),
		},
		name: customerName{
			value: name,
		},
		age: customerAge{
			value: age,
		},
	}

	validate := validator.New(validator.WithPrivateFieldValidation())
	err := validate.Struct(c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
