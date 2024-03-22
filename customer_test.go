package customer

import (
	"testing"
)

func TestNewCustomer(t *testing.T) {
    name := "John Snow"
    age := 20
    customer, err := newCustomer(name, age)

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if customer.name != name {
        t.Errorf("Expected customer name to be %v, got %v", name, customer.name)
    }

    if customer.age != age {
        t.Errorf("Expected customer age to be %v, got %v", age, customer.age)
    }

    if customer.id.value.String() == "" {
        t.Error("Expected customer id to be set")
    }
}

func TestCustomerNameGuard(t *testing.T) {
    _, err := newCustomer("", 20)

    if err == nil {
        t.Error("Expected error for empty name, got nil")
    }
}

func TestCustomerNameLessThan5Guard(t *testing.T) {
    _, err := newCustomer("John", 20)

    if err == nil {
        t.Error("Expected error for short name, got nil")
    }
}

func TestCustomerAgeGuard(t *testing.T) {
    _, err := newCustomer("John Snow", 17)

    if err == nil {
        t.Error("Expected error for age less than 18, got nil")
    }
}