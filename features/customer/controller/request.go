package controller

import "github.com/pradanadp/simple-resto-app/features/customer"

type RegisterRequest struct {
	FirstName   string `json:"firstname" form:"firstname"`
	LastName    string `json:"lastname" form:"lastname"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone" form:"phone"`
	Password    string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToEntity(data interface{}) customer.CustomerEntity {
	res := customer.CustomerEntity{}
	switch v := data.(type) {
	case RegisterRequest:
		res.FirstName = v.FirstName
		res.Email = v.Email
		res.PhoneNumber = v.PhoneNumber
		res.Password = v.Password
	case LoginRequest:
		res.Email = v.Email
		res.Password = v.Password
	default:
		return customer.CustomerEntity{}
	}

	return res
}
