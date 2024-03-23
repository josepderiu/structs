# Go Structs and Validations Learning Project

This project is a learning resource for understanding structs and validations in Go. It provides a basic implementation of a customer management system, demonstrating the use of structs for data organization and the `github.com/go-playground/validator/v10` package for data validation.

## Structs

The project defines four main structs: `customerId`, `Customer`, `customerName`, and `customerAge`.

`customerId` is a struct that contains a single field, `value`, which is a globally unique identifier (GUID). This struct is not directly accessible outside the package.

`Customer` is a struct that represents a customer in the system. It contains three fields:

- `id`: a `customerId` that represents the unique identifier of the customer.
- `name`: a `customerName` that represents the name of the customer.
- `age`: a `customerAge` that represents the age of the customer.

`customerName` and `customerAge` are value objects that encapsulate specific validation rules for the name and age of a customer, respectively.

## Validations

The project uses the `github.com/go-playground/validator/v10` package to validate `Customer` instances. The validation rules are as follows:

- `id` must be present.
- `name` cannot be empty and must be at least 5 characters long.
- `age` must be greater than or equal to 18.

These rules are defined using struct tags, demonstrating how to use tags for validation purposes in Go.

## Factory Function

The project provides a public factory function `NewCustomer` that creates a new `Customer` instance. The factory function validates the `Customer` instance before returning it. If the instance fails validation, the factory function returns an error. This demonstrates how to encapsulate validation logic within a function.

## How to Use

To use this project, you can import the `Customer` struct and the `NewCustomer` function into your own code. Here's an example of how you can do this:

```go
package main

import (
    "fmt"
    "github.com/josepderiu/structs"
)

func main() {
    customer, err := customer.NewCustomer("John Doe", 20)
    if err != nil {
        fmt.Println("Error creating customer:", err)
        return
    }

    fmt.Println("Created customer:", customer)
}