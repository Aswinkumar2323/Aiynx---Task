package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1"`
	DOB  string `json:"dob" validate:"required,date_yyyy_mm_dd"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1"`
	DOB  string `json:"dob" validate:"required,date_yyyy_mm_dd"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  *int   `json:"age,omitempty"` // omitempty, or just serialize it as int when requested
}

var Validate = validator.New()

func InitValidator() {
	_ = Validate.RegisterValidation("date_yyyy_mm_dd", func(fl validator.FieldLevel) bool {
		str := fl.Field().String()
		_, err := time.Parse("2006-01-02", str)
		return err == nil
	})
}
