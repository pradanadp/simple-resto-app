package repository

import (
	"errors"

	"github.com/pradanadp/simple-resto-app/app/middlewares"
	"github.com/pradanadp/simple-resto-app/features/customer"

	utils "github.com/pradanadp/simple-resto-app/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	log = logrus.New()
)

type customerQuery struct {
	db *gorm.DB
}

// Login implements customer.CustomerRepository.
func (cq *customerQuery) Login(req customer.CustomerEntity) (customer.CustomerEntity, string, error) {
	result := Customer{}
	query := cq.db.Table("customers").Where("email = ?", req.Email).First(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("customer record not found, invalid email and password")
		return customer.CustomerEntity{}, "", errors.New("invalid email and password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no customer has been created")
		return customer.CustomerEntity{}, "", errors.New("no row affected")
	}

	if !utils.MatchPassword(req.Password, result.Password) {
		log.Warn("password does not match")
		return customer.CustomerEntity{}, "", errors.New("password does not match")
	}

	token, err := middlewares.GenerateToken(result.CustomerID)
	if err != nil {
		log.Error("error while creating jwt token")
		return customer.CustomerEntity{}, "", errors.New("error while creating jwt token")
	}

	log.Infof("user has been logged in: %s", result.CustomerID)
	return CustomerModelToEntity(result), token, nil
}

// Register implements customer.CustomerRepository.
func (cq *customerQuery) Register(req customer.CustomerEntity) (customer.CustomerEntity, error) {
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Error("error while hashing password")
		return customer.CustomerEntity{}, errors.New("error while hashing password")
	}

	req.Password = hashedPass
	request := CustomerEntityToModel(req)
	query := cq.db.Table("customers").Create(&request)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return customer.CustomerEntity{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no customer has been created")
		return customer.CustomerEntity{}, errors.New("no row affected")
	}

	log.Infof("new customer has been created: %s", req.Email)
	return CustomerModelToEntity(request), nil
}

func New(db *gorm.DB) customer.Repository {
	return &customerQuery{
		db: db,
	}
}
