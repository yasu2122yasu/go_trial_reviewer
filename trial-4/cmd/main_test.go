package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestTop(t *testing.T) {
	Response := `{"200":"The Index Method is passed"}`
	// Response := `{"400":"The Index Method is passed"}`
	r := SetUpRouter()
	r.GET("/", Top)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, Response, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreate(t *testing.T) {
	r := SetUpRouter()
	r.POST("/information", Create)
	// r.GET("/information", Create)
	informationId := xid.New().String()
	information := Information{
		ID:      informationId,
		Name:    "Yamada Taro",
		Address: "Sendai",
		Pin:     7889121,
	}
	jsonValue, _ := json.Marshal(information)
	req, _ := http.NewRequest("POST", "/information", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestIndex(t *testing.T) {
	r := SetUpRouter()
	r.GET("/informations", Index)
	req, _ := http.NewRequest("GET", "/informations", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var informations []Information
	json.Unmarshal(w.Body.Bytes(), &informations)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, informations)
}

func TestUpdate(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/information/:id", Update)
	information := Information{
		ID:      `1`,
		Name:    "Yamada Taro",
		Address: "Sendai",
		Pin:     7889121,
	}
	jsonValue, _ := json.Marshal(information)
	reqFound, _ := http.NewRequest("PUT", "/information/"+information.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	// reqNotFound, _ := http.NewRequest("PUT", "/information/12", bytes.NewBuffer(jsonValue))
	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, reqNotFound)
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
