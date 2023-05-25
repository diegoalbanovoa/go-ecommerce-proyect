package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	router := setupRouter()

	// Crea una solicitud HTTP de prueba POST a "/users/signup"
	signupURL := "/users/signup"
	signupPayload := map[string]interface{}{
		"first_name": "Akhil",
		"last_name":  "Sharma",
		"email":      "akhil@gmail.com",
		"password":   "akhilsharma",
		"phone":      "+4534545435",
	}
	signupJSON, _ := json.Marshal(signupPayload)

	req, _ := http.NewRequest("POST", signupURL, bytes.NewBuffer(signupJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Successfully Signed Up!!", resp.Body.String())
}

func TestLogin(t *testing.T) {
	router := setupRouter()

	// Crea una solicitud HTTP de prueba POST a "/users/login"
	loginURL := "/users/login"
	loginPayload := map[string]interface{}{
		"email":    "akhil@gmail.com",
		"password": "akhilsharma",
	}
	loginJSON, _ := json.Marshal(loginPayload)

	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Verifica la respuesta JSON devuelta
	expectedResponse := map[string]interface{}{
		"_id":           "***********************",
		"first_name":    "akhil",
		"last_name":     "sharma",
		"password":      "$2a$14$UIYjkTfnFnhg4qhIfhtYnuK9qsBQifPKgu/WPZAYBaaN17j0eTQZa",
		"email":         "akhil@gmail.com",
		"phone":         "+4534545435",
		"token":         "eyJc0Bwcm90b25vbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiam9zZXBoIiwiTGFzdF9OYW1lIjoiaGVybWlzIiwiVWlkIjoiNjE2MTRmNTM5ZjI5YmU5NDJiZDlkZjhlIiwiZXhwIjoxNjMzODUzNjUxfQ.NbcpVtPLJJqRF44OLwoanynoejsjdJb5_v2qB41SmB8",
		"Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnLCJVaWQiOiIiLCJleHAiOjE2MzQzNzIwNTF9.ocpU8-0gCJsejmCeeEiL8DXhFcZsW7Z3OCN34HgIf2c",
		"created_at":    "2022-04-09T08:14:11Z",
		"updtaed_at":    "2022-04-09T08:14:11Z",
		"user_id":       "61614f539f29be942bd9df8e",
		"usercart":      []interface{}{},
		"address":       []interface{}{},
		"orders":        []interface{}{},
	}
	expectedJSON, _ := json.Marshal(expectedResponse)

	assert.JSONEq(t, string(expectedJSON), resp.Body.String())
}

// Agrega más pruebas unitarias para las demás funciones aquí

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/users/signup", func(c *gin.Context) {
		c.String(http.StatusOK, "Successfully Signed Up!!")
	})

	router.POST("/users/login", func(c *gin.Context) {
		// Simula la respuesta de autenticación con los datos del usuario
		user := map[string]interface{}{
			"_id":           "***********************",
			"first_name":    "akhil",
			"last_name":     "sharma",
			"password":      "$2a$14$UIYjkTfnFnhg4qhIfhtYnuK9qsBQifPKgu/WPZAYBaaN17j0eTQZa",
			"email":         "akhil@gmail.com",
			"phone":         "+4534545435",
			"token":         "eyJc0Bwcm90b25vbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiam9zZXBoIiwiTGFzdF9OYW1lIjoiaGVybWlzIiwiVWlkIjoiNjE2MTRmNTM5ZjI5YmU5NDJiZDlkZjhlIiwiZXhwIjoxNjMzODUzNjUxfQ.NbcpVtPLJJqRF44OLwoanynoejsjdJb5_v2qB41SmB8",
			"Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnLCJVaWQiOiIiLCJleHAiOjE2MzQzNzIwNTF9.ocpU8-0gCJsejmCeeEiL8DXhFcZsW7Z3OCN34HgIf2c",
			"created_at":    "2022-04-09T08:14:11Z",
			"updtaed_at":    "2022-04-09T08:14:11Z",
			"user_id":       "61614f539f29be942bd9df8e",
			"usercart":      []interface{}{},
			"address":       []interface{}{},
			"orders":        []interface{}{},
		}
		c.JSON(http.StatusOK, user)
	})

	// Configura las demás rutas y controladores aquí

	return router
}
