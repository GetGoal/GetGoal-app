package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	// "gopkg.in/go-playground/validator.v8"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})

	// Check if the provided error is of type validator.ValidationErrors
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErrs {
			// Can translate each error one at a time.
			if v.Param() != "" {
				res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
			} else {
				res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
			}
		}
	}

	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func Validate(modelValidate interface{}) error {
	validate := validator.New()
	err := validate.Struct(modelValidate)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, _ := json.MarshalIndent(messages, "", "  ")
		return errors.New(string(jsonMessage))
	}
	return nil
}

func GetTimeNow() time.Time {
	return time.Now()
}
