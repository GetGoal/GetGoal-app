package common_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xbklyn/getgoal-app/common"
)

func TestNewError(t *testing.T) {
	// Test creating a new error with a key and message
	errKey := "testKey"
	errMessage := errors.New("test error message")

	commonErr := common.NewError(errKey, errMessage)

	assert.NotNil(t, commonErr)
	assert.NotNil(t, commonErr.Errors)
	assert.Equal(t, 1, len(commonErr.Errors))
	assert.Equal(t, errMessage.Error(), commonErr.Errors[errKey])

	// Test creating a new error with an empty key
	commonErrEmptyKey := common.NewError("", errMessage)

	assert.NotNil(t, commonErrEmptyKey)
	assert.NotNil(t, commonErrEmptyKey.Errors)
	assert.Equal(t, 1, len(commonErrEmptyKey.Errors))
	assert.Equal(t, errMessage.Error(), commonErrEmptyKey.Errors[""])
}

type TestObject struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func TestBind(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Create a mock request
	body := []byte("name=John&age=30")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)

	// Create a mock gin.Context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Create an object to bind
	obj := &TestObject{}

	// Call the Bind function
	err := common.Bind(c, obj)

	// Assertions
	assert.NoError(t, err, "Bind should not return an error")
	assert.Equal(t, "John", obj.Name, "Name should be 'John'")
	assert.Equal(t, 30, obj.Age, "Age should be 30")
}

func TestValidate(t *testing.T) {
	// Test validating a struct with missing required fields
	type SampleStruct struct {
		Email string `json:"email" binding:"required,email" validate:"required,email"`
	}

	invalidStruct := SampleStruct{}

	err := common.Validate(invalidStruct)

	assert.Error(t, err)

	var errMsg []map[string]interface{}
	json.Unmarshal([]byte(err.Error()), &errMsg)

	assert.Len(t, errMsg, 1)
	assert.Equal(t, "Email", errMsg[0]["field"])
	assert.Contains(t, errMsg[0]["message"], "this field is required")

	// Test validating a struct with valid fields
	validStruct := SampleStruct{
		Email: "test@example.com",
	}

	errValid := common.Validate(validStruct)

	assert.NoError(t, errValid)
}

func TestGetTimeNow(t *testing.T) {
	// Test getting the current time
	currentTime := common.GetTimeNow()

	assert.NotNil(t, currentTime)
	assert.True(t, currentTime.Before(time.Now().Add(1*time.Second)))
	assert.True(t, currentTime.After(time.Now().Add(-1*time.Second)))
}
