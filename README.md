# Go Structs and Validations Learning Project

This project is a learning resource for understanding structs and validations in Go. It provides a basic implementation of a customer management system, demonstrating the use of structs for data organization and the `github.com/go-playground/validator/v10` package for data validation.

## Structs

The project defines two main structs: `customerId` and `customer`.

`customerId` is a struct that contains a single field, `value`, which is a globally unique identifier (GUID).

`customer` is a struct that represents a customer in the system. It contains three fields:

- `id`: a `customerId` that represents the unique identifier of the customer.
- `name`: a string that represents the name of the customer.
- `age`: an integer that represents the age of the customer.

## Validations

The project uses the `github.com/go-playground/validator/v10` package to validate `customer` instances. The validation rules are as follows:

- `id` must be present.
- `name` cannot be empty.
- `age` must be greater than or equal to 18.

These rules are defined using struct tags, demonstrating how to use tags for validation purposes in Go.

## Factory Function

The project provides a private factory function `newCustomer` that creates a new `customer` instance. The factory function validates the `customer` instance before returning it. If the instance fails validation, the factory function returns an error. This demonstrates how to encapsulate validation logic within a function.

## How to Use

To use this project, you can import the structs into your own code. As the `newCustomer` function is private, it can only be used within the same package. Here's an example of how you can create a new `customer` instance within the same package:

```go
package main

import (
    "fmt"
)

func main() {
    customer, err := newCustomer("John Doe", 20)
    if err != nil {
        fmt.Println("Error creating customer:", err)
        return
    }

    fmt.Println("Created customer:", customer)
}