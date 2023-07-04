package controllers_test

import (
	controllers "agustinzunda/usersadmingo/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := controllers.NewUserController()

	router := gin.Default()
	router.GET("/ping", controller.PingHandler)

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}\n", w.Body.String())
}
