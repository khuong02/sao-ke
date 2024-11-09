package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interview-rest/pkg/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ResponseData struct {
	Message string `json:"message"`
}

func TestRecoverPanic(t *testing.T) {
	type testCase struct {
		body           model.User
		expectResponse ResponseData
	}

	testCases := []testCase{
		{
			body: model.User{
				Email: "abc@gmail.com",
			},
			expectResponse: ResponseData{Message: "Internal server error"},
		},
	}
	svc := NewServer(8080)
	r := gin.Default()

	r.Use(svc.recoverPanic())
	r.POST("/validate", svc.validateUserHandler)

	for _, tc := range testCases {
		data, _ := json.Marshal(tc.body)
		resp := ResponseData{}
		req, _ := http.NewRequest("POST", "/validate", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		_ = json.Unmarshal(responseData, &resp)
		assert.Equal(t, tc.expectResponse.Message, resp.Message)
	}
}
