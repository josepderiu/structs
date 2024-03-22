package customer

import (
	"github.com/beevik/guid"

	"github.com/go-playground/validator/v10"
)

type customerId struct {
	value guid.Guid
}

type customer struct {
	id customerId `validate:"required"`
	name string `validate:"required,min=5"`
	age int `validate:"gte=18"`
}

func newCustomer(name string, age int) (*customer, error) {

	c := &customer{
		id: customerId{
			value: *guid.New(),
		},
		name: name,
		age: age,
	}

	validate := validator.New(validator.WithPrivateFieldValidation())
	err := validate.Struct(c)

	if err != nil {
		return nil, err
	}

	return c, nil
}