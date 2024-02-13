package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xbklyn/getgoal-app/controller"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
)

type mockLabelService struct{}

// FindAllLabels returns a slice of labels
func (m *mockLabelService) FindAllLabels() ([]entity.Label, error) {
	return []entity.Label{
		{
			LabelID:   1,
			LabelName: "Label 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			LabelID:   2,
			LabelName: "Label 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

// FindLabelByID returns a label by ID
func (m *mockLabelService) FindLabelByID(id uint64) (*entity.Label, error) {
	if id == 1 {
		return &entity.Label{
			LabelID:   1,
			LabelName: "Label 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}
	return nil, errors.New("label not found")
}

// GetSearchLabel returns a slice of labels
func (m *mockLabelService) GetSearchLabel() ([]entity.Label, error) {
	return []entity.Label{
		{
			LabelID:   1,
			LabelName: "Label 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

// Save creates a new label
func (m *mockLabelService) Save(label model.LabelRequest) (*entity.Label, error) {
	if label.LabelName == "" {
		return nil, errors.New("empty label name")
	}
	return &entity.Label{
		LabelID:   1,
		LabelName: "New Label",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *mockLabelService) Update(id uint64, label model.LabelRequest) (*entity.Label, error) {
	if id == 1 {
		if label.LabelName == "" {
			return nil, errors.New("label_name is required")
		}
		return &entity.Label{LabelID: id, LabelName: "Updated Label"}, nil
	}
	return nil, errors.New("label not found")
}

func (m *mockLabelService) Delete(id uint64) error {
	if id == 1 {
		return nil
	}
	return errors.New("label not found")
}

var (
	r               *gin.Engine
	labelService    *mockLabelService
	labelController *controller.LabelController
)

func setup() {
	gin.SetMode(gin.TestMode)
	r = gin.New()
	labelService = &mockLabelService{}
	labelController = controller.NewLabelController(labelService)
	api := r.Group("/api")
	labelController.Route(api)
}

func TestFindAllLabels_Success(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/labels", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
	assert.Equal(t, 2, response.Count) // Check the count of labels
}

func TestFindLabelByLabelID_Success(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/labels/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
	assert.Equal(t, 1, response.Count) // Check the count of labels
}

func TestGetSearchLabel_Success(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/labels/search", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Success", response.Message)
	assert.Equal(t, 1, response.Count) // Check the count of labels
}

func TestSaveLabel_Success(t *testing.T) {
	setup()

	payload := `{"label_name":"New Label"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/labels", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "Success", response.Message)
	assert.Equal(t, 1, response.Count) // Check the count of labels
	assert.Equal(t, "New Label", response.Data.(map[string]interface{})["label_name"])
}

func TestSaveLabel_EmptyLabelName(t *testing.T) {
	setup()

	payload := `{"label_name":""}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/labels", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "Bad Request", response.Message)
	assert.Equal(t, 0, response.Count) // Check the count of labels
	assert.Contains(t, response.Error, "empty label name")
}

func TestFindLabelByLabelID_LabelNotFound(t *testing.T) {
	setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/labels/2", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "Bad Request", response.Message)
	assert.Equal(t, 0, response.Count) // Check the count of labels
	assert.Contains(t, response.Error, "label not found")
}
func TestUpdateLabel_Success(t *testing.T) {
	setup()

	id := uint64(1)
	payload := `{"label_name":"Updated Label"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/labels/"+strconv.FormatUint(id, 10), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Label", response.Data.(map[string]interface{})["label_name"])
}

func TestUpdateLabel_InvalidID(t *testing.T) {
	setup()

	id := uint64(100) // Assuming ID doesn't exist
	payload := `{"label_name":"Updated Label"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/labels/"+strconv.FormatUint(id, 10), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "label not found", response.Error)
}

func TestUpdateLabel_InvalidPayload(t *testing.T) {
	setup()

	id := uint64(1)
	payload := `{"invalid_field":"Updated Label"}` // Invalid payload

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/labels/"+strconv.FormatUint(id, 10), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Bad Request", response.Message)
	assert.Contains(t, response.Error, "label_name is required")
}

func TestDeleteLabel_Success(t *testing.T) {
	setup()

	id := uint64(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/labels/"+strconv.FormatUint(id, 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Deleted id 1 successfully", response.Message)
}

func TestDeleteLabel_InvalidID(t *testing.T) {
	setup()

	id := uint64(100) // Assuming ID doesn't exist

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/labels/"+strconv.FormatUint(id, 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response model.GeneralResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "label not found", response.Error)
}
