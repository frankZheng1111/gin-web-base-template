package controller

import (
	"encoding/json"
	db "gin-web-base-template/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidationErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", ValidationErrorResponse)
	r.ServeHTTP(w, req)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "ParamsValidationError", respJson["msg"])
}

func TestMissingLockErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", MissingLockErrorResponse)
	r.ServeHTTP(w, req)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "OverFrequency", respJson["msg"])
}

func TestCommonPanicHandle(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", CommonPanicHandle(func(c *gin.Context) {
		panic("testError")
	}))
	r.ServeHTTP(w, req)

	var respJson map[string]interface{}

	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "SERVER_ERROR", respJson["msg"])
}

func TestParsePaginateFromQuery(t *testing.T) {
	var paginate *db.Paginate
	req, err := http.NewRequest("GET", "/test?page=2&page_size=19", nil)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		paginate = ParsePaginateFromQuery(c)
	})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 2, paginate.Page)
	assert.Equal(t, 19, paginate.PageSize)
}

func TestGetAndSession(t *testing.T) {
	var sessionValue string
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/test", func(c *gin.Context) {
		SetSession(c, "SessionKey", "SessionValue")
		sessionValue = GetSession(c, "SessionKey").(string)
	})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, "SessionValue", sessionValue)
}
