package service

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pradanadp/simple-resto-app/features/customer"
	"github.com/pradanadp/simple-resto-app/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type customerService struct {
	customerRepository customer.CustomerRepository
	validator          *validator.Validate
}

// Login implements customer.CustomerService.
func (cs *customerService) Login(req customer.CustomerEntity) (customer.CustomerEntity, string, error) {
	err := cs.validator.Struct(req)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Email"):
			log.Warn("invalid email format")
			return customer.CustomerEntity{}, "", errors.New("invalid email format")
		case strings.Contains(err.Error(), "Password"):
			log.Warn("password cannot be empty")
			return customer.CustomerEntity{}, "", errors.New("password cannot be empty")
		}
	}

	result, token, err := cs.customerRepository.Login(req)
	if err != nil {
		message := ""
		switch {
		case strings.Contains(err.Error(), "invalid email and password"):
			log.Error("invalid email and password")
			message = "invalid email and password"
		case strings.Contains(err.Error(), "password does not match"):
			log.Error("password does not match")
			message = "password does not match"
		case strings.Contains(err.Error(), "no row affected"):
			log.Error("no row affected")
			message = "no row affected"
		case strings.Contains(err.Error(), "error while creating jwt token"):
			log.Error("error while creating jwt token")
			message = "error while creating jwt token"
		default:
			log.Error("internal server error")
			message = "internal server error"
		}
		return customer.CustomerEntity{}, "", errors.New(message)
	}

	log.Infof("user has been logged in: %s", result.CustomerID)
	return result, token, nil
}

// Register implements customer.CustomerService.
func (cs *customerService) Register(req customer.CustomerEntity) (customer.CustomerEntity, error) {
	customerID := utils.GenerateCustomerID()
	req.CustomerID = customerID

	err := utils.ValidatePassword(req.Password)
	if err != nil {
		log.Error(err.Error())
		return customer.CustomerEntity{}, err
	}

	_, isValid := utils.ValidateMailAddress(req.Email)
	if !isValid {
		log.Error("wrong email format")
		return customer.CustomerEntity{}, errors.New("wrong email format")
	}

	if req.FirstName == "" {
		log.Error("firstname is required")
		return customer.CustomerEntity{}, errors.New("firstname is required")
	}

	if req.PhoneNumber == "" {
		log.Error("phone number is required")
		return customer.CustomerEntity{}, errors.New("phone number is required")
	}

	newUser, err := cs.customerRepository.Register(req)
	if err != nil {
		log.Error(err.Error())
		return customer.CustomerEntity{}, err
	}

	return newUser, nil
}

func New(c customer.CustomerRepository, v *validator.Validate) customer.CustomerService {
	return &customerService{
		customerRepository: c,
		validator:          v,
	}
}
