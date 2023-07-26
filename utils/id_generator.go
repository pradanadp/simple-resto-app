package utils

import "github.com/google/uuid"

func GenerateCustomerID() string {
	return "cstm" + uuid.New().String()
}

func GenerateCartID() string {
	return "crt" + uuid.New().String()
}
